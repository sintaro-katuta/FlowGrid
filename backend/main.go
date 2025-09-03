// main.go
package main

import (
	"github.com/sintaro/FlowGrid/backend/api"
	"github.com/sintaro/FlowGrid/backend/api/handler" // handlerパッケージもインポート
)

func main() {
	// 1. 依存関係のインスタンス化
	// まずハンドラーを生成
	authHandler := handler.NewAuthHandler()

	// 2. 依存性の注入
	// 生成したハンドラーをルーターに渡す
	router := api.SetupRouter(authHandler)

	// 3. サーバーの起動
	router.Run(":8080")
}
