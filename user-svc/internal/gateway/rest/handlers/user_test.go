package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/stretchr/testify/mock"
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
