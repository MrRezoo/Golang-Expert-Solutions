package main

// user.go
import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = make(map[int]User)
var idCounter = 1

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = idCounter
	idCounter++
	users[user.ID] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/users/"):])
	user, exists := users[id]
	if !exists {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/users", CreateUser)
	http.HandleFunc("/users/", GetUser)
	http.ListenAndServe(":8080", nil)
}
