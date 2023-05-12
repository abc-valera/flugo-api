package usecase

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
)

// TODO: some strange logic, can be remade

type LikeUsecase interface {
	// DeleteJoke checks if like is deleted by its owner and deletes the like.
	// Returns error if like is being deleted by another (who hasn't created it) user.
	//
	// Returned codes:
	//  - NotFound (if user with provided username doesn't exist)
	//  - PermissionDenied (if like is created by another user)
	//  - Internal
	DeleteLike(c context.Context, jokeID int, username string) error
}

type likeUsecase struct {
	userRepo repository.UserRepository
	likeRepo repository.LikeRepository
}

func newLikeUsecase(userRepo repository.UserRepository, likeRepo repository.LikeRepository) LikeUsecase {
	return &likeUsecase{
		userRepo: userRepo,
		likeRepo: likeRepo,
	}
}

func (s *likeUsecase) DeleteLike(c context.Context, jokeID int, username string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != username {
		return domain.NewErrWithMsg(domain.CodePermissionDenied, "Operation can be performed only by creator user")
	}

	return s.DeleteLike(c, jokeID, username)
}
