package application

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
)

type JokeService interface {
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

type jokeService struct {
	userRepo domain.UserRepository
	jokeRepo domain.JokeRepository
}

func newJokeService(userRepo domain.UserRepository, jokeRepo domain.JokeRepository) JokeService {
	return &jokeService{
		userRepo: userRepo,
		jokeRepo: jokeRepo,
	}
}

func (s *jokeService) UpdateJokeExplanation(c context.Context, jokeID int, username, explanation string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != username {
		return domain.NewErrWithMsg(domain.CodePermissionDenied, "Operation can be performed only by creator user")
	}

	return s.jokeRepo.UpdateJokeExplanation(c, jokeID, explanation)
}

func (s *jokeService) DeleteJoke(c context.Context, jokeID int, username string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != username {
		return domain.NewErrWithMsg(domain.CodePermissionDenied, "Operation can be performed only by creator user")
	}

	return s.jokeRepo.DeleteJoke(c, jokeID)
}
