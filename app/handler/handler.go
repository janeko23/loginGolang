package handler

import (

	"login-go/app/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"

)
var client *redis.Client

func Simple(c *gin.Context)() {
	c.String(200, "Welcome to Go and Gin! in Simple!!")
}
var user = models.User{
	ID:            1,
	UserName: "username",
	Hash: 	"password",
}

func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.UserName != u.UserName || user.Hash != u.Hash {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	
	c.JSON(http.StatusOK, "Pumm loggeaste")
}

