package service

import (
	"userMicroService/dbaccess"
	"userMicroService/model"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	var userCarrier model.User
	err := c.BindJSON(&userCarrier)
	if err != nil {
		log.Fatal("(RegisterUser) c.BindJSON", err)
	}

	query := `INSERT INTO user (first_name, last_name, email,password) VALUES (?, ?,?,?)`
	res, err := db.Exec(query, userCarrier.FirstName, userCarrier.LastName, userCarrier.Email, userCarrier.Password)
	if err != nil {
		log.Fatal("(RegisterUser) db.Exec", err)
	}
	userCarrier.ID, err = res.LastInsertId()
	if err != nil {
		log.Fatal("(CreateProduct) res.LastInsertId", err)
	}

	c.JSON(http.StatusOK, userCarrier)
}

func GetUsers(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	query := "SELECT * FROM user"
	res, err := db.Query(query)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetProducts) db.Query", err)
	}

	products := []model.User{}
	for res.Next() {
		var product model.User
		err := res.Scan(&product.ID, &product.FirstName, &product.LastName, &product.Password, &product.Email)
		if err != nil {
			log.Fatal("(GetProducts) res.Scan", err)
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}
