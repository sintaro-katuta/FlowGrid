// Cloudflare Workers用アダプター
//go:build cloudflare
// +build cloudflare

package main

import (
	"encoding/json"
	"net/http"
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

	// シンプルなレスポンスを返す
	response := map[string]string{
		"message": "FlowGrid API is running on Cloudflare Workers",
		"status":  "success",
	}
	
	jsonResponse(w, http.StatusOK, response)
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
