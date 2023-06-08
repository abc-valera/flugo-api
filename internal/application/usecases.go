package application

import (
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/domain/service"
)

type Usecases struct {
	SignUsecase    SignUsecase
	UserUsecase    UserService
	JokeUsecase    JokeUsecase
	LikeUsecase    LikeUsecase
	CommentUsecase CommentUsecase
}

func NewUsecases(
	repos *repository.Repositories,
	services *service.Services,
	msgBroker service.MessagingBroker,
) *Usecases {
	return &Usecases{
		SignUsecase:    newSignUsecase(repos.UserRepo, services.PasswordMaker, services.TokenMaker, services.EmailSender, msgBroker),
		UserUsecase:    newUserService(repos.UserRepo, services.PasswordMaker),
		JokeUsecase:    newJokeUsecase(repos.UserRepo, repos.JokeRepo),
		LikeUsecase:    newLikeUsecase(repos.UserRepo, repos.LikeRepo),
		CommentUsecase: newCommentUsecase(repos.UserRepo, repos.CommentRepo),
	}
}
