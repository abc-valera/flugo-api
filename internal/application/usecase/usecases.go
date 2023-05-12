package usecase

import (
	"github.com/abc-valera/flugo-api/internal/application/service"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
)

type Usecases struct {
	SignUsecase    SignUsecase
	UserUsecase    UserService
	JokeUsecase    JokeUsecase
	LikeUsecase    LikeUsecase
	CommentUsecase CommentUsecase
}

func NewUsecases(repos *repository.Repositories, services *service.Services) *Usecases {
	return &Usecases{
		SignUsecase:    newSignUsecase(repos.UserRepo, services.PasswordService, services.TokenService, services.EmailService),
		UserUsecase:    newUserService(repos.UserRepo, services.PasswordService),
		JokeUsecase:    newJokeUsecase(repos.UserRepo, repos.JokeRepo),
		LikeUsecase:    newLikeUsecase(repos.UserRepo, repos.LikeRepo),
		CommentUsecase: newCommentUsecase(repos.UserRepo, repos.CommentRepo),
	}
}
