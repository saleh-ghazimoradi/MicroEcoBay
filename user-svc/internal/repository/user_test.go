package repository

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/customErr"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB(t *testing.T) (*gorm.DB, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Address{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db, db
}

func TestUserRepository_CreateUser(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewUserRepository(dbWrite, dbRead)
	ctx := context.Background()

	tests := []struct {
		name    string
		user    *domain.User
		wantErr error
	}{
		{name: "successfully create user",
			user: &domain.User{
				Email:     "test@example.com",
				Password:  "password",
				FirstName: "John",
				LastName:  "Doe",
			},
			wantErr: nil,
		},
		{
			name: "duplicate email",
			user: &domain.User{
				Email:     "test@example.com",
				Password:  "password",
				FirstName: "John",
				LastName:  "Doe",
			},
			wantErr: customErr.ErrDuplicateUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateUser(ctx, tt.user)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)

				var createdUser domain.User
				err = dbRead.WithContext(ctx).Where("email=?", tt.user.Email).First(&createdUser).Error
				assert.NoError(t, err)
				assert.Equal(t, tt.user.Email, createdUser.Email)
			}
		})
	}
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewUserRepository(dbWrite, dbRead)
	ctx := context.Background()

	user := &domain.User{
		Email:     "find@example.com",
		Password:  "password",
		FirstName: "Jane",
		LastName:  "Doe",
		Address: domain.Address{
			AddressLine1: "123 Main St",
			City:         "Tehran",
			PostCode:     "12345",
			Country:      "Iran",
		},
	}
	err := dbWrite.WithContext(ctx).Create(user).Error
	assert.NoError(t, err)

	tests := []struct {
		name    string
		email   string
		wantErr error
	}{
		{
			name:    "user found",
			email:   "find@example.com",
			wantErr: nil,
		},
		{
			name:    "user not found",
			email:   "notfound@example.com",
			wantErr: customErr.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foundUser, err := repo.FindUserByEmail(ctx, tt.email)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, foundUser)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, foundUser)
				assert.Equal(t, tt.email, foundUser.Email)
				assert.Equal(t, user.Address.AddressLine1, foundUser.Address.AddressLine1)
			}
		})
	}
}

func TestUserRepository_SaveUser(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewUserRepository(dbWrite, dbRead)
	ctx := context.Background()

	user := &domain.User{
		Email:     "Save@example.com",
		Password:  "password",
		FirstName: "John",
		LastName:  "Doe",
	}

	err := dbWrite.WithContext(ctx).Create(user).Error
	assert.NoError(t, err)

	user.FirstName = "UpdatedJohn"
	err = repo.SaveUser(ctx, user)
	assert.NoError(t, err)

	var updatedUser domain.User
	err = dbRead.WithContext(ctx).Where("email=?", user.Email).First(&updatedUser).Error
	assert.NoError(t, err)
	assert.Equal(t, user.Email, updatedUser.Email)
}

func TestUserRepository_FindUserByResetToken(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewUserRepository(dbWrite, dbRead)
	ctx := context.Background()

	user := &domain.User{
		Email:      "token@example.com",
		Password:   "password",
		ResetToken: "rest123",
	}
	err := dbWrite.WithContext(ctx).Create(user).Error
	assert.NoError(t, err)

	tests := []struct {
		name    string
		token   string
		wantErr error
	}{
		{
			name:    "user found by token",
			token:   "rest123",
			wantErr: nil,
		},
		{
			name:    "invalid token",
			token:   "invalid",
			wantErr: customErr.ErrInvalidUserResetToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foundUser, err := repo.FindUserByResetToken(ctx, tt.token)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, foundUser)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, foundUser)
				assert.Equal(t, user.Email, foundUser.Email)
			}
		})
	}
}

func TestUserRepository_FindUserById(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewUserRepository(dbWrite, dbRead)
	ctx := context.Background()

	user := &domain.User{
		Email:     "id@example.com",
		Password:  "password",
		FirstName: "John",
		LastName:  "Doe",
	}

	err := dbWrite.WithContext(ctx).Create(user).Error
	assert.NoError(t, err)

	tests := []struct {
		name    string
		id      uint
		wantErr error
	}{
		{
			name:    "user found by id",
			id:      user.ID,
			wantErr: nil,
		},
		{
			name:    "user not found",
			id:      999,
			wantErr: customErr.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foundUser, err := repo.FindUserById(ctx, tt.id)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, foundUser)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, foundUser)
				assert.Equal(t, user.Email, foundUser.Email)
			}
		})
	}
}
