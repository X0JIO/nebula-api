package middleware

import "context"

type contextKey string

const (
	ContextUserID    contextKey = "user_id"
	ContextRole      contextKey = "role"
	ContextProjectID contextKey = "project_id"
)

func WithProjectID(
	ctx context.Context,
	projectID string,
) context.Context {

	return context.WithValue(
		ctx,
		ContextProjectID,
		projectID,
	)
}

func ProjectID(
	ctx context.Context,
) string {

	projectID, ok := ctx.Value(ContextProjectID).(string)
	if !ok {
		return ""
	}

	return projectID
}

func UserID(
	ctx context.Context,
) string {

	userID, ok := ctx.Value(ContextUserID).(string)
	if !ok {
		return ""
	}

	return userID
}

func Role(
	ctx context.Context,
) string {

	role, ok := ctx.Value(ContextRole).(string)
	if !ok {
		return ""
	}

	return role
}
