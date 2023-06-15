package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"userMicroService/model"
	"userMicroService/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAnimal(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Create the request body
	var animal = model.Animal{
		UserID:      1,
		AnimalID:    123,
		Picture:     "https://example.com/picture.jpg",
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "2000-01-01",
		Gender:      "Male",
		Type:        "Dog",
		Breed:       "Labrador Retriever",
		Shelter:     "ABC Shelter",
		Address:     "123 Main St",
		PostalCode:  "12345",
		Bio:         "A friendly and playful dog.",
	}
	
	payload, err := json.Marshal(animal)
	if err != nil {
		t.Fatal("Failed to marshal JSON payload:", err)
	}
	request := httptest.NewRequest(http.MethodPost, "/create-animal", bytes.NewBuffer(payload))

	// Create a test response recorder
	recorder := httptest.NewRecorder()
	// Call the CreateAnimal handler function
	table := "animal_profile_test" // Provide the table name
	router.POST("/create-animal", func(c *gin.Context) {
		service.CreateAnimal(c, table)
	})
	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAnimal(t *testing.T) {
	// Set up the router
	router := gin.Default()
	
	request := httptest.NewRequest(http.MethodGet, "/getanimal", nil)

	// Create a test response recorder
	recorder := httptest.NewRecorder()
	// Call the CreateAnimal handler function
	table := "animal_profile_test" // Provide the table name
	router.GET("/getanimal", func(c *gin.Context) {
		service.GetAnimals(c, table)
	})
	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
}
