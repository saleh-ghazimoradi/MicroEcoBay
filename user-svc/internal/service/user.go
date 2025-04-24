package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/infra/queue"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/customErr"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/repository"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, input *dto.UserSignup) error
	Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error)
	ForgotPassword(ctx context.Context, email string) error
	SetPassword(ctx context.Context, token, password string) error
	CreateProfile(ctx context.Context, profile *dto.UserProfile) error
	GetProfile(ctx context.Context, id uint) (*domain.User, error)
	Authenticate(ctx *fiber.Ctx) (*domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	producer       queue.Producer
}

func (u *userService) findUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return u.userRepository.FindUserByEmail(ctx, email)
}

func (u *userService) findUserById(ctx context.Context, id uint) (*domain.User, error) {
	return u.userRepository.FindUserById(ctx, id)
}

func (u *userService) Register(ctx context.Context, input *dto.UserSignup) error {
	hashedPassword, err := generateHashedPassword(input.Password)
	if err != nil {
		return err
	}
	// TODO: token must be generated for the registeration
	if err = u.userRepository.CreateUser(ctx, &domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	}); err != nil {
		return err
	}
	return nil
}

func (u *userService) Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error) {
	user, err := u.findUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, customErr.ErrUserNotFound
	}

	if err = verifyPassword(input.Password, user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) ForgotPassword(ctx context.Context, email string) error {
	return nil
}

func (u *userService) SetPassword(ctx context.Context, token, password string) error {
	return nil
}

func (u *userService) CreateProfile(ctx context.Context, profile *dto.UserProfile) error {
	user, err := u.findUserById(ctx, profile.UserId)
	if err != nil {
		return err
	}

	if profile.FirstName != nil {
		user.FirstName = *profile.FirstName
	}

	if profile.LastName != nil {
		user.LastName = *profile.LastName
	}

	if profile.Email != nil {
		user.Email = *profile.Email
	}

	if profile.Phone != nil {
		user.Phone = *profile.Phone
	}

	if profile.Address.AddressLine1 != nil {
		user.Address.AddressLine1 = *profile.Address.AddressLine1
	}

	if profile.Address.AddressLine2 != nil {
		user.Address.AddressLine2 = *profile.Address.AddressLine2
	}

	if profile.Address.City != nil {
		user.Address.City = *profile.Address.City
	}

	if profile.Address.Country != nil {
		user.Address.Country = *profile.Address.Country
	}

	if profile.Address.PostCode != nil {
		user.Address.PostCode = *profile.Address.PostCode
	}

	if err = u.userRepository.SaveUser(ctx, user); err != nil {
		return errors.New("error saving user")
	}

	return nil
}

func (u *userService) GetProfile(ctx context.Context, id uint) (*domain.User, error) {
	user, err := u.findUserById(ctx, id)
	if err != nil {
		return nil, customErr.ErrUserNotFound
	}
	return user, nil
}

func (u *userService) Authenticate(ctx *fiber.Ctx) (*domain.User, error) {
	return nil, nil
}

func generateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slg.Logger.Error("error hashing password", "error", err.Error())
		return "", errors.New("error generating hashed password")
	}
	return string(hashedPassword), nil
}

func verifyPassword(plainPassword string, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		slg.Logger.Error("error verifying password", "error", err.Error())
		return errors.New("password does not match")
	}
	return nil
}

func NewUserService(userRepository repository.UserRepository, producer queue.Producer) UserService {
	return &userService{
		userRepository: userRepository,
		producer:       producer,
	}
}
