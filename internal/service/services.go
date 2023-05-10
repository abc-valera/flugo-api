package service

import (
	"github.com/abc-valera/flugo-api/internal/infrastructure/framework"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
)

type Services struct {
	UserService    UserService
	JokeService    JokeService
	CommentService CommentService
}

func NewServices(repos *repository.Repositories, frs *framework.Frameworks) *Services {
	return &Services{
		UserService:    newUserService(repos, frs),
		JokeService:    newJokeService(repos),
		CommentService: newCommentService(repos),
	}
}
