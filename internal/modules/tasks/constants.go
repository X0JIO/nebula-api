package tasks

import "github.com/X0JIO/nebula-api/internal/shared/apperrors"

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusReview     = "review"
	StatusDone       = "done"
	StatusArchived   = "archived"
)

const (
	PriorityLow      = "low"
	PriorityMedium   = "medium"
	PriorityHigh     = "high"
	PriorityCritical = "critical"
)

func ValidateStatus(status string) error {

	switch status {

	case
		StatusTodo,
		StatusInProgress,
		StatusReview,
		StatusDone,
		StatusArchived:

		return nil

	default:

		return apperrors.ErrInvalidStatus

	}
}

func ValidatePriority(priority string) error {

	switch priority {

	case
		PriorityLow,
		PriorityMedium,
		PriorityHigh,
		PriorityCritical:

		return nil

	default:

		return apperrors.ErrInvalidPriority

	}
}
