// main.go
package main

import (
	"log"
	"os"
	"github.com/sintaro/FlowGrid/backend/api"
	"github.com/sintaro/FlowGrid/backend/api/handler" // handlerパッケージもインポート
	"github.com/sintaro/FlowGrid/backend/models" // modelsパッケージをインポート
)

func main() {
	// 1. データベース接続の確立
	db, err := models.DBConnect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// 2. 依存関係のインスタンス化
	// データベース接続を渡してハンドラーを生成
	authHandler := handler.NewAuthHandler(db)
	taskHandler := handler.NewTaskHandler(db)
	projectHandler := handler.NewProjectHandler(db)

	// 3. 依存性の注入
	// 生成したハンドラーをルーターに渡す
	router := api.SetupRouter(authHandler, taskHandler, projectHandler)

	// 4. サーバーの起動
	port := getPort()
	router.Run(":" + port)
}

// getPort は環境変数 PORT からポート番号を取得し、デフォルトは "8080"
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
