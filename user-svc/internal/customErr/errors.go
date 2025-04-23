package customErr

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrDuplicateUser         = errors.New("user already exists")
	ErrInvalidUserResetToken = errors.New("invalid user reset token")
)
