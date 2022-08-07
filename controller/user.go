package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kaushik13k/rest-GoApi/config"
	"github.com/kaushik13k/rest-GoApi/models"
)

func GetAllUser(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
	// c.String(200, "HELLO WORLD1234")
}

func GetSpecificUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Find(&user)
	c.JSON(200, &user)
	// c.String(200, "HELLO WORLD1234")
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
