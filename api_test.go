package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"server/middleware"
	"strings"
	"testing"
)

func TestGetAllTask(t *testing.T) {
	req, err := http.NewRequest("GET", "/task", nil)
	if err != nil {
			t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware.GetAllTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := ""
	if  rr.Body.String() == expected {
			t.Errorf("Not valid response %v, %v", rr.Body.String(),expected)
	}
}

func TestCreateTask(t *testing.T) {

	var jsonStr = []byte(`{"task": "Deneme task"}`)

	req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware.CreateTask)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `"task": "Deneme task"`
	if  strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


