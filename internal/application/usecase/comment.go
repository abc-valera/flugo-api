package usecase

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
)

type CommentUsecase interface {
	// DeleteComment checks if comment is deleted by its owner and deletes the comment.
	// Returns error if comment is being deleted by another (who hasn't created it) user.
	//
	// Returned codes:
	//  - NotFound (if user with provided username doesn't exist)
	//  - PermissionDenied (if comment is created by another user)
	//  - Internal
	DeleteComment(c context.Context, commentID int, username string) error
}

type commentUsecase struct {
	userRepo    repository.UserRepository
	commentRepo repository.CommentRepository
}

func newCommentUsecase(userRepo repository.UserRepository, commentRepo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{
		userRepo:    userRepo,
		commentRepo: commentRepo,
	}
}

func (s *commentUsecase) DeleteComment(c context.Context, commentID int, username string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != username {
		return domain.NewErrWithMsg(domain.CodePermissionDenied, "Operation can be performed only by creator user")
	}

	return s.commentRepo.DeleteComment(c, commentID)
}
