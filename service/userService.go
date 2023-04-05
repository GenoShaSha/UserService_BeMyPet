package service

import (
	"userMicroService/dbaccess"
	"userMicroService/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	ctx, _, usercollection := dbaccess.ConnectToDb()

	var userCarrier model.User
	c.BindJSON(&userCarrier)

	_, err := usercollection.InsertOne(ctx, userCarrier)
	if err != nil {
		panic(err)
	}
}
