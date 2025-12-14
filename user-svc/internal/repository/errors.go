package repository

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrDuplicateUser         = errors.New("user already exists")
	ErrInvalidUserResetToken = errors.New("invalid user reset token")
	ErrSaveUser              = errors.New("failed to save user")
)
