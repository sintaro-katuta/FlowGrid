// api/handler/auth_handler.go
package handler // パッケージ名を handler にします

import (
	"database/sql"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/sintaro/FlowGrid/backend/db"
 )

// --- グローバルではなく、ハンドラー構造体の一部として定義 ---
type AuthHandler struct {
	DB db.Database
}

// NewAuthHandler はAuthHandlerのインスタンスを生成します
func NewAuthHandler(db db.Database) *AuthHandler {
	return &AuthHandler{DB: db}
}

// --- データとJWT関連の定義 (本来はより適切な場所に移動) ---
var users = make(map[string]string)
var blacklistedTokens = make(map[string]bool) // ブラックリストされたトークンを保存

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// getJWTKey は環境変数からJWTシークレットを取得します
func getJWTKey() []byte {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return []byte("fallback-secret-key")
	}
	return []byte(jwtSecret)
}
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct { Email string `json:"email"`; Password string `json:"password"` }
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"} ); return }
	
	// パスワードをハッシュ化（メモリ内認証用）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"} ); return }
	
	// まずユーザーが既に存在するかチェック
	var existingUserID int
	err = h.DB.QueryRow("SELECT id FROM users WHERE email = ?", req.Email).Scan(&existingUserID)
	if err == nil {
		// ユーザーが既に存在する場合
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	} else if err != sql.ErrNoRows {
		// データベースエラーの場合
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	
	// ユーザーが存在しない場合は新規作成
	// nameフィールドには一旦emailと同じ値を設定
	_, err = h.DB.Exec("INSERT INTO users (name, email, role_id) VALUES (?, ?, ?)", 
		req.Email, req.Email, 2)
	if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user to database"} )
		return 
	}
	
	// パスワードはメモリ内に保存（データベースには保存しない）
	users[req.Email] = string(hashedPassword)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"} )
}

// Login はログイン処理を行います
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct { Email string `json:"email"`; Password string `json:"password"` }
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"} ); return }
	storedPassword, ok := users[req.Email]
	if !ok { c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"} ); return }
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password)); err != nil { c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"} ); return }
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{ Username: req.Email, RegisteredClaims: jwt.RegisteredClaims{ ExpiresAt: jwt.NewNumericDate(expirationTime), }, }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTKey())
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"} ); return }
	c.JSON(http.StatusOK, gin.H{"token": tokenString} )
}

// Logout はログアウト処理を行います
func (h *AuthHandler) Logout(c *gin.Context) {
	// Authorizationヘッダーからトークンを取得
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}

	// "Bearer "プレフィックスを確認
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header format must be Bearer {token}"})
		return
	}

	tokenString := parts[1]

	// トークンを検証
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTKey(), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// トークンをブラックリストに追加
	blacklistedTokens[tokenString] = true

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// Profile はプロフィール表示処理を行います
func (h *AuthHandler) Profile(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{ "message":  "This is a protected route", "username": username, } )
}
