package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func setUpTest(t *testing.T) (*fiber.App, *MockUserService, *UserHandler) {
	app := fiber.New()
	mockUserService := new(MockUserService)
	handler := NewUserHandler(mockUserService)
	return app, mockUserService, handler
}

func TestUserHandler_Register(t *testing.T) {
	app, mockUserService, handler := setUpTest(t)
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
