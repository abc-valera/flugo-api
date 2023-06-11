package dto

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
)

// InsertComment represents comment data which should be added into the database
type InsertComment struct {
	Username string `db:"username"`
	JokeID   int    `db:"joke_id"`
	Text     string `db:"text"`
}

func NewInsertComment(comment *domain.Comment) *InsertComment {
	return &InsertComment{
		Username: comment.Username,
		JokeID:   comment.JokeID,
		Text:     comment.Text,
	}
}

// ReturnComment represents comment data which is returned from the database
type ReturnComment struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	JokeID    int       `db:"joke_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func NewDomainComment(comment *ReturnComment) *domain.Comment {
	return &domain.Comment{
		ID:        comment.ID,
		Username:  comment.Username,
		JokeID:    comment.JokeID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}

// ReturnComments represents slice of ReturnComment type returned from the database
type ReturnComments []*ReturnComment

func NewDomainComments(dbComments ReturnComments) domain.Comments {
	comments := make(domain.Comments, len(dbComments))
	for i, comment := range dbComments {
		comments[i] = NewDomainComment(comment)
	}
	return comments
}
