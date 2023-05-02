package handler

import (
	"github.com/abc-valera/flugo-api/internal/delivery/dto"
	"github.com/abc-valera/flugo-api/internal/delivery/middleware"
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
	"github.com/abc-valera/flugo-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	commentRepo    repository.CommentRepository
	commentService service.CommentService
}

func newCommentHandler(repos *repository.Repositories, services *service.Services) *CommentHandler {
	return &CommentHandler{
		commentRepo:    repos.CommentRepository,
		commentService: services.CommentService,
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
		domain.NewComment(
			c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
			req.Text,
			req.JokeID),
	); err != nil {
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
