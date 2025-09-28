// api/handler/project_handler.go
package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sintaro/FlowGrid/backend/db"
)

// ProjectHandler はプロジェクト関連のAPIハンドラー
type ProjectHandler struct {
	DB db.Database
}

// NewProjectHandler はProjectHandlerのインスタンスを生成します
func NewProjectHandler(db db.Database) *ProjectHandler {
	return &ProjectHandler{DB: db}
}

// ProjectProgressResponse はプロジェクトの進捗率レスポンス
type ProjectProgressResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	TotalTasks  int     `json:"total_tasks"`
	DoneTasks   int     `json:"done_tasks"`
	Progress    float64 `json:"progress"`
	SprintName  string  `json:"sprint_name,omitempty"`
	StartDate   string  `json:"start_date,omitempty"`
	EndDate     string  `json:"end_date,omitempty"`
}

// GetAllProjectsProgress は全プロジェクトの進捗率を取得します
func (h *ProjectHandler) GetAllProjectsProgress(c *gin.Context) {
	query := `
		SELECT 
			p.id,
			p.name,
			COUNT(t.id) as total_tasks,
			SUM(CASE WHEN st.name = 'done' THEN 1 ELSE 0 END) as done_tasks,
			s.name as sprint_name,
			s.start_date,
			s.end_date
		FROM projects p
		LEFT JOIN tasks t ON p.id = t.project_id
		LEFT JOIN statuses st ON t.status_id = st.id
		LEFT JOIN sprint s ON t.sprint_id = s.id
		GROUP BY p.id, p.name, s.name, s.start_date, s.end_date
		ORDER BY p.name
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project progress"})
		return
	}
	defer rows.Close()

	var projects []ProjectProgressResponse

	for rows.Next() {
		var project ProjectProgressResponse
		var totalTasks, doneTasks int
		var sprintName, startDate, endDate sql.NullString

		err := rows.Scan(
			&project.ID,
			&project.Name,
			&totalTasks,
			&doneTasks,
			&sprintName,
			&startDate,
			&endDate,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan project data"})
			return
		}

		project.TotalTasks = totalTasks
		project.DoneTasks = doneTasks
		
		// 進捗率計算（完了タスク数 ÷ 全タスク数 × 100）
		if totalTasks > 0 {
			project.Progress = (float64(doneTasks) / float64(totalTasks)) * 100
		} else {
			project.Progress = 0
		}

		if sprintName.Valid {
			project.SprintName = sprintName.String
		}
		if startDate.Valid {
			project.StartDate = startDate.String
		}
		if endDate.Valid {
			project.EndDate = endDate.String
		}

		projects = append(projects, project)
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// GetProjectProgress は特定のプロジェクトの進捗率を取得します
func (h *ProjectHandler) GetProjectProgress(c *gin.Context) {
	projectID := c.Param("id")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	query := `
		SELECT 
			p.id,
			p.name,
			COUNT(t.id) as total_tasks,
			SUM(CASE WHEN st.name = 'done' THEN 1 ELSE 0 END) as done_tasks,
			s.name as sprint_name,
			s.start_date,
			s.end_date
		FROM projects p
		LEFT JOIN tasks t ON p.id = t.project_id
		LEFT JOIN statuses st ON t.status_id = st.id
		LEFT JOIN sprint s ON t.sprint_id = s.id
		WHERE p.id = ?
		GROUP BY p.id, p.name, s.name, s.start_date, s.end_date
	`

	var project ProjectProgressResponse
	var totalTasks, doneTasks int
	var sprintName, startDate, endDate sql.NullString

	err := h.DB.QueryRow(query, projectID).Scan(
		&project.ID,
		&project.Name,
		&totalTasks,
		&doneTasks,
		&sprintName,
		&startDate,
		&endDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project progress"})
		return
	}

	project.TotalTasks = totalTasks
	project.DoneTasks = doneTasks
	
	// 進捗率計算
	if totalTasks > 0 {
		project.Progress = (float64(doneTasks) / float64(totalTasks)) * 100
	} else {
		project.Progress = 0
	}

	if sprintName.Valid {
		project.SprintName = sprintName.String
	}
	if startDate.Valid {
		project.StartDate = startDate.String
	}
	if endDate.Valid {
		project.EndDate = endDate.String
	}

	c.JSON(http.StatusOK, gin.H{"project": project})
}

// GetSprintProgress は特定のスプリントの進捗率を取得します
func (h *ProjectHandler) GetSprintProgress(c *gin.Context) {
	sprintID := c.Param("id")
	if sprintID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sprint ID is required"})
		return
	}

	query := `
		SELECT 
			s.id,
			s.name,
			COUNT(t.id) as total_tasks,
			SUM(CASE WHEN st.name = 'done' THEN 1 ELSE 0 END) as done_tasks,
			p.name as project_name,
			s.start_date,
			s.end_date
		FROM sprint s
		LEFT JOIN tasks t ON s.id = t.sprint_id
		LEFT JOIN statuses st ON t.status_id = st.id
		LEFT JOIN projects p ON s.project_id = p.id
		WHERE s.id = ?
		GROUP BY s.id, s.name, p.name, s.start_date, s.end_date
	`

	var sprint struct {
		ID          uint    `json:"id"`
		Name        string  `json:"name"`
		TotalTasks  int     `json:"total_tasks"`
		DoneTasks   int     `json:"done_tasks"`
		Progress    float64 `json:"progress"`
		ProjectName string  `json:"project_name"`
		StartDate   string  `json:"start_date"`
		EndDate     string  `json:"end_date"`
	}

	var totalTasks, doneTasks int
	var projectName, startDate, endDate string

	err := h.DB.QueryRow(query, sprintID).Scan(
		&sprint.ID,
		&sprint.Name,
		&totalTasks,
		&doneTasks,
		&projectName,
		&startDate,
		&endDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sprint not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sprint progress"})
		return
	}

	sprint.TotalTasks = totalTasks
	sprint.DoneTasks = doneTasks
	sprint.ProjectName = projectName
	sprint.StartDate = startDate
	sprint.EndDate = endDate
	
	// 進捗率計算
	if totalTasks > 0 {
		sprint.Progress = (float64(doneTasks) / float64(totalTasks)) * 100
	} else {
		sprint.Progress = 0
	}

	c.JSON(http.StatusOK, gin.H{"sprint": sprint})
}
