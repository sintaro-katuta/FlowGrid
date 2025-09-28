// main.go
package main

import (
	"log"
	"github.com/sintaro/FlowGrid/backend/api"
	"github.com/sintaro/FlowGrid/backend/api/handler" // handlerパッケージもインポート
	"github.com/sintaro/FlowGrid/backend/db"
	"github.com/sintaro/FlowGrid/backend/models" // modelsパッケージをインポート
)

func main() {
	// 1. データベース接続の確立
	dbConn, err := models.DBConnect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer dbConn.Close()

	// 2. データベースインターフェースのラップ
	sqlDB := &db.SQLDatabase{DB: dbConn}

	// 3. 依存関係のインスタンス化
	// データベース接続を渡してハンドラーを生成
	authHandler := handler.NewAuthHandler(sqlDB)
	taskHandler := handler.NewTaskHandler(sqlDB)
	projectHandler := handler.NewProjectHandler(sqlDB)

	// 4. 依存性の注入
	// 生成したハンドラーをルーターに渡す
	router := api.SetupRouter(authHandler, taskHandler, projectHandler)

	// 5. サーバーの起動
	router.Run(":8080")
}
