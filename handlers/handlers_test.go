package handlers

import (
	"bytes"
	"encoding/json"
	"go-gin-api/data"
	"go-gin-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAddHabit(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Reset data
	data.HabitDB = []models.Habit{}
	data.IDCounter = 1

	router := gin.New()
	router.POST("/habits", AddHabit)

	habit := models.Habit{Name: "Exercise", Category: "Health"}
	reqBody, _ := json.Marshal(habit)

	req, _ := http.NewRequest("POST", "/habits", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	if len(data.HabitDB) != 1 {
		t.Errorf("Expected 1 habit in DB, got %d", len(data.HabitDB))
	}
}
