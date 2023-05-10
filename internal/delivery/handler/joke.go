package handler

import (
	"github.com/abc-valera/flugo-api/internal/delivery/dto"
	"github.com/abc-valera/flugo-api/internal/delivery/middleware"
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
	"github.com/abc-valera/flugo-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type JokeHandler struct {
	jokeRepo    repository.JokeRepository
	jokeService service.JokeService
}

func newJokeHandler(repos *repository.Repositories, services *service.Services) *JokeHandler {
	return &JokeHandler{
		jokeRepo:    repos.JokeRepository,
		jokeService: services.JokeService,
	}
}

// CreateMyJoke godoc
//
//	@Summary	Create new joke for current user.
//	@Tags		my jokes
//	@Accept		json
//	@Produce	plain
//	@Param		authorization	header	string					true	"access token"
//	@Param		joke_data		body	dto.CreateMyJokeRequest	true	"request for creating joke"
//	@Success	201
//	@Failure	400			{object}	api.errorResponse
//	@Failure	500			{object}	api.errorResponse
//	@Router		/me/jokes	[post]
func (h *JokeHandler) CreateMyJoke(c *fiber.Ctx) error {
	req := new(dto.CreateMyJokeRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.jokeRepo.CreateJoke(
		c.Context(),
		domain.NewJoke(
			c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
			req.Title,
			req.Text,
			req.Explanation),
	); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// GetJoke godoc
//
//	@Summary	Returns joke with specified ID from the database.
//	@Tags		jokes
//	@Accept		json
//	@Produce	json
//	@Param		authorization		header		string	true	"access token"
//	@Param		joke_id				path		uint	true	"id of required joke"
//	@Success	200					{object}	dto.JokeResponse
//	@Failure	400					{object}	api.errorResponse
//	@Failure	500					{object}	api.errorResponse
//	@Router		/jokes/{joke_id}	[get]
func (h *JokeHandler) GetJoke(c *fiber.Ctx) error {
	jokeID, err := c.ParamsInt("joke_id")
	if err != nil {
		return domain.NewErrWithMsg(domain.CodeInvalidArgument, "JokeID (uint) must be provided")
	}

	joke, err := h.jokeRepo.GetJokeByID(c.Context(), jokeID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewJokeResponse(joke))
}

// GetMyJokes godoc
//
//	@Summary	Returns all jokes (limited by limit and offset) created by the current user.
//	@Tags		my jokes
//	@Accept		json
//	@Produce	json
//	@Param		authorization	header		string					true	"access token"
//	@Param		sorting			query		dto.SelectParamsQuery	true	"params for sorting"
//	@Success	200				{object}	dto.JokeResponse
//	@Failure	400				{object}	api.errorResponse
//	@Failure	500				{object}	api.errorResponse
//	@Router		/me/jokes [get]
func (h *JokeHandler) GetMyJokes(c *fiber.Ctx) error {
	query := new(dto.SelectParamsQuery)
	c.QueryParser(query)

	jokes, err := h.jokeRepo.GetJokesByAuthor(
		c.Context(),
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
		dto.NewDomainSelectParams(query))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.NewJokesResponse(jokes))
}

// GetUserJokes godoc
//
//	@Summary	Returns all jokes (limited by limit and offset) by specified user.
//	@Tags		jokes
//	@Accept		json
//	@Produce	json
//	@Param		authorization			header		string					true	"access token"
//	@Param		username				path		string					true	"username of required user"
//	@Param		sorting					query		dto.SelectParamsQuery	true	"params for sorting"
//	@Success	200						{object}	dto.JokesResponse
//	@Failure	400						{object}	api.errorResponse
//	@Failure	500						{object}	api.errorResponse
//	@Router		/jokes/by/{username}	[get]
func (h *JokeHandler) GetUserJokes(c *fiber.Ctx) error {
	username := c.Params("username")
	query := new(dto.SelectParamsQuery)
	c.QueryParser(query)

	jokes, err := h.jokeRepo.GetJokesByAuthor(
		c.Context(),
		username,
		dto.NewDomainSelectParams(query))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.NewJokesResponse(jokes))
}

// GetAllJokes godoc
//
//	@Summary	Returns all jokes (limited by limit and offset) from the database.
//	@Tags		jokes
//	@Accept		json
//	@Produce	json
//	@Param		authorization	header		string					true	"access token"
//	@Param		sorting			query		dto.SelectParamsQuery	true	"params for sorting"
//	@Success	200				{object}	dto.JokesResponse
//	@Failure	400				{object}	api.errorResponse
//	@Failure	500				{object}	api.errorResponse
//	@Router		/jokes [get]
func (h *JokeHandler) GetAllJokes(c *fiber.Ctx) error {
	query := new(dto.SelectParamsQuery)
	c.QueryParser(query)

	jokes, err := h.jokeRepo.GetJokes(
		c.Context(),
		dto.NewDomainSelectParams(query))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.NewJokesResponse(jokes))
}

// TODO: Implement
func (h *JokeHandler) SearchJokesByTitle(c *fiber.Ctx) error {
	return nil
}

// UpdateMyJokeExplanation godoc
//
//	@Summary	Updates explanation of specified joke created by current user.
//	@Tags		my jokes
//	@Accept		json
//	@Produce	plain
//	@Param		authorization	header	string								true	"access token"
//	@Param		explanation		body	dto.UpdateMyJokeExplanationRequest	true	"request for updating joke explanation"
//	@Success	200
//	@Failure	400						{object}	api.errorResponse
//	@Failure	500						{object}	api.errorResponse
//	@Router		/me/jokes/explanation	[put]
func (h *JokeHandler) UpdateMyJokeExplanation(c *fiber.Ctx) error {
	req := new(dto.UpdateMyJokeExplanationRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.jokeRepo.UpdateJokeExplanation(c.Context(), req.JokeID, req.Explanation); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DeleteMyJoke godoc
//
//	@Summary	Deletes specified joke created by current user.
//	@Tags		my jokes
//	@Accept		json
//	@Produce	plain
//	@Param		authorization	header	string	true	"access token"
//	@Param		joke_id			path	uint	true	"ID of required joke"
//	@Success	204
//	@Failure	400					{object}	api.errorResponse
//	@Failure	500					{object}	api.errorResponse
//	@Router		/me/jokes/{joke_id}	[delete]
func (h *JokeHandler) DeleteMyJoke(c *fiber.Ctx) error {
	jokeID, err := c.ParamsInt("joke_id")
	if err != nil {
		return domain.NewErrWithMsg(domain.CodeInvalidArgument, "JokeID (uint) must be provided")
	}

	if err := h.jokeRepo.DeleteJoke(c.Context(), jokeID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
