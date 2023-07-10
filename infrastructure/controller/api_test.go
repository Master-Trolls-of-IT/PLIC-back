package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"gaia-api/application/interface"
	"gaia-api/domain/entity"
	"gaia-api/domain/mock"
	"gaia-api/domain/service"
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
	mockUserRepo := mock.NewMockUserInterface(ctrl)

	// Create an instance of the AuthService using the mock user repository
	authService := service.NewAuthService(mockUserRepo)

	// Create a mock ReturnAPIData
	returnAPIData := &interfaces.ReturnAPIData{}

	// Create a Server instance with the mock dependencies
	server := &Server{
		authService:   authService,
		returnAPIData: returnAPIData,
	}

	// Create a mock context
	w := httptest.NewRecorder()
	validUserJSON := `{
		"Id": 1,
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
	invalidUserJSON := `{
		"Id": 1,
		"Email": 2334,
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

	user := entity.User{}
	invalidUser := entity.User{}
	//Bind the JSON data to the entities.User struct
	err := json.Unmarshal([]byte(validUserJSON), &user)
	if err != nil {
		t.Errorf("Error during the register test: %s", err)
	}
	err = json.Unmarshal([]byte(invalidUserJSON), &invalidUser)
	if err == nil {
		t.Errorf("Error during the register test: Invalid user data should not be unmarshalled")
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

func TestLoginUserWithInvalidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Create a mock user repository
	mockUserRepo := mock.NewMockUserInterface(ctrl)
	//Create expectations on the mock user and password
	//mockUserRepo.EXPECT().GetUserByEmail(gomock.Any()).Return(entities.User{}, nil) // Create an instance of the AuthService using the mock user repository
	mockUserRepo.EXPECT().CheckLogin(gomock.Any()).Return(false, errors.New("invalid login credentials"))
	//mockUserRepo.EXPECT().CheckLogin(gomock.Any()).Return(error(), nil)
	authService := service.NewAuthService(mockUserRepo)

	// Create a mock ReturnAPIData
	returnAPIData := &interfaces.ReturnAPIData{}

	// Create a Server instance with the mock dependencies
	server := &Server{
		authService:   authService,
		returnAPIData: returnAPIData,
	}
	// Create a mock context
	w := httptest.NewRecorder()
	invalidUserJSON := `{
		"Email": "invalid@example.com",
		"Password": "invalidpassword"
	}`
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(invalidUserJSON)))
	if err != nil {
		t.Errorf("Error during the login test: %s", err)
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the test function and verify the expected response code (e.g., unauthorized)
	server.login(c)

	// Verify the expected response code (e.g., unauthorized)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected response code for the login test: %d, obtained: %d", http.StatusUnauthorized, w.Code)
	}
}

func TestLoginUserSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Create a mock user repository
	mockUserRepo := mock.NewMockUserInterface(ctrl)
	//Create expectations on the mock user and password
	mockUserRepo.EXPECT().GetUserByEmail(gomock.Any()).Return(entity.User{}, nil) // Create an instance of the AuthService using the mock user repository
	// the Check Login method should return nil error
	mockUserRepo.EXPECT().CheckLogin(gomock.Any()).Return(true, nil) //mockUserRepo.EXPECT().CheckLogin(gomock.Any()).Return(error(), nil)
	// Create an instance of the AuthService using the mock user repository
	authService := service.NewAuthService(mockUserRepo)

	// Create a mock ReturnAPIData
	returnAPIData := &interfaces.ReturnAPIData{}

	// Create a Server instance with the mock dependencies
	server := &Server{
		authService:   authService,
		returnAPIData: returnAPIData,
	}

	// Create a mock context
	w := httptest.NewRecorder()
	validUserJSON := `{
		"Email": "valid@example.com",
		"Password": "validpassword"
	}`
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(validUserJSON)))
	if err != nil {
		t.Errorf("Error during the login test: %s", err)
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the test function
	server.login(c)

	// Verify the expected response code (e.g., accepted)
	if w.Code != http.StatusAccepted {
		t.Errorf("Expected response code for the login test: %d, obtained: %d", http.StatusAccepted, w.Code)
	}
}
