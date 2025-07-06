package repository

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/customErr"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
	"gorm.io/gorm"
	"strings"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	SaveUser(ctx context.Context, user *domain.User) error
	FindUserByResetToken(ctx context.Context, token string) (*domain.User, error)
	FindUserById(ctx context.Context, id uint) (*domain.User, error)
}

type userRepository struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	err := u.dbWrite.WithContext(ctx).Create(user).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "UNIQUE constraint failed"):
			slg.Logger.Error("failed to create user: duplicate email", "error", err.Error())
			return customErr.ErrDuplicateUser
		default:
			return err
		}
	}
	return nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.dbRead.WithContext(ctx).Preload("Address").First(&user, "email = ?", email).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			slg.Logger.Warn("user not found", "email", email)
			return nil, customErr.ErrUserNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (u *userRepository) SaveUser(ctx context.Context, user *domain.User) error {
	if err := u.dbWrite.WithContext(ctx).Save(user).Error; err != nil {
		slg.Logger.Error("failed to save user", "error", err.Error())
		return customErr.ErrSaveUser
	}
	return nil
}

func (u *userRepository) FindUserByResetToken(ctx context.Context, token string) (*domain.User, error) {
	var user domain.User
	if err := u.dbRead.WithContext(ctx).Where("reset_token = ?", token).First(&user).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			slg.Logger.Warn("user not found", "token", token)
			return nil, customErr.ErrInvalidUserResetToken
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (u *userRepository) FindUserById(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := u.dbRead.WithContext(ctx).First(&user, id).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			slg.Logger.Warn("user not found", "id", id)
			return nil, customErr.ErrUserNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func NewUserRepository(dbWrite *gorm.DB, dbRead *gorm.DB) UserRepository {
	return &userRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
