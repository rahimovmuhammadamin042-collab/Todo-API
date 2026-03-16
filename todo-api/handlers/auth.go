package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todo-api/models"
	"todo-api/utils"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	err := h.DB.QueryRow(
		"INSERT INTO users(username,password) VALUES($1,$2) RETURNING id",
		user.Username,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	var dbUser models.User

	err := h.DB.QueryRow(
		"SELECT id,password FROM users WHERE username=$1",
		user.Username,
	).Scan(&dbUser.ID, &dbUser.Password)

	if err != nil || dbUser.Password != user.Password {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	token, _ := utils.GenerateJWT(dbUser.ID)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
