// Cloudflare Workers用メインコード
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Cloudflare Workersのエントリーポイント
func main() {
	// Cloudflare Workersではmain()は空で、exportされた関数が呼び出される
}

// メインハンドラー - Cloudflare Workersのエントリーポイント
//export HandleRequest
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// CORS設定
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// パスに基づいてルーティング
	path := r.URL.Path
	method := r.Method

	switch {
	case path == "/health" && method == "GET":
		handleHealthCheck(w, r)
	case strings.HasPrefix(path, "/auth"):
		handleAuthRoutes(w, r)
	case strings.HasPrefix(path, "/projects"):
		handleProjectRoutes(w, r)
	case strings.HasPrefix(path, "/tasks"):
		handleTaskRoutes(w, r)
	default:
		jsonResponse(w, http.StatusNotFound, map[string]string{
			"error": "Route not found",
		})
	}
}

// ヘルスチェック
func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "FlowGrid API is running on Cloudflare Workers",
	})
}

// 認証ルートの処理
func handleAuthRoutes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	switch {
	case path == "/auth/register" && method == "POST":
		handleRegister(w, r)
	case path == "/auth/login" && method == "POST":
		handleLogin(w, r)
	default:
		jsonResponse(w, http.StatusNotFound, map[string]string{
			"error": "Auth route not found",
		})
	}
}

// プロジェクトルートの処理
func handleProjectRoutes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	switch {
	case path == "/projects" && method == "GET":
		handleGetAllProjects(w, r)
	case strings.HasPrefix(path, "/projects/") && method == "GET":
		// プロジェクトIDの抽出
		parts := strings.Split(path, "/")
		if len(parts) >= 3 {
			projectID := parts[2]
			handleGetProject(w, r, projectID)
		} else {
			jsonResponse(w, http.StatusBadRequest, map[string]string{
				"error": "Invalid project ID",
			})
		}
	case strings.HasPrefix(path, "/projects/sprint/") && method == "GET":
		// スプリントIDの抽出
		parts := strings.Split(path, "/")
		if len(parts) >= 4 {
			sprintID := parts[3]
			handleGetSprint(w, r, sprintID)
		} else {
			jsonResponse(w, http.StatusBadRequest, map[string]string{
				"error": "Invalid sprint ID",
			})
		}
	default:
		jsonResponse(w, http.StatusNotFound, map[string]string{
			"error": "Project route not found",
		})
	}
}

// タスクルートの処理
func handleTaskRoutes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	switch {
	case path == "/tasks" && method == "GET":
		handleGetAllTasks(w, r)
	case path == "/tasks/status" && method == "GET":
		handleGetTasksByStatus(w, r)
	default:
		jsonResponse(w, http.StatusNotFound, map[string]string{
			"error": "Task route not found",
		})
	}
}

// 簡易的なハンドラー関数
func handleRegister(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]string{
		"message": "Register endpoint - to be implemented",
	})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]string{
		"message": "Login endpoint - to be implemented",
	})
}

func handleGetAllProjects(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"projects": []map[string]interface{}{
			{"id": 1, "name": "Project 1", "progress": 75},
			{"id": 2, "name": "Project 2", "progress": 50},
		},
	})
}

func handleGetProject(w http.ResponseWriter, r *http.Request, projectID string) {
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"id":      projectID,
		"name":    fmt.Sprintf("Project %s", projectID),
		"progress": 75,
	})
}

func handleGetSprint(w http.ResponseWriter, r *http.Request, sprintID string) {
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"id":      sprintID,
		"name":    fmt.Sprintf("Sprint %s", sprintID),
		"progress": 60,
	})
}

func handleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"tasks": map[string][]map[string]interface{}{
			"todo": {
				{"id": 1, "title": "Task 1", "status": "todo"},
				{"id": 2, "title": "Task 2", "status": "todo"},
			},
			"in_progress": {
				{"id": 3, "title": "Task 3", "status": "in_progress"},
			},
			"done": {
				{"id": 4, "title": "Task 4", "status": "done"},
			},
		},
	})
}

func handleGetTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		status = "todo"
	}
	
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"status": status,
		"tasks": []map[string]interface{}{
			{"id": 1, "title": fmt.Sprintf("Task for %s", status), "status": status},
			{"id": 2, "title": fmt.Sprintf("Another task for %s", status), "status": status},
		},
	})
}

// 簡易的なJSONレスポンス関数
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
