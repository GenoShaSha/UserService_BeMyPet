package service

import (
	"fmt"
	"os"
	"time"
	"userMicroService/dbaccess"
	"userMicroService/model"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var animalsCache *cache.Cache

func init() {
	// Initialize the cache with a default expiration time of 5 minutes
	animalsCache = cache.New(5*time.Minute, 10*time.Minute)
}

//Create Animal
func CreateAnimal(c *gin.Context) {
	db := dbaccess.ConnectToDb()
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("CreateAnimalLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("Error opening file:", err)
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Set the output of the logger to the file
	logger.SetOutput(file)

	// Get client IP
	clientIP := c.ClientIP()

	var animalCarrier model.Animal
	err = c.BindJSON(&animalCarrier)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"IP":     clientIP,
			"Status": "Error",
		}).Error("(CreateAnimal) c.BindJSON", err)
		log.Fatal("(CreateAnimal) c.BindJSON", err)
	}

	query1 := `SELECT first_name, last_name FROM animal_profile WHERE first_name = ? AND last_name = ?`
	rows, err := db.Query(query1, animalCarrier.FirstName, animalCarrier.LastName)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"IP":     clientIP,
			"Status": "Error",
		}).Error(err)
		log.Fatal(err)
	}

	// Process the result
	for rows.Next() {
		var first_name string
		var last_name string
		err := rows.Scan(&first_name, &last_name)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"IP":     clientIP,
				"Status": "Error",
			}).Error(err)
			log.Fatal(err)
		}
		if first_name != "" && last_name != "" {
			logger.WithFields(logrus.Fields{
				"IP":     clientIP,
				"Status": "Error",
			}).Error("The profile is already exist")
			message := []byte("The profile is already exist")
			n, err := file.Write(message)
			if err != nil {
				logger.WithFields(logrus.Fields{
					"IP":     clientIP,
					"Status": "Error",
				}).Error("Error writing to file:", err)
				fmt.Println("Error writing to file:", err)
				return
			}
			logger.WithFields(logrus.Fields{
				"IP":           clientIP,
				"Status":       "Error",
				"BytesWritten": n,
			}).Info("Bytes written to file")
			fmt.Printf("Bytes written: %d\n", n)
			c.JSON(http.StatusOK, "The profile is already exist")
			return
		}
	}

	query := `INSERT INTO animal_profile (user_id,picture, first_name, last_name, date_of_birth, gender, type, breed, shelter, address, postal_code, bio) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := db.Exec(query, animalCarrier.UserID, animalCarrier.Picture, animalCarrier.FirstName, animalCarrier.LastName, animalCarrier.DateOfBirth, animalCarrier.Gender, animalCarrier.Type, animalCarrier.Breed, animalCarrier.Shelter, animalCarrier.Address, animalCarrier.PostalCode, animalCarrier.Bio)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"IP":     clientIP,
			"Status": "Error",
		}).Error("(CreateAnimal) db.Exec", err)
		log.Fatal("(CreateAnimal) db.Exec", err)
	}
	animalCarrier.AnimalID, err = res.LastInsertId()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"IP":     clientIP,
			"Status": "Error",
			"Query":  query,
		}).Error("(CreateAnimal) res.LastInsertId", err)
		log.Fatal("(CreateAnimal) res.LastInsertId", err)
	}

	// Log the successful registration with client IP
	logger.WithFields(logrus.Fields{
		"IP":     clientIP,
		"Status": "Success",
	}).Info("Animal created successfully")

	animalsCache.Flush()

	c.JSON(http.StatusOK, animalCarrier)
}

//Get Animals
func GetAnimals(c *gin.Context) {
	type AnimalsResponse struct {
		Animals []model.Animal `json:"animals"`
	}

	// Check if the users are already cached
	if cachedAnimals, found := animalsCache.Get("animals"); found {
		// If the users are cached, return the cached data
		c.JSON(http.StatusOK, AnimalsResponse{Animals: cachedAnimals.([]model.Animal)})
		return
	}

	db := dbaccess.ConnectToDb()
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("GetAnimalsLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("Error opening file:", err)
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Set the output of the logger to the file
	logger.SetOutput(file)

	// Get client IP
	clientIP := c.ClientIP()

	query := "SELECT * FROM animal_profile"
	res, err := db.Query(query)
	defer res.Close()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"IP":     clientIP,
			"Status": "Error",
		}).Error("(GetProducts) db.Query", err)
		log.Fatal("(GetProducts) db.Query", err)
	}

	animals := []model.Animal{}
	for res.Next() {
		var animal model.Animal

		err := res.Scan(&animal.UserID, &animal.AnimalID, &animal.Picture, &animal.FirstName, &animal.LastName, &animal.DateOfBirth, &animal.Gender, &animal.Type, &animal.Breed, &animal.Shelter, &animal.Address, &animal.PostalCode, &animal.Bio)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"IP":     clientIP,
				"Status": "Error",
			}).Error("(GetAnimals) res.Scan", err)
			log.Fatal("(GetAnimals) res.Scan", err)
		}
		animals = append(animals, animal)
	}

	// Store the retrieved users in the cache
	animalsCache.Set("animals", animals, cache.DefaultExpiration)

	// Log the result to the file with client IP
	logger.WithFields(logrus.Fields{
		"IP":     clientIP,
		"Status": "Success",
	}).Info("Animals retrieved successfully")

	// Wrap the users array within an object
	response := AnimalsResponse{Animals: animals}

	c.JSON(http.StatusOK, response)
}
