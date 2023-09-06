package map_user_hobby_controller

import (
	"net/http"

	"github.com/Arifkrniawan/FirstFinplanProject/database"
	"github.com/Arifkrniawan/FirstFinplanProject/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var map_user_hobbies []models.MapUserHobby

	database.DB.Find(&map_user_hobbies)
	c.JSON(http.StatusOK, gin.H{"map_user_hobbies": map_user_hobbies})
}

func Create(c *gin.Context) {
	var map_user_hobby models.MapUserHobby

	if err := c.ShouldBindJSON(&map_user_hobby); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&map_user_hobby)
	c.JSON(http.StatusOK, gin.H{"map_user_hobby": map_user_hobby})

}

func Update(c *gin.Context) {
	var map_user_hobby models.MapUserHobby

	id := c.Param("id")

	if err := c.ShouldBindJSON(&map_user_hobby); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&map_user_hobby).Where("id=?", id).Updates(&map_user_hobby).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat merubah data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "BERHASIL MENAMBAHKAN DATA"})
}
