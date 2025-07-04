package main

import (
	"go-gin-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/habits", handlers.AddHabit)

	r.GET("/habits", handlers.ListHabits)
	r.GET("/habits/:id", handlers.GetHabit)

	r.PUT("/habits/:id", handlers.UpdateHabit)

	r.DELETE("/habits/:id", handlers.DeleteHabit)

	r.Run() // listens on :8080 by default
}
