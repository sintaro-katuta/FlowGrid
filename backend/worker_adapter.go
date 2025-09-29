// Cloudflare Workers用アダプター
package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sintaro/FlowGrid/backend/api"
	"github.com/sintaro/FlowGrid/backend/api/handler"
	"github.com/sintaro/FlowGrid/backend/db"
)

// Cloudflare Workers用のハンドラー
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// CORS設定
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Cloudflare Workers環境では、データベース接続はWranglerのバインド経由で処理される
	// ここではモックデータベースを使用（実際のデプロイ時は適切なD1アダプターに置き換え）
	var database db.Database
	
	// 環境変数で本番/開発環境を判定
	if os.Getenv("ENVIRONMENT") == "production" {
		database = &db.MockDatabase{} // 本番環境用モック（実際はD1アダプター）
	} else {
		database = &db.MockDatabase{} // 開発環境用モック
	}

	// 依存関係のインスタンス化
	authHandler := handler.NewAuthHandler(database)
	taskHandler := handler.NewTaskHandler(database)
	projectHandler := handler.NewProjectHandler(database)

	// ルーターのセットアップ
	router := api.SetupRouter(authHandler, taskHandler, projectHandler)

	// GinルーターをHTTPハンドラーとして使用
	router.ServeHTTP(w, r)
}

// 簡易的なJSONレスポンス関数
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Cloudflare Workersのメインハンドラー
func main() {
	// Cloudflare Workersではmain()は空で、HandleRequestが自動的に呼び出される
}
