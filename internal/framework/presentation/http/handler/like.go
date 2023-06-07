package handler

import (
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/dto"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type LikeHandler struct {
	likeRepo repository.LikeRepository
	*baseHandler
}

func newLikeHandler(
	repos *repository.Repositories,
	basebaseHandler *baseHandler,
) *LikeHandler {
	return &LikeHandler{
		likeRepo:    repos.LikeRepo,
		baseHandler: basebaseHandler,
	}
}

// CreateMyLike godoc
//
//	@Summary	Create new like by current user for specified joke.
//	@Tags		my likes
//	@Accept		json
//	@Produce	plain
//	@Param		authorization	header	string	true	"access token"
//	@Param		joke_id			path	uint	true	"ID of required joke"
//	@Success	201
//	@Failure	400					{object}	api.errorResponse
//	@Failure	500					{object}	api.errorResponse
//	@Router		/me/likes/{joke_id}	[post]
func (h *LikeHandler) CreateMyLike(c *fiber.Ctx) error {
	jokeID, err := c.ParamsInt("joke_id")
	if err != nil {
		return domain.NewErrWithMsg(domain.CodeInvalidArgument, "JokeID (uint) must be provided")
	}

	if err := h.likeRepo.CreateLike(
		c.Context(),
		&domain.Like{
			Username: c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
			JokeID:   jokeID,
		},
	); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// JokeLikes godoc
//
//	@Summary	Returns likes number for specified joke.
//	@Tags		likes
//	@Accept		json
//	@Produce	plain
//	@Param		authorization		header		string	true	"access token"
//	@Param		joke_id				path		uint	true	"ID of required joke"
//	@Success	200					{object}	dto.JokeLikesResponse
//	@Failure	400					{object}	api.errorResponse
//	@Failure	500					{object}	api.errorResponse
//	@Router		/likes/{joke_id}	[get]
func (h *LikeHandler) JokeLikes(c *fiber.Ctx) error {
	jokeID, err := c.ParamsInt("joke_id")
	if err != nil {
		return domain.NewErrWithMsg(domain.CodeInvalidArgument, "JokeID (uint) must be provided")
	}

	likesNumber, err := h.likeRepo.CalcLikesOfJoke(c.Context(), jokeID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.JokeLikesResponse{
		LikesNumber: likesNumber,
	})
}

// TODO: docs
func (h *LikeHandler) UsersWhoLikedJoke(c *fiber.Ctx) error {
	jokeID, err := c.ParamsInt("joke_id")
	if err != nil {
		return domain.NewErrWithMsg(domain.CodeInvalidArgument, "JokeID (uint) must be provided")
	}
	query := new(dto.SelectParamsQuery)
	c.QueryParser(query)

	users, err := h.likeRepo.GetUsersWhoLikedJoke(
		c.Context(),
		jokeID,
		dto.NewDomainSelectParams(query),
	)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewUsersResponse(users))
}

// TODO: docs
func (h *LikeHandler) MyLikedJokes(c *fiber.Ctx) error {
	query := new(dto.SelectParamsQuery)
	c.QueryParser(query)

	jokes, err := h.likeRepo.GetJokesUserLiked(
		c.Context(),
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
		dto.NewDomainSelectParams(query),
	)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewJokesResponse(jokes))
}

// TODO: docs
func (h *LikeHandler) JokesUserLiked(c *fiber.Ctx) error {
	username := c.Params("username")
	query := new(dto.SelectParamsQuery)
	c.QueryParser(query)

	jokes, err := h.likeRepo.GetJokesUserLiked(
		c.Context(),
		username,
		dto.NewDomainSelectParams(query),
	)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewJokesResponse(jokes))
}

// DeleteMyLike godoc
//
//	@Summary	Deletes like (made by current user) of a specified joke.
//	@Tags		my likes
//	@Accept		json
//	@Produce	plain
//	@Param		authorization	header	string	true	"access token"
//	@Param		joke_id			path	uint	true	"ID of required joke"
//	@Success	204
//	@Failure	400					{object}	api.errorResponse
//	@Failure	500					{object}	api.errorResponse
//	@Router		/me/likes/{joke_id}	[delete]
func (h *LikeHandler) DeleteMyLike(c *fiber.Ctx) error {
	jokeID, err := c.ParamsInt("joke_id")
	if err != nil {
		return domain.NewErrWithMsg(domain.CodeInvalidArgument, "JokeID (uint) must be provided")
	}

	if err := h.likeRepo.DeleteLike(
		c.Context(),
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
		jokeID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
