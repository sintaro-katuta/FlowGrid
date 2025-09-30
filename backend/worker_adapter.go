// Cloudflare Workers用のシンプルなHello World API
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

	// パスに基づいて異なるレスポンスを返す
	switch r.URL.Path {
	case "/":
		response := map[string]string{
			"message": "Hello World from FlowGrid API!",
			"status":  "success",
			"version": "1.0.0",
		}
		jsonResponse(w, http.StatusOK, response)
	case "/health":
		response := map[string]string{
			"status": "healthy",
		}
		jsonResponse(w, http.StatusOK, response)
	default:
		response := map[string]string{
			"error": "Not Found",
		}
		jsonResponse(w, http.StatusNotFound, response)
	}
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
