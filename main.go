package main

import (
	"userMicroService/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/register", service.CreateUser)

	router.Run(":5000")
}
