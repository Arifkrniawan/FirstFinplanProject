package result_controller

import (
	"net/http"

	"github.com/Arifkrniawan/FirstFinplanProject/database"
	"github.com/Arifkrniawan/FirstFinplanProject/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var result []models.Result

	database.DB.Find(&result)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
