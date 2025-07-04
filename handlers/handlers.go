package handlers

import (
	"go-gin-api/data"
	"go-gin-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddHabit(c *gin.Context) {
	var newItem models.Habit

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem.ID = data.IDCounter
	data.HabitDB = append(data.HabitDB, newItem)
	data.IDCounter++

	c.JSON(http.StatusCreated, newItem)
}

func GetHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, item := range data.HabitDB {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func ListHabits(c *gin.Context) {
	c.JSON(http.StatusOK, data.HabitDB)
}

func UpdateHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedItem models.Habit
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, item := range data.HabitDB {
		if item.ID == id {
			data.HabitDB[index].Name = updatedItem.Name
			data.HabitDB[index].Category = updatedItem.Category
			c.JSON(http.StatusOK, data.HabitDB[index])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func DeleteHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, item := range data.HabitDB {
		if item.ID == id {
			data.HabitDB = append(data.HabitDB[:index], data.HabitDB[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}
