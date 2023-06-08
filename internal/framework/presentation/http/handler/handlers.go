package handler

import (
	"github.com/abc-valera/flugo-api/internal/application"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/domain/service"
)

type baseHandler struct {
	Log service.Logger
}

func newBaseHandler(logger service.Logger) *baseHandler {
	return &baseHandler{
		Log: logger,
	}
}

type Handlers struct {
	SignHandler    *SignHandler
	UserHandler    *UserHandler
	JokeHandler    *JokeHandler
	LikeHandler    *LikeHandler
	CommentHandler *CommentHandler
}

func NewHandlers(
	repos *repository.Repositories,
	services *service.Services,
	usecases *application.Usecases,
) *Handlers {
	baseHandler := newBaseHandler(services.Logger)
	return &Handlers{
		SignHandler:    newSignHandler(repos, usecases, baseHandler),
		UserHandler:    newUserHandler(repos, usecases, baseHandler),
		JokeHandler:    newJokeHandler(repos, usecases, baseHandler),
		LikeHandler:    newLikeHandler(repos, baseHandler),
		CommentHandler: newCommentHandler(repos, usecases, baseHandler),
	}
}
