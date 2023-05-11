package application

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
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
	userRepo    domain.UserRepository
	commentRepo domain.CommentRepository
}

func newCommentService(userRepo domain.UserRepository, commentRepo domain.CommentRepository) CommentService {
	return &commentService{
		userRepo:    userRepo,
		commentRepo: commentRepo,
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
