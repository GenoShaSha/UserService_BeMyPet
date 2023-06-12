package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"userMicroService/model"
	"userMicroService/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAnimal(t *testing.T) {
	// Set up the router
	router := gin.Default()
	router.POST("/animal", service.CreateAnimal)

	// Create a mock DB connection and get the mock instance
	db, mock, _ := sqlmock.New()

	// Inject the mock DB connection into the service

	service.CheckDB = false

	service.DBTest = db

	// Create the request body
	animal := model.Animal{
		UserID:      1,
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "2000-01-01",
		Gender:      "Male",
		Type:        "Dog",
		Breed:       "Labrador",
		Shelter:     "XYZ Shelter",
		Address:     "123 Main St",
		PostalCode:  "12345",
		Bio:         "Lorem ipsum dolor sit amet",
	}
	body, _ := json.Marshal(animal)

	// Set up the request
	req, _ := http.NewRequest("POST", "/animal", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Set up the mock database expectations
	mock.ExpectQuery(`SELECT first_name, last_name FROM animal_profile`).WithArgs(animal.FirstName, animal.LastName).WillReturnRows(sqlmock.NewRows([]string{"first_name", "last_name"}))
	mock.ExpectExec(`INSERT INTO animal_profile`).WillReturnResult(sqlmock.NewResult(1, 1))

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verify the response
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify the expected response body
	expectedResponse := model.Animal{
		UserID:      1,
		AnimalID:    1,
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "2000-01-01",
		Gender:      "Male",
		Type:        "Dog",
		Breed:       "Labrador",
		Shelter:     "XYZ Shelter",
		Address:     "123 Main St",
		PostalCode:  "12345",
		Bio:         "Lorem ipsum dolor sit amet",
	}
	actualResponse := model.Animal{}
	_ = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.Equal(t, expectedResponse, actualResponse)

	// Verify the mock expectations
	assert.Nil(t, mock.ExpectationsWereMet())
}
