// api/handler/task_handler.go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sintaro/FlowGrid/backend/db"
)

// TaskHandler はタスク関連のAPIハンドラー
type TaskHandler struct {
	DB db.Database
}

// NewTaskHandler はTaskHandlerのインスタンスを生成します
func NewTaskHandler(db db.Database) *TaskHandler {
	return &TaskHandler{DB: db}
}

// TaskResponse はタスクのレスポンス構造体
type TaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserName    string `json:"user_name"`
	ProjectName string `json:"project_name"`
	SprintName  string `json:"sprint_name"`
	StatusName  string `json:"status_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TasksByStatusResponse はステータスごとにグループ化されたタスクのレスポンス
type TasksByStatusResponse struct {
	Todo       []TaskResponse `json:"todo"`
	InProgress []TaskResponse `json:"in_progress"`
	Done       []TaskResponse `json:"done"`
}

// GetAllTasksGroupedByStatus は全タスクをステータスごとにグループ化して取得します
func (h *TaskHandler) GetAllTasksGroupedByStatus(c *gin.Context) {
	// データベースからタスクを取得（ステータスごとにグループ化）
	query := `
		SELECT 
			t.id, 
			t.title, 
			t.description, 
			u.name as user_name,
			p.name as project_name,
			s.name as sprint_name,
			st.name as status_name,
			t.created_at,
			t.updated_at
		FROM tasks t
		JOIN users u ON t.user_id = u.id
		JOIN projects p ON t.project_id = p.id
		JOIN sprint s ON t.sprint_id = s.id
		JOIN statuses st ON t.status_id = st.id
		ORDER BY st.name, t.created_at
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer rows.Close()

	// ステータスごとにタスクをグループ化
	tasksByStatus := make(map[string][]TaskResponse)

	for rows.Next() {
		var task TaskResponse
		var createdAt, updatedAt string

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.UserName,
			&task.ProjectName,
			&task.SprintName,
			&task.StatusName,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan task data"})
			return
		}

		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt

		// ステータス名でグループ化
		tasksByStatus[task.StatusName] = append(tasksByStatus[task.StatusName], task)
	}

	// レスポンス形式に変換
	response := TasksByStatusResponse{
		Todo:       []TaskResponse{},
		InProgress: []TaskResponse{},
		Done:       []TaskResponse{},
	}

	// ステータス名に基づいて分類
	for statusName, tasks := range tasksByStatus {
		switch statusName {
		case "todo", "To Do", "TODO":
			response.Todo = tasks
		case "in progress", "In Progress", "IN_PROGRESS":
			response.InProgress = tasks
		case "done", "Done", "DONE":
			response.Done = tasks
		default:
			// 不明なステータスはtodoに分類
			response.Todo = append(response.Todo, tasks...)
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetTasksByStatus は特定のステータスのタスクを取得します
func (h *TaskHandler) GetTasksByStatus(c *gin.Context) {
	statusID := c.Query("status")
	if statusID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status parameter is required"})
		return
	}

	query := `
		SELECT 
			t.id, 
			t.title, 
			t.description, 
			u.name as user_name,
			p.name as project_name,
			s.name as sprint_name,
			st.name as status_name,
			t.created_at,
			t.updated_at
		FROM tasks t
		JOIN users u ON t.user_id = u.id
		JOIN projects p ON t.project_id = p.id
		JOIN sprint s ON t.sprint_id = s.id
		JOIN statuses st ON t.status_id = st.id
		WHERE t.status_id = ?
		ORDER BY t.created_at
	`

	rows, err := h.DB.Query(query, statusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer rows.Close()

	var tasks []TaskResponse

	for rows.Next() {
		var task TaskResponse
		var createdAt, updatedAt string

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.UserName,
			&task.ProjectName,
			&task.SprintName,
			&task.StatusName,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan task data"})
			return
		}

		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
