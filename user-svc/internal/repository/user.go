package repository

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"gorm.io/gorm"
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
	return u.dbWrite.WithContext(ctx).Create(user).Error
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.dbRead.WithContext(ctx).Preload("Address").First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) SaveUser(ctx context.Context, user *domain.User) error {
	return u.dbWrite.WithContext(ctx).Save(user).Error
}

func (u *userRepository) FindUserByResetToken(ctx context.Context, token string) (*domain.User, error) {
	var user domain.User
	if err := u.dbRead.WithContext(ctx).Where("reset_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindUserById(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := u.dbRead.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(dbWrite *gorm.DB, dbRead *gorm.DB) UserRepository {
	return &userRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
