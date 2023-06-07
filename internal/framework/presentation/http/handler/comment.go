package handler

import (
	"github.com/abc-valera/flugo-api/internal/application/usecase"
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/dto"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	commentRepo    repository.CommentRepository
	commentService usecase.CommentUsecase
	*baseHandler
}

func newCommentHandler(
	repos *repository.Repositories,
	usecases *usecase.Usecases,
	baseHandler *baseHandler,
) *CommentHandler {
	return &CommentHandler{
		commentRepo:    repos.CommentRepo,
		commentService: usecases.CommentUsecase,
		baseHandler:    baseHandler,
	}
}

// CreateMyJoke godoc
//
//	@Summary	Creates new comment for a specified joke by current user.
//	@Tags		my comments
//	@Accept		json
//	@Produce	plain
//	@Param		authorization	header	string					true	"access token"
//	@Param		comment_data	body	dto.NewMyCommentRequest	true	"request for creating comment"
//	@Success	201
//	@Failure	400				{object}	api.errorResponse
//	@Failure	500				{object}	api.errorResponse
//	@Router		/me/comments	[post]
func (h *CommentHandler) NewMyComment(c *fiber.Ctx) error {
	req := new(dto.NewMyCommentRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.commentRepo.CreateComment(
		c.Context(),
		&domain.Comment{
			Username: c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
			Text:     req.Text,
			JokeID:   req.JokeID,
		}); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// GetJokeComments godoc
//
//	@Summary	Returns comments for specified joke.
//	@Tags		comments
//	@Accept		json
//	@Produce	plain
//	@Param		authorization		header		string					true	"access token"
//	@Param		joke_id				path		uint					true	"ID of required joke"
//	@Param		sorting				query		dto.SelectParamsQuery	true	"params for sorting"
//	@Success	200					{object}	dto.CommentsResponse
//	@Failure	400					{object}	api.errorResponse
//	@Failure	500					{object}	api.errorResponse
//	@Router		/comments/{joke_id}	[get]
func (h *CommentHandler) GetJokeComments(c *fiber.Ctx) error {
	jokeID, err := c.ParamsInt("joke_id")
	if err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "JokeID must be an integer"}
	}
	query := new(dto.SelectParamsQuery)
	c.QueryParser(query)

	comments, err := h.commentRepo.GetCommentsOfJoke(
		c.Context(),
		jokeID,
		dto.NewDomainSelectParams(query))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewCommentsResponse(comments))
}

// DeleteMyComment godoc
//
//	@Summary	Deletes specified comment by current user.
//	@Tags		my comments
//	@Accept		json
//	@Produce	plain
//	@Param		authorization	header	string	true	"access token"
//	@Param		comment_id		path	uint	true	"ID of required comment"
//	@Success	204
//	@Failure	400							{object}	api.errorResponse
//	@Failure	500							{object}	api.errorResponse
//	@Router		/me/comments/{comment_id}	[delete]
func (h *CommentHandler) DeleteMyComment(c *fiber.Ctx) error {
	commentID, err := c.ParamsInt("comment_id")
	if err != nil {
		return domain.NewErrWithMsg(domain.CodeInvalidArgument, "JokeID (uint) must be provided")
	}

	if err := h.commentService.DeleteComment(
		c.Context(),
		commentID,
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
	); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
