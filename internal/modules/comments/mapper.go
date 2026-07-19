package comments

import db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

func ToResponse(comment db.Comment) CommentResponse {

	return CommentResponse{
		ID:        comment.ID.String(),
		TaskID:    comment.TaskID.String(),
		AuthorID:  comment.AuthorID.String(),
		Body:      comment.Body,
		CreatedAt: comment.CreatedAt.Time,
		UpdatedAt: comment.UpdatedAt.Time,
	}
}

func ToResponses(comments []db.Comment) []CommentResponse {

	result := make([]CommentResponse, 0, len(comments))

	for _, comment := range comments {
		result = append(result, ToResponse(comment))
	}

	return result
}
