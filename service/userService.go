package service

import (
	"os"
	"time"
	"userMicroService/dbaccess"
	"userMicroService/model"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	var userCarrier model.User
	err := c.BindJSON(&userCarrier)
	if err != nil {
		log.Fatal("(RegisterUser) c.BindJSON", err)
	}

	query1 := `SELECT email FROM user WHERE email = ?`
	rows, err1 := db.Query(query1, userCarrier.Email)
	if err1 != nil {
		log.Fatal(err1)
	}

	// Process the result
	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			log.Fatal(err)
		}
		if email != "" {
			c.JSON(http.StatusOK, "Email already exists")
			return
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userCarrier.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO user (first_name, last_name, email,password) VALUES (?, ?,?,?)`
	res, err := db.Exec(query, userCarrier.FirstName, userCarrier.LastName, userCarrier.Email, hash)
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
		err := res.Scan(&product.ID, &product.FirstName, &product.LastName, &product.Email, &product.Password)
		if err != nil {
			log.Fatal("(GetProducts) res.Scan", err)
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func Login(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	var attept model.LoginAttept
	if c.BindJSON(&attept) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read the body",
		})
		return
	}

	query1 := `SELECT * FROM user WHERE email = ?`
	rows, err1 := db.Query(query1, attept.Email)
	if err1 != nil {
		log.Fatal(err1)
	}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
			log.Fatal(err)
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(attept.Password))

		if err != nil {
			c.JSON(http.StatusBadRequest, "Wrong email or password")
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":     user.ID,
			"expire": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to make token")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
