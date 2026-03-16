package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"todo-api/models"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	DB *sql.DB
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {

	rows, err := h.DB.Query("SELECT id,title,description,status,user_id FROM tasks")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {

		var t models.Task

		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID)

		tasks = append(tasks, t)
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	err := h.DB.QueryRow(
		"INSERT INTO tasks(title,description,status,user_id) VALUES($1,$2,$3,$4) RETURNING id",
		task.Title,
		task.Description,
		task.Status,
		task.UserID,
	).Scan(&task.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	taskID, _ := strconv.Atoi(id)

	_, err := h.DB.Exec("DELETE FROM tasks WHERE id=$1", taskID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
