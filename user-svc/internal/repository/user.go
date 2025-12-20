package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"gorm.io/gorm"
	"strings"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	SaveUser(ctx context.Context, user *domain.User) error
	FindUserByResetToken(ctx context.Context, token string) (*domain.User, error)
	FindUserById(ctx context.Context, id uint) (*domain.User, error)
	WithTx(tx *gorm.DB) UserRepository
}

type userRepository struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
	tx      *gorm.DB
	logger  *zerolog.Logger
}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	err := exec(u.dbWrite, u.tx).WithContext(ctx).Create(user).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "UNIQUE constraint failed"):
			u.logger.Warn().Err(err).Msg("duplicate user detected")
			return ErrDuplicateUser
		default:
			return err
		}
	}
	return nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := exec(u.dbRead, u.tx).WithContext(ctx).Preload("Address").First(&user, "email = ?", email).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			u.logger.Debug().Str("email", email).Msg("user not found")
			return nil, ErrUserNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (u *userRepository) SaveUser(ctx context.Context, user *domain.User) error {
	if err := exec(u.dbWrite, u.tx).WithContext(ctx).Save(user).Error; err != nil {
		u.logger.Warn().Err(err).Uint("user_id", user.ID).Msg("failed to save user")
		return ErrSaveUser
	}
	return nil
}

func (u *userRepository) FindUserByResetToken(ctx context.Context, token string) (*domain.User, error) {
	var user domain.User
	if err := exec(u.dbRead, u.tx).WithContext(ctx).Where("reset_token = ?", token).First(&user).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			u.logger.Info().Msg("invalid reset token")
			return nil, ErrInvalidUserResetToken
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (u *userRepository) FindUserById(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := exec(u.dbRead, u.tx).WithContext(ctx).First(&user, id).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			u.logger.Debug().Uint("user_id", id).Msg("user not found")
			return nil, ErrUserNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (u *userRepository) WithTx(tx *gorm.DB) UserRepository {
	return &userRepository{
		dbWrite: u.dbWrite,
		dbRead:  u.dbRead,
		tx:      tx,
	}
}

func NewUserRepository(dbWrite *gorm.DB, dbRead *gorm.DB, logger *zerolog.Logger) UserRepository {
	return &userRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
		logger:  logger,
	}
}
