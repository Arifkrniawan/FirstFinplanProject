package user_controller

import (
	"fmt"
	"net/http"

	"github.com/Arifkrniawan/FirstFinplanProject/database"
	"github.com/Arifkrniawan/FirstFinplanProject/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Create(c *gin.Context) {
	var users models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})

}

func Show(c *gin.Context) {
	var users models.User
	var result models.Result

	id := c.Param("id")
	fmt.Println("PAMRAM ID: ", id)

	if err := database.DB.Table("map_user_hobbies").Select("users.id, users.name, users.gender, users.status, hobbies.id, hobbies.name, hobbies.level").
		Joins("JOIN users on users.id=map_user_hobbies.id_user").
		Joins("JOIN hobbies on hobbies.id = map_user_hobbies.id_hobby").
		Where("users.id=?", id).Scan(&result).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "CANT FOUND DATA"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": users})
}

func Update(c *gin.Context) {
	var users models.User

	id := c.Param("id")

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&users).Where("id=?", id).Updates(&users).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat merubah data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "BERHASIL MENAMBAHKAN DATA"})
}
