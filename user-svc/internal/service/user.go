package service

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/infra/queue"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/repository"
)

type UserService interface {
	Register(ctx context.Context, input *dto.UserSignup) error
	Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error)
	ForgotPassword(ctx context.Context, email string) error
	SetPassword(ctx context.Context, token, password string) error
	CreateProfile(ctx context.Context, profile *dto.UserProfile) error
	GetProfile(ctx context.Context, id int) (*domain.User, error)
	Authenticate(ctx *fiber.Ctx) (*domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	producer       queue.Producer
}

func (u *userService) Register(ctx context.Context, input *dto.UserSignup) error {
	return nil
}

func (u *userService) Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error) {
	return nil, nil
}

func (u *userService) ForgotPassword(ctx context.Context, email string) error {
	return nil
}

func (u *userService) SetPassword(ctx context.Context, token, password string) error {
	return nil
}

func (u *userService) CreateProfile(ctx context.Context, profile *dto.UserProfile) error {
	return nil
}

func (u *userService) GetProfile(ctx context.Context, id int) (*domain.User, error) {
	return nil, nil
}

func (u *userService) Authenticate(ctx *fiber.Ctx) (*domain.User, error) {
	return nil, nil
}

func NewUserService(userRepository repository.UserRepository, producer queue.Producer) UserService {
	return &userService{
		userRepository: userRepository,
		producer:       producer,
	}
}
