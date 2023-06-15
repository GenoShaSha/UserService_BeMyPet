package main

import (
	"userMicroService/kafkaaccess"
	"userMicroService/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	tableAnimal := "animal_profile" // Provide the table name

	router.POST("/register", service.RegisterUser)
	router.POST("/login", service.Login)
	router.GET("/users", service.GetUsers)
	router.POST("/createAnimal", func(c *gin.Context) {
		service.CreateAnimal(c, tableAnimal)
	})
	router.GET("/animals", func(c *gin.Context) {
		service.GetAnimals(c, tableAnimal)
	})

	// router.GET("/products", GetProducts)
	// router.GET("/products/:productId", GetSingleProduct)
	// router.PUT("/products/:productId", UpdateProduct)
	// router.DELETE("/products/:productId", DeleteProduct)

	// Run the router
	router.Run(":3000")

	kafkaaccess.ConnectAndWriteMessage()
	kafkaaccess.ConnectAndConsumeMessage()

	//Testing Sonar qube

}
