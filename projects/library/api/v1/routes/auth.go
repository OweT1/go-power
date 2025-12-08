package routes

import (
	"encoding/json"
	"fmt"
	"library/db"
	"library/utils"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Repo *db.Repository
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user utils.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPass)
	result := h.Repo.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "User already exists!", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("User %s successfully created!", user.Username)})
}


func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var body, user utils.User
	json.NewDecoder(r.Body).Decode(&body)

	result := h.Repo.DB.Where("username = ?", body.Username).First(&user)
	if result.Error != nil {
		http.Error(w, "User not found!", http.StatusNotFound)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "Invalid Password!", http.StatusUnauthorized)
		return
	}

	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	JWT_SECRET := os.Getenv("JWT_SECRET")
	tokenString, err := jwt_token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		http.Error(w, "Error Creating JWT", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("jwt_token.txt", []byte(tokenString), 0666)
	if err != nil {
		http.Error(w, "Error writing JWT", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Logged in successfully!"})
}
