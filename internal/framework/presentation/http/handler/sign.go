package handler

import (
	"github.com/abc-valera/flugo-api/internal/application"
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/dto"
	"github.com/gofiber/fiber/v2"
)

type SignHandler struct {
	userRepo    repository.UserRepository
	signUsecase application.SignUsecase
	*baseHandler
}

func newSignHandler(
	repos *repository.Repositories,
	usecases *application.Usecases,
	baseHandler *baseHandler,
) *SignHandler {
	return &SignHandler{
		baseHandler: baseHandler,
		userRepo:    repos.UserRepo,
		signUsecase: usecases.SignUsecase,
	}
}

// SignUp godoc
//
//	@Summary		Create new user account
//	@Description	Performs sign-up of the new user account.
//	@Tags			sign
//	@Accept			json
//	@Produce		plain
//	@Param			sign-up	body	dto.SignUpRequest	true	"sign-up request"
//	@Success		200
//	@Failure		400			{object}	api.errorResponse
//	@Failure		500			{object}	api.errorResponse
//	@Router			/sign_up	[post]
func (h *SignHandler) SignUp(c *fiber.Ctx) error {
	req := new(dto.SignUpRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.signUsecase.SignUp(
		c.Context(),
		&domain.User{
			Username: req.Username,
			Email:    req.Email,
		},
		req.Password,
	); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// SignIn godoc
//
//	@Summary		Sign-in to existing account
//	@Description	Performs sign-in with email to an existing user account.
//	@Description	Returns user with access token for accessing secured endpoints and refresh token for renewing access token.
//	@Tags			sign
//	@Accept			json
//	@Produce		json
//	@Param			sign-in		body		dto.SignInRequest	true	"sign-in request"
//	@Success		200			{object}	dto.SignInResponse
//	@Failure		400			{object}	api.errorResponse
//	@Failure		500			{object}	api.errorResponse
//	@Router			/sign_in	[get]
func (h *SignHandler) SignIn(c *fiber.Ctx) error {
	req := new(dto.SignInRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	user, access, refresh, err := h.signUsecase.SignIn(
		c.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&dto.SignInResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		User:         dto.NewUserResponse(user),
	})
}

// SignRefresh godoc
//
//	@Summary		Refresh the access token
//	@Description	Exchanges the refresh token for a new access token.
//	@Tags			sign
//	@Accept			json
//	@Produce		json
//	@Param			sign-refresh		body		dto.SignRefreshRequest	true	"sign-refresh request"
//	@Success		200			{object}	dto.SignRefreshResponse
//	@Failure		400			{object}	api.errorResponse
//	@Failure		401			{object}	api.errorResponse
//	@Failure		500			{object}	api.errorResponse
//	@Router			/sign_refresh	[get]
func (h *SignHandler) SignRefresh(c *fiber.Ctx) error {
	req := new(dto.SignRefreshRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	access, err := h.signUsecase.SignRefresh(c.Context(), req.RefreshToken)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&dto.SignRefreshResponse{
		AccessToken: access,
	})
}
