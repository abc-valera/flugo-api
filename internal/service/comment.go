package service

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
)

type CommentService interface {
	// DeleteComment checks if comment is deleted by its owner and deletes the comment.
	// Returns error if comment is being deleted by another (who hasn't created it) user.
	//
	// Returned codes:
	//  - NotFound (if user with provided username doesn't exist)
	//  - PermissionDenied (if comment is created by another user)
	//  - Internal
	DeleteComment(c context.Context, commentID int, username string) error
}

type commentService struct {
	userRepo    repository.UserRepository
	commentRepo repository.CommentRepository
}

func newCommentService(repos *repository.Repositories) CommentService {
	return &commentService{
		userRepo:    repos.UserRepository,
		commentRepo: repos.CommentRepository,
	}
}

func (s *commentService) DeleteComment(c context.Context, commentID int, username string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != username {
		return domain.NewErrWithMsg(domain.CodePermissionDenied, "Operation can be performed only by creator user")
	}

	return s.commentRepo.DeleteComment(c, commentID)
}
