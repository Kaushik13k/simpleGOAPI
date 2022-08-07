package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaushik13k/rest-GoApi/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/getalluser", controller.GetAllUser)
	router.POST("/createuser", controller.CreateUser)
	router.DELETE("/deleteuser/:id", controller.DeleteUser)
	router.PUT("/updateuser/:id", controller.UpdateUser)
	router.GET("/getspecificuser/:id", controller.GetSpecificUser)
}
