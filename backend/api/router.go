// api/router.go
package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	// handlerパッケージをインポート
	"github.com/sintaro/FlowGrid/backend/api/handler"
 )

// SetupRouter はAuthHandlerとTaskHandlerを受け取り、ルーターをセットアップします
func SetupRouter(authHandler *handler.AuthHandler, taskHandler *handler.TaskHandler, projectHandler *handler.ProjectHandler) *gin.Engine {
	r := gin.Default()

	// ヘルスチェックエンドポイント
	r.GET("/health", healthCheck)

	// ルーティング設定
	// ハンドラーのメソッドを呼び出す形に修正
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	authRequired := r.Group("/")
	// ミドルウェアもauthHandlerのメソッドとして定義するとより綺麗ですが、
	// ここでは簡単のため、このファイル内に定義します。
	authRequired.Use(authMiddleware())
	{
		authRequired.GET("/profile", authHandler.Profile)
		authRequired.POST("/logout", authHandler.Logout)
		
		// タスク関連のルート
		authRequired.GET("/tasks", taskHandler.GetAllTasksGroupedByStatus)
		authRequired.GET("/tasks/filter", taskHandler.GetTasksByStatus)
		
		// プロジェクト進捗率関連のルート
		authRequired.GET("/projects/progress", projectHandler.GetAllProjectsProgress)
		authRequired.GET("/projects/:id/progress", projectHandler.GetProjectProgress)
		authRequired.GET("/sprints/:id/progress", projectHandler.GetSprintProgress)
	}

	return r
}

// healthCheck はヘルスチェックエンドポイントのハンドラー
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"message": "Server is running",
	})
}

var blacklistedTokens = make(map[string]bool) // ブラックリストされたトークンを保存（handler側と共有）

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" { c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"} ); c.Abort(); return }
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" { c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"} ); c.Abort(); return }
		tokenString := parts[1]
		
		// ブラックリストされたトークンをチェック
		if blacklistedTokens[tokenString] {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been revoked"})
			c.Abort()
			return
		}
		
		// JWTシークレットキーを環境変数から取得
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret not configured"})
			c.Abort()
			return
		}
		
		claims := &Claims{} // ここでClaimsを使えるようにする
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return []byte(jwtSecret), nil })
		if err != nil || !token.Valid { c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"} ); c.Abort(); return }
		c.Set("username", claims.Username)
		c.Next()
	}
}
