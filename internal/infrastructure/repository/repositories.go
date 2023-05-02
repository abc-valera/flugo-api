package repository

import "github.com/jmoiron/sqlx"

type Repositories struct {
	UserRepository    UserRepository
	JokeRepository    JokeRepository
	LikeRepository    LikeRepository
	CommentRepository CommentRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository:    newUserRepository(db),
		JokeRepository:    newJokeRepository(db),
		LikeRepository:    newLikeRepository(db),
		CommentRepository: newCommentRepository(db),
	}
}
