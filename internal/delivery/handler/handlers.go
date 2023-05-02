package handler

import (
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
	"github.com/abc-valera/flugo-api/internal/service"
)

type Handlers struct {
	UserHandler    *UserHandler
	JokeHandler    *JokeHandler
	LikeHandler    *LikeHandler
	CommentHandler *CommentHandler
}

func NewHandlers(repos *repository.Repositories, services *service.Services) *Handlers {
	return &Handlers{
		UserHandler:    newUserHandler(repos, services),
		JokeHandler:    newJokeHandler(repos, services),
		LikeHandler:    newLikeHandler(repos),
		CommentHandler: newCommentHandler(repos, services),
	}
}
