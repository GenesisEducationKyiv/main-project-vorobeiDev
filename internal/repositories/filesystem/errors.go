package filesystem

import "errors"

var ErrUserExists = errors.New("user already exists")
var ErrInvalidEmail = errors.New("invalid email address")
