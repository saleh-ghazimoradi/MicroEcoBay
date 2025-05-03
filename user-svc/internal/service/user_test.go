package service

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

type MockProducer struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*domain.User), args.Error(1)
}

func (m *MockUserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByResetToken(ctx context.Context, token string) (*domain.User, error) {
	args := m.Called(ctx, token)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindUserById(ctx context.Context, id uint) (*domain.User, error) {
	args := m.Called(ctx, id)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*domain.User), args.Error(1)
}

func (m *MockProducer) PublishMessage(ctx context.Context, key, value []byte) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *MockProducer) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestRegister_Success(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	input := dto.UserSignup{
		Email:    "test@test.com",
		Password: "password",
		Phone:    "123456789",
	}

	userRepository.On("CreateUser", ctx, mock.MatchedBy(func(user *domain.User) bool {
		assert.Equal(t, input.Email, user.Email)
		assert.Len(t, user.Password, 60)
		assert.Equal(t, input.Phone, user.Phone)
		return true
	})).Return(nil)

	err := uService.Register(ctx, &input)
	assert.NoError(t, err)
	userRepository.AssertExpectations(t)
}
