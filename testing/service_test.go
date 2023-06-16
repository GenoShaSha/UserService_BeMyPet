package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"userMicroService/model"
	"userMicroService/service"
	"time"
	"fmt"
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



func TestRegisterUser(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Create a unique email with current date and time
	email := fmt.Sprintf("test-%s@example.com", time.Now().Format("20060102150405"))

	// Create the request body
	var user = model.User{
		Email:    email,
		Password: "password123",
	}

	payload, err := json.Marshal(user)
	if err != nil {
		t.Fatal("Failed to marshal JSON payload:", err)
	}
	request := httptest.NewRequest(http.MethodPost, "/register-user", bytes.NewBuffer(payload))

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the RegisterUser handler function
	table := "user_test" // Provide the table name
	router.POST("/register-user", func(c *gin.Context) {
		service.RegisterUser(c, table)
	})

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Assert the response body
	var responseUser model.User
	err = json.Unmarshal(recorder.Body.Bytes(), &responseUser)
	if err != nil {
		t.Fatal("Failed to unmarshal response body:", err)
	}

	// Assert the user properties
	assert.Equal(t, user.Email, responseUser.Email)
	// Assert other properties as needed
}

func TestLoginUser(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Create the request body
	credentials := map[string]string{
		"email":    "test-20230616165330@example.com",
		"password": "password123",
	}

	payload, err := json.Marshal(credentials)
	if err != nil {
		t.Fatal("Failed to marshal JSON payload:", err)
	}
	request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the Login handler function
	table := "user_test" // Provide the table name
	router.POST("/login", func(c *gin.Context) {
		service.Login(c, table)
	})

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Assert the response body
	var responseBody gin.H
	err = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal("Failed to unmarshal response body:", err)
	}

	// Assert the token exists in the response body
	token, ok := responseBody["token"].(string)
	assert.True(t, ok, "Token not found in response body")
	assert.NotEmpty(t, token, "Token is empty")

	// Perform additional assertions as needed
}
