package controllers

import (
	"bytes"
	"encoding/json"
	"gaia-api/application/interfaces"
	"gaia-api/domain/entities"
	"gaia-api/domain/mocks"
	"gaia-api/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Create a mock UserRepo instance
func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock user repository
	mockUserRepo := mocks.NewMockUserInterface(ctrl)

	// Create an instance of the AuthService using the mock user repository
	authService := services.NewAuthService(mockUserRepo)

	// Create a mock ReturnAPIData
	returnAPIData := &interfaces.ReturnAPIData{}

	// Create a Server instance with the mock dependencies
	server := &Server{
		authService:   authService,
		returnAPIData: returnAPIData,
	}

	// Create a mock context
	w := httptest.NewRecorder()
	/*invalidUserJSON := `{
		"Id": "invalid",
		"Email": "invalid_email",
		"Username": "example",
		"Password": "password123",
		"Pseudo": "user123",
		"Birthdate": "1990-01-01",
		"Weight": 70.5,
		"Height": 180,
		"Gender": 0,
		"Sportiveness": 1,
		"BasalMetabolism": 1500
	}`*/
	validUserJSON := `{
		"Id": "prout",
		"Email": "henry.sargerson@hotmail.fr",
		"Username": "example",
		"Password": "password123",
		"Pseudo": "user123",
		"Birthdate": "1990-01-01",
		"Weight": 70.5,
		"Height": 180,
		"Gender": 0,
		"Sportiveness": 1,
		"BasalMetabolism": 1500
	}`
	user := entities.User{}
	//Bind the JSON data to the entities.User struct
	err := json.Unmarshal([]byte(validUserJSON), &user)
	if err != nil {
		t.Errorf("Error during the register test: %s", err)
	}
	jsonData, err := json.Marshal(user)

	if err != nil {
		t.Errorf("Error during the register test: %s", err)
	}
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("Error during the register test: %s", err)
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set the expectation for the Register method
	mockUserRepo.EXPECT().Register(gomock.Any()).Return(true, nil)

	// Call the test function
	server.register(c)

	// Verify the expected response code
	if w.Code != http.StatusOK {
		t.Errorf("Expected response code for the register test: %d, obtained: %d", http.StatusOK, w.Code)
	}
}
