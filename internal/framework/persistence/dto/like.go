package dto

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
)

// InsertLike represents like data which should be added into the database
type InsertLike struct {
	Username string `db:"username"`
	JokeID   int    `db:"joke_id"`
}

func NewInsertLike(like *domain.Like) *InsertLike {
	return &InsertLike{
		Username: like.Username,
		JokeID:   like.JokeID,
	}
}

// ReturnLike represents like data which is returned from the database
type ReturnLike struct {
	Username  string    `db:"username"`
	JokeID    int       `db:"joke_id"`
	CreatedAt time.Time `db:"created_at"`
}

func NewDomainLike(like *ReturnLike) *domain.Like {
	return &domain.Like{
		Username:  like.Username,
		JokeID:    like.JokeID,
		CreatedAt: like.CreatedAt,
	}
}

// ReturnLikes represents slice of ReturnLike type returned from the database
type ReturnLikes []*ReturnLike

func NewDomainLikes(dbLikes ReturnLikes) domain.Likes {
	likes := make(domain.Likes, len(dbLikes))
	for i, like := range dbLikes {
		likes[i] = NewDomainLike(like)
	}
	return likes
}
