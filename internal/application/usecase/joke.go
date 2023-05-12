package usecase

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
)

type JokeUsecase interface {
	// UpdateJokeExplanation checks if joke is updated by its owner and updates joke's explanation.
	// Returns error if joke is being updated by other (not owner) user.
	//
	// Returned codes:
	//  - NotFound (if user with provided username doesn't exist)
	//  - PermissionDenied (if joke is created by another user)
	//  - Internal
	UpdateJokeExplanation(c context.Context, jokeID int, username, explanation string) error

	// DeleteJoke checks if joke is deleted by its owner and deletes the joke.
	// Returns error if joke is being deleted by another (who hasn't created it) user.
	//
	// Returned codes:
	//  - NotFound (if user with provided username doesn't exist)
	//  - PermissionDenied (if joke is created by another user)
	//  - Internal
	DeleteJoke(c context.Context, jokeID int, username string) error
}

type jokeUsecase struct {
	userRepo repository.UserRepository
	jokeRepo repository.JokeRepository
}

func newJokeUsecase(userRepo repository.UserRepository, jokeRepo repository.JokeRepository) JokeUsecase {
	return &jokeUsecase{
		userRepo: userRepo,
		jokeRepo: jokeRepo,
	}
}

func (s *jokeUsecase) UpdateJokeExplanation(c context.Context, jokeID int, username, explanation string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != username {
		return domain.NewErrWithMsg(domain.CodePermissionDenied, "Operation can be performed only by creator user")
	}

	return s.jokeRepo.UpdateJokeExplanation(c, jokeID, explanation)
}

func (s *jokeUsecase) DeleteJoke(c context.Context, jokeID int, username string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != username {
		return domain.NewErrWithMsg(domain.CodePermissionDenied, "Operation can be performed only by creator user")
	}

	return s.jokeRepo.DeleteJoke(c, jokeID)
}
