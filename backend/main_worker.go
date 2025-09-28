// Cloudflare Workers用メインコード
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/sintaro/FlowGrid/backend/api/handler"

	"github.com/gin-gonic/gin"
)

// Cloudflare Workersの環境変数
type Env struct {
	DB interface {
		Exec(query string, args ...interface{}) error
		Prepare(query string) (*sql.Stmt, error)
	} `json:"DB"`
	JWT_SECRET string `json:"JWT_SECRET"`
}

// Cloudflare Workersのエントリーポイント
func main() {
	// この関数はCloudflare Workersによって呼び出される
	// 実際のリクエスト処理はHandleRequest関数で行う
}

// D1データベースアダプター
type D1DatabaseAdapter struct{}

func (d *D1DatabaseAdapter) Exec(query string, args ...interface{}) (sql.Result, error) {
	// Cloudflare D1のExecメソッドをシミュレート
	// 実際の実装ではenv.DB.Exec()を使用
	return nil, nil
}

func (d *D1DatabaseAdapter) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// Cloudflare D1のQueryメソッドをシミュレート
	return nil, nil
}

func (d *D1DatabaseAdapter) QueryRow(query string, args ...interface{}) *sql.Row {
	// Cloudflare D1のQueryRowメソッドをシミュレート
	return nil
}

func (d *D1DatabaseAdapter) Prepare(query string) (*sql.Stmt, error) {
	// Cloudflare D1のPrepareメソッドをシミュレート
	return nil, nil
}

// Cloudflare Workers用のハンドラー関数
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// 環境変数の取得（Cloudflare Workersから注入）
	env := &Env{
		JWT_SECRET: "your-jwt-secret", // 実際は環境変数から取得
	}

	// リクエストの処理
	router := setupGinRouter(env)
	
	// Ginルーターでリクエストを処理
	router.ServeHTTP(w, r)
}


// Ginルーターのセットアップ
func setupGinRouter(env *Env) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// CORS設定
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// データベースアダプターの作成
	db := &D1DatabaseAdapter{}

	// ハンドラーの初期化
	authHandler := handler.NewAuthHandler(db)
	taskHandler := handler.NewTaskHandler(db)
	projectHandler := handler.NewProjectHandler(db)

	// ルーティング設定
	setupRoutes(router, authHandler, taskHandler, projectHandler)

	return router
}

// ルーティング設定（既存のapi.SetupRouterから移植）
func setupRoutes(router *gin.Engine, authHandler *handler.AuthHandler, taskHandler *handler.TaskHandler, projectHandler *handler.ProjectHandler) {
	// 認証ルート
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// プロジェクトルート
	projects := router.Group("/projects")
	{
		projects.GET("/", projectHandler.GetAllProjectsProgress)
		projects.GET("/:id", projectHandler.GetProjectProgress)
		projects.GET("/sprint/:id", projectHandler.GetSprintProgress)
	}

	// タスクルート
	tasks := router.Group("/tasks")
	{
		tasks.GET("/", taskHandler.GetAllTasksGroupedByStatus)
		tasks.GET("/status", taskHandler.GetTasksByStatus)
	}

	// ヘルスチェック
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "FlowGrid API is running on Cloudflare Workers",
		})
	})
}

// 簡易的なJSONレスポンス関数
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
