package apperrors

import "errors"

var (

	// Common

	ErrNotFound     = errors.New("resource not found")
	ErrForbidden    = errors.New("forbidden")
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadRequest   = errors.New("bad request")

	// User

	ErrUserNotFound = errors.New("user not found")

	// Project

	ErrProjectNotFound = errors.New("project not found")

	// Task

	ErrTaskNotFound = errors.New("task not found")

	ErrInvalidStatus   = errors.New("invalid task status")
	ErrInvalidPriority = errors.New("invalid task priority")

	ErrTitleRequired = errors.New("title is required")
)

// VPN

var (
	ErrProtocolRequired    = errors.New("protocol is required")
	ErrUnsupportedProtocol = errors.New("unsupported protocol")
	ErrVPNUserNotFound     = errors.New("vpn user not found")
	ErrVPNConfigNotFound   = errors.New("vpn config not found")
)
