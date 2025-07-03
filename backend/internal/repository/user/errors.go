package user

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidObjectID    = errors.New("invalid object id")
	ErrDatabaseConnection = errors.New("database connection error")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrOAuthUserNotFound  = errors.New("oauth user not found")
	ErrUpdateFailed       = errors.New("update operation failed")
)
