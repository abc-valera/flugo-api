package application

import "github.com/abc-valera/flugo-api/internal/domain"

type Services struct {
	SignService    SignService
	UserService    UserService
	JokeService    JokeService
	LikeService    LikeService
	CommentService CommentService
}

func NewServices(repos *domain.Repositories, pkgs *domain.Packages) *Services {
	return &Services{
		SignService:    newSignService(repos.UserRepo, pkgs.PasswordPkg, pkgs.TokenPkg, pkgs.EmailPkg),
		UserService:    newUserService(repos.UserRepo, pkgs.PasswordPkg),
		JokeService:    newJokeService(repos.UserRepo, repos.JokeRepo),
		LikeService:    newLikeService(repos.UserRepo, repos.LikeRepo),
		CommentService: newCommentService(repos.UserRepo, repos.CommentRepo),
	}
}
