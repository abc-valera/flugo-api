package repository

import (
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

func NewRepositories(db *sqlx.DB) *repository.Repositories {
	return &repository.Repositories{
		UserRepo:    newUserRepository(db),
		JokeRepo:    newJokeRepository(db),
		LikeRepo:    newLikeRepository(db),
		CommentRepo: newCommentRepository(db),
	}
}
