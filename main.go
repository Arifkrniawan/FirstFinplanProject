package main

import (
	"github.com/Arifkrniawan/FirstFinplanProject/database"
	"github.com/Arifkrniawan/FirstFinplanProject/repository/hobby_controller"
	"github.com/Arifkrniawan/FirstFinplanProject/repository/map_user_hobby_controller"
	"github.com/Arifkrniawan/FirstFinplanProject/repository/result_controller"
	"github.com/Arifkrniawan/FirstFinplanProject/repository/user_controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.GetConnection()

	r.GET("/api/users", user_controller.Index)                    //done
	r.GET("/api/hobbies", hobby_controller.Index)                 //done
	r.GET("/api/map_user_hobby", map_user_hobby_controller.Index) //done
	r.GET("/api/users/detail", user_controller.Show)
	r.GET("/api/result", result_controller.Index)
	r.POST("/api/users", user_controller.Create)                       //done
	r.POST("/api/map_user_hobby", map_user_hobby_controller.Create)    //done
	r.PUT("/api/users/:id", user_controller.Update)                    //done
	r.PUT("/api/map_user_hobby/:id", map_user_hobby_controller.Update) //done

	r.Run()
}
