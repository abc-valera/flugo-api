package repository

import (
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/jmoiron/sqlx"
)

func NewRepositories(db *sqlx.DB) *domain.Repositories {
	return &domain.Repositories{
		UserRepo:    newUserRepository(db),
		JokeRepo:    newJokeRepository(db),
		LikeRepo:    newLikeRepository(db),
		CommentRepo: newCommentRepository(db),
	}
}
