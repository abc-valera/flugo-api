package handler

import (
	"github.com/abc-valera/flugo-api/internal/application"
	"github.com/abc-valera/flugo-api/internal/domain"
)

type Handlers struct {
	UserHandler    *UserHandler
	JokeHandler    *JokeHandler
	LikeHandler    *LikeHandler
	CommentHandler *CommentHandler
}

func NewHandlers(repos *domain.Repositories, services *application.Services) *Handlers {
	return &Handlers{
		UserHandler:    newUserHandler(repos, services),
		JokeHandler:    newJokeHandler(repos, services),
		LikeHandler:    newLikeHandler(repos),
		CommentHandler: newCommentHandler(repos, services),
	}
}
