package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) GenerateToken(userID uint, email string) (string, error) {
	args := m.Called(userID, email)
	return args.String(0), args.Error(1)
}

func (m *MockTokenService) VerifyToken(tokenString string) (*dto.AuthResponse, error) {
	args := m.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.AuthResponse), args.Error(1)
}

func (m *MockTokenService) AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		authResponse, err := m.VerifyToken(authHeader)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		ctx.Locals("userId", authResponse.UserId)
		ctx.Locals("user", authResponse)
		return ctx.Next()
	}
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(ctx context.Context, input *dto.UserSignup) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

func (m *MockUserService) Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error) {
	args := m.Called(ctx, input)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*domain.User), args.Error(1)
}

func (m *MockUserService) ForgotPassword(ctx context.Context, input *dto.ForgotPassword) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

func (m *MockUserService) SetPassword(ctx context.Context, input *dto.SetPassword) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

func (m *MockUserService) CreateProfile(ctx context.Context, profile *dto.UserProfile) error {
	args := m.Called(ctx, profile)
	return args.Error(0)
}

func (m *MockUserService) GetProfile(ctx context.Context, id uint) (*domain.User, error) {
	args := m.Called(ctx, id)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*domain.User), args.Error(1)
}

func (m *MockUserService) Authenticate(ctx *fiber.Ctx) (*domain.User, error) {
	args := m.Called(ctx)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*domain.User), args.Error(1)
}

func setUpTest(t *testing.T) (*fiber.App, *MockUserService, *MockTokenService, *UserHandler) {
	app := fiber.New()
	mockUserService := new(MockUserService)
	mockTokenService := new(MockTokenService)
	handler := NewUserHandler(mockUserService, mockTokenService)

	return app, mockUserService, mockTokenService, handler
}

func TestUserHandler_Register(t *testing.T) {
	app, mockUserService, _, handler := setUpTest(t)
	app.Post("/register", handler.Register)

	body := dto.UserSignup{
		Email:    "test@test.com",
		Password: "password123",
		Phone:    "+123456789012",
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	mockUserService.On("Register", mock.Anything, &body).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUserHandler_MissingField(t *testing.T) {
	app, mockUserService, _, handler := setUpTest(t)
	app.Post("/register", handler.Register)

	body := dto.UserSignup{
		Email:    "",
		Password: "",
		Phone:    "+123456789012",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	mockUserService.On("Register", mock.Anything, &body).Return(nil)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUserHandler_ServiceError(t *testing.T) {
	app, mockUserService, _, handler := setUpTest(t)
	app.Post("/register", handler.Register)

	body := dto.UserSignup{
		Email:    "test@test.com",
		Password: "password123",
		Phone:    "+123456789012",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	mockUserService.On("Register", mock.Anything, &body).Return(fiber.NewError(fiber.StatusInternalServerError, "service error"))
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestLogin_Success(t *testing.T) {
	app, mockUserService, mockTokenService, handler := setUpTest(t)
	app.Post("/login", handler.Login)
	body := dto.UserLogin{
		Email:    "test@test.com",
		Password: "password123",
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	user := &domain.User{
		ID:       uint(1),
		Email:    body.Email,
		Password: body.Password,
	}

	mockUserService.On("Login", mock.Anything, &body).Return(user, nil)
	mockToken := "mocked-jwt-token"
	mockTokenService.On("GenerateToken", user.ID, user.Email).Return(mockToken, nil)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestLogin_Unauthorized(t *testing.T) {
	app, mockUserService, _, handler := setUpTest(t)
	app.Post("/login", handler.Login)
	body := dto.UserLogin{
		Email:    "test@test.com",
		Password: "wrong_password",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	mockUserService.On("Login", mock.Anything, &body).Return(nil, fiber.NewError(fiber.StatusUnauthorized, "unauthorized error"))
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestForgotPassword_Success(t *testing.T) {
	app, mockUserService, _, handler := setUpTest(t)
	app.Post("/forgot-password", handler.ForgotPassword)

	body := dto.ForgotPassword{
		Email: "test@test.com",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	mockUserService.On("ForgotPassword", mock.Anything, &body).Return(nil)
	req := httptest.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSetPassword_Success(t *testing.T) {
	app, mockUserService, _, handler := setUpTest(t)
	app.Post("/set-password", handler.SetPassword)

	body := dto.SetPassword{
		Token:    "valid_token",
		Password: "password123",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	mockUserService.On("SetPassword", mock.Anything, &body).Return(nil)
	req := httptest.NewRequest(http.MethodPost, "/set-password", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateProfile_Success(t *testing.T) {
	app, mockUserService, mockTokenService, handler := setUpTest(t)

	app.Use(mockTokenService.AuthMiddleware())
	app.Post("/profile", handler.CreateProfile)

	firstName := "John"
	lastName := "Doe"
	phone := "+123456789012"

	body := dto.UserProfile{
		UserId:    1,
		FirstName: &firstName,
		LastName:  &lastName,
		Phone:     &phone,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	authResponse := &dto.AuthResponse{
		UserId: 1,
		Email:  "test@test.com",
		Exp:    float64(time.Now().Add(24 * time.Hour).Unix()),
	}
	mockTokenService.On("VerifyToken", "Bearer valid_token").Return(authResponse, nil)

	mockUserService.On("CreateProfile", mock.AnythingOfType("*fasthttp.RequestCtx"), &body).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/profile", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	mockUserService.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}

func TestCreateProfile_InvalidToken(t *testing.T) {
	app, _, mockTokenService, handler := setUpTest(t)
	app.Use(mockTokenService.AuthMiddleware())
	app.Post("/profile", handler.CreateProfile)

	firstName := "John"
	lastName := "Doe"
	phone := "+123456789012"
	body := dto.UserProfile{
		UserId:    1,
		FirstName: &firstName,
		LastName:  &lastName,
		Phone:     &phone,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	mockTokenService.On("VerifyToken", "Bearer invalid_token").Return(nil, errors.New("invalid token"))

	req := httptest.NewRequest(http.MethodPost, "/profile", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer invalid_token")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	mockTokenService.AssertExpectations(t)
}

func TestGetProfile_Success(t *testing.T) {
	app, mockUserService, mockTokenService, handler := setUpTest(t)
	
	app.Use(mockTokenService.AuthMiddleware())
	app.Get("/profile", handler.GetProfile)

	user := &domain.User{
		ID:    uint(1),
		Email: "test@test.com",
	}

	authResponse := &dto.AuthResponse{
		UserId: 1,
		Email:  "test@test.com",
		Exp:    float64(time.Now().Add(24 * time.Hour).Unix()),
	}
	mockTokenService.On("VerifyToken", "Bearer valid_token").Return(authResponse, nil)

	mockUserService.On("GetProfile", mock.AnythingOfType("*fasthttp.RequestCtx"), uint(1)).Return(user, nil)

	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockUserService.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}

func TestAuth_Success(t *testing.T) {
	app, mockUserService, _, handler := setUpTest(t)
	app.Get("/auth", handler.Authentication)
	user := &domain.User{
		ID:    uint(1),
		Email: "test@test.com",
	}

	mockUserService.On("Authenticate", mock.Anything).Return(user, nil)
	req := httptest.NewRequest(http.MethodGet, "/auth", nil)
	req.Header.Set("Authorization", "Bearer valid_token")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestMe_Success(t *testing.T) {
	app, mockUserService, mockTokenService, handler := setUpTest(t)

	app.Use(mockTokenService.AuthMiddleware())
	app.Get("/me", handler.Me)

	user := &domain.User{
		ID:    uint(1),
		Email: "test@test.com",
	}

	authResponse := &dto.AuthResponse{
		UserId: 1,
		Email:  "test@test.com",
		Exp:    float64(time.Now().Add(24 * time.Hour).Unix()),
	}
	mockTokenService.On("VerifyToken", "Bearer valid_token").Return(authResponse, nil)

	mockUserService.On("GetProfile", mock.AnythingOfType("*fasthttp.RequestCtx"), uint(1)).Return(user, nil)

	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockUserService.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}

func TestMe_InvalidToken(t *testing.T) {
	app, mockUserService, mockTokenService, handler := setUpTest(t)
	app.Use(mockTokenService.AuthMiddleware())
	app.Get("/me", handler.Me)

	mockTokenService.On("VerifyToken", "Bearer invalid_token").Return(nil, errors.New("invalid token"))

	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	mockTokenService.AssertExpectations(t)
	mockUserService.AssertExpectations(t)
}
