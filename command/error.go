package command

import "errors"

// ErrAlreadyExistChannelKey is an error that indicates the channel key already exists.
var (
	ErrAlreadyExistChannelKey = errors.New("channel key already exists")
	ErrChannelNotFound        = errors.New("channel not found")
)
