package dto

import "github.com/abc-valera/flugo-api/internal/domain"

type CommentReponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	JokeID   int    `json:"joke_id"`
	Text     string `json:"text"`
}

func NewCommentResponse(comment *domain.Comment) *CommentReponse {
	return &CommentReponse{
		ID:       comment.ID,
		Username: comment.Username,
		JokeID:   comment.JokeID,
		Text:     comment.Text,
	}
}

type CommentsResponse []*CommentReponse

func NewCommentsResponse(comments domain.Comments) CommentsResponse {
	commentsResponse := make(CommentsResponse, len(comments))
	for i, comment := range comments {
		commentsResponse[i] = NewCommentResponse(comment)
	}
	return commentsResponse
}

type NewMyCommentRequest struct {
	JokeID int    `json:"joke_id"`
	Text   string `json:"text"`
}

type DeleteMyCommentRequest struct {
	CommentID int `json:"comment_id"`
}
