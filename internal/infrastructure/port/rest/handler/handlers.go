package handler

import (
	"github.com/abc-valera/flugo-api/internal/application/usecase"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
)

type Handlers struct {
	UserHandler    *UserHandler
	JokeHandler    *JokeHandler
	LikeHandler    *LikeHandler
	CommentHandler *CommentHandler
}

func NewHandlers(repos *repository.Repositories, services *usecase.Usecases) *Handlers {
	return &Handlers{
		UserHandler:    newUserHandler(repos, services),
		JokeHandler:    newJokeHandler(repos, services),
		LikeHandler:    newLikeHandler(repos),
		CommentHandler: newCommentHandler(repos, services),
	}
}
