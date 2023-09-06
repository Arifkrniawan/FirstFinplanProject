package hobby_controller

import (
	"net/http"

	"github.com/Arifkrniawan/FirstFinplanProject/database"
	"github.com/Arifkrniawan/FirstFinplanProject/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var hobbies []models.Hobby

	database.DB.Find(&hobbies)
	c.JSON(http.StatusOK, gin.H{"hobbies": hobbies})
}
