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
	var results models.Result

	id := c.Query("id")
	fmt.Println("PAMRAM ID: ", id)

	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "CANT FOUND DATA"})
		return
	}

	if err := database.DB.Table("hobbies").Select("hobbies.*").
		Joins("INNER JOIN map_user_hobbies ON hobbies.id = map_user_hobbies.id_hobby").
		Where("map_user_hobbies.id_user = ?", id).Scan(&results.Hobby).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "CANT FOUND DATA"})
		return
	}

	results.Id = users.Id
	results.Name = users.Name
	results.Gender = users.Gender
	results.Status = users.Status

	c.JSON(http.StatusOK, gin.H{"Data": results})
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
