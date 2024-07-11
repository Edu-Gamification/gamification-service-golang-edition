package errors

import "errors"

var (
	UserNotFoundErr      = errors.New("user not found")
	ClanNotFoundErr      = errors.New("clan not found")
	EventTypeNotFoundErr = errors.New("event type not found")
	ErrNotFound          = errors.New("not found")
)
