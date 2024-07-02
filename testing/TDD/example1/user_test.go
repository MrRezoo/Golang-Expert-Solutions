package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := User{Name: "John Doe", Email: "john@example.com"}
	jsonUser, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdUser User
	json.NewDecoder(rr.Body).Decode(&createdUser)

	if createdUser.Name != user.Name {
		t.Errorf("handler returned unexpected body: got %v want %v", createdUser.Name, user.Name)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var user User
	json.NewDecoder(rr.Body).Decode(&user)

	if user.ID != 1 {
		t.Errorf("handler returned unexpected body: got %v want %v", user.ID, 1)
	}
}
