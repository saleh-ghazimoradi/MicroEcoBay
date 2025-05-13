package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/customErr"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
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

func TestRegister_DuplicateUser(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	input := dto.UserSignup{
		Email:    "test@test.com",
		Password: "password",
		Phone:    "123456789",
	}

	userRepository.On("CreateUser", ctx, mock.Anything).Return(customErr.ErrDuplicateUser)
	err := uService.Register(ctx, &input)
	assert.ErrorIs(t, err, customErr.ErrDuplicateUser)
	userRepository.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	input := dto.UserLogin{
		Email:    "test@test.com",
		Password: "password",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := &domain.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	userRepository.On("FindUserByEmail", ctx, user.Email).Return(user, nil)
	result, err := uService.Login(ctx, &input)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Password, result.Password)
	userRepository.AssertExpectations(t)
}

func TestLogin_InvalidPassword(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	input := dto.UserLogin{
		Email:    "test@test.com",
		Password: "wrong_password",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("right_password"), bcrypt.DefaultCost)
	user := &domain.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	userRepository.On("FindUserByEmail", ctx, user.Email).Return(user, nil)
	result, err := uService.Login(ctx, &input)
	assert.Nil(t, result)
	assert.Error(t, err, "invalid email or password")
}

func TestForgotPassword_UserNotFound(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	input := &dto.ForgotPassword{
		Email: "test@test.com",
	}

	userRepository.On("FindUserByEmail", ctx, input.Email).Return(nil, customErr.ErrUserNotFound)
	err := uService.ForgotPassword(ctx, input)
	assert.Error(t, err, customErr.ErrUserNotFound)

}

func TestSetPassword_InvalidToken(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	input := &dto.SetPassword{
		Token:    "invalid_token",
		Password: "new_password",
	}

	userRepository.On("FindUserByResetToken", ctx, input.Token).Return((*domain.User)(nil), errors.New("not found"))
	err := uService.SetPassword(ctx, input)
	assert.EqualError(t, err, "invalid or expired token")
	userRepository.AssertExpectations(t)
}

func TestSetPassword_Success(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	input := &dto.SetPassword{
		Token:    "valid_token",
		Password: "new_password",
	}

	userRepository.On("FindUserByResetToken", ctx, input.Token).Return(&domain.User{
		ID:    1,
		Email: "test@test.com",
	}, nil)

	userRepository.On("SaveUser", ctx, mock.MatchedBy(func(user *domain.User) bool {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
		return user.Email == "test@test.com" && err == nil
	})).Return(nil)

	err := uService.SetPassword(ctx, input)
	assert.NoError(t, err)
	userRepository.AssertExpectations(t)
}

func TestGetProfile_Success(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	userId := uint(1)
	user := &domain.User{
		ID: userId,
	}

	userRepository.On("FindUserById", ctx, userId).Return(user, nil)
	profile, err := uService.GetProfile(ctx, userId)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, profile.ID)
	userRepository.AssertExpectations(t)
}

func TestGetProfile_NotFound(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	userId := uint(1)
	userRepository.On("FindUserById", ctx, userId).Return(nil, customErr.ErrUserNotFound)
	profile, err := uService.GetProfile(ctx, userId)
	assert.Nil(t, profile)
	assert.Error(t, err, "profile not found")
	userRepository.AssertExpectations(t)
}

func TestAuthenticate_Success(t *testing.T) {
	userRepository := new(MockUserRepository)
	producer := new(MockProducer)
	uService := NewUserService(userRepository, producer)
	ctx := context.Background()

	userId := uint(1)
	user := &domain.User{
		ID: userId,
	}

	userRepository.On("FindUserById", ctx, userId).Return(user, nil)
	profile, err := uService.GetProfile(ctx, userId)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, profile.ID)
	userRepository.AssertExpectations(t)
}
