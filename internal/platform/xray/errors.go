package xray

import "errors"

var (
	ErrDisabled      = errors.New("xray client disabled")
	ErrRequestFailed = errors.New("xray request failed")
	ErrUnauthorized  = errors.New("xray unauthorized")
	ErrNotFound      = errors.New("xray object not found")
	ErrUnavailable   = errors.New("xray unavailable")
)
