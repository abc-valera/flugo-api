package handler

import (
	"github.com/abc-valera/flugo-api/internal/application"
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/dto"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepo    repository.UserRepository
	signUsecase application.SignUsecase
	userUsecase application.UserService
	*baseHandler
}

func newUserHandler(
	repos *repository.Repositories,
	usecases *application.Usecases,
	baseHandler *baseHandler,
) *UserHandler {
	return &UserHandler{
		userRepo:    repos.UserRepo,
		signUsecase: usecases.SignUsecase,
		userUsecase: usecases.UserUsecase,
		baseHandler: baseHandler,
	}
}

// GetMe godoc
//
//	@Summary					Returns current user.
//	@Description				Returns data about user whose auth credentials provided.
//	@Tags						me
//	@Accept						json
//	@Produce					json
//	@Param						authorization	header		string	true	"access token"
//	@Success					200				{object}	dto.UserResponse
//	@Failure					401				{object}	api.errorResponse
//	@Failure					500				{object}	api.errorResponse
//	@securityDefinitions.basic	BasicAuth
//	@Router						/me [get]
func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	user, err := h.userRepo.GetUserByUsername(c.Context(), c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.NewUserResponse(user))
}

// SearchUsersByUsername godoc
//
//	@Summary					Searches for users in specified range.
//	@Description				Searches for users in specified range whose usernames follow the pattern *<username>*.
//	@Tags						users
//	@Accept						json
//	@Produce					json
//	@Param						authorization	header		string					true	"access token"
//	@Param						username		path		string					true	"pattern for a username"
//	@Param						sorting			query		dto.SelectParamsQuery	true	"params for sorting"
//	@Success					200				{object}	dto.UsersResponse
//	@Failure					401				{object}	api.errorResponse
//	@Failure					500				{object}	api.errorResponse
//	@SecurityDefinitions.basic	BasicAuth
//	@Router						/users/search/{username} [get]
func (h *UserHandler) SearchUsersByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	query := new(dto.SelectParamsQuery)
	c.QueryParser(query)

	users, err := h.userRepo.SearchUsersByUsername(c.Context(), username, dto.NewDomainSelectParams(query))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewUsersResponse(users))
}

// UpdateMyPassword godoc
//
//	@Summary					Updates current user's password.
//	@Description				Compares provided password with original password hash and updates user with new password.
//	@Tags						me
//	@Accept						json
//	@Produce					json
//	@Param						authorization	header		string						true	"access token"
//	@Param						passwords		body		dto.UpdateMyPasswordRequest	true	"old and new passwords"
//	@Success					200				{array}		dto.UserResponse
//	@Failure					400				{object}	api.errorResponse
//	@Failure					401				{object}	api.errorResponse
//	@Failure					500				{object}	api.errorResponse
//	@SecurityDefinitions.basic	BasicAuth
//	@Router						/me/password [put]
func (h *UserHandler) UpdateMyPassword(c *fiber.Ctx) error {
	req := new(dto.UpdateMyPasswordRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.userUsecase.UpdateUserPassword(
		c.Context(),
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
		req.OldPassword,
		req.NewPassword); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// UpdateMyFullname godoc
//
//	@Summary					Updates current user's fullname.
//	@Tags						me
//	@Accept						json
//	@Produce					json
//	@Param						authorization	header		string						true	"access token"
//	@Param						fullname		body		dto.UpdateMyFullnameRequest	true	"new fullname"
//	@Success					200				{array}		dto.UserResponse
//	@Failure					401				{object}	api.errorResponse
//	@Failure					500				{object}	api.errorResponse
//	@SecurityDefinitions.basic	BasicAuth
//	@Router						/me/fullname [put]
func (h *UserHandler) UpdateMyFullname(c *fiber.Ctx) error {
	req := new(dto.UpdateMyFullnameRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.userRepo.UpdateUserFullname(
		c.Context(),
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
		req.Fullname); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// UpdateMyStatus godoc
//
//	@Summary					Updates current user's status.
//	@Tags						me
//	@Accept						json
//	@Produce					json
//	@Param						authorization	header		string						true	"access token"
//	@Param						status			body		dto.UpdateMyStatusRequest	true	"new status"
//	@Success					200				{array}		dto.UserResponse
//	@Failure					401				{object}	api.errorResponse
//	@Failure					500				{object}	api.errorResponse
//	@SecurityDefinitions.basic	BasicAuth
//	@Router						/me/status [put]
func (h *UserHandler) UpdateMyStatus(c *fiber.Ctx) error {
	req := new(dto.UpdateMyStatusRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.userRepo.UpdateUserStatus(
		c.Context(),
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
		req.Status); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// UpdateMyBio godoc
//
//	@Summary					Updates current user's bio.
//	@Tags						me
//	@Accept						json
//	@Produce					json
//	@Param						authorization	header		string					true	"access token"
//	@Param						status			body		dto.UpdateMyBioRequest	true	"new bio"
//	@Success					200				{array}		dto.UserResponse
//	@Failure					401				{object}	api.errorResponse
//	@Failure					500				{object}	api.errorResponse
//	@SecurityDefinitions.basic	BasicAuth
//	@Router						/me/bio [put]
func (h *UserHandler) UpdateMyBio(c *fiber.Ctx) error {
	req := new(dto.UpdateMyBioRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.userRepo.UpdateUserBio(
		c.Context(),
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
		req.Bio); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DeleteMe godoc
//
//	@Summary					Deletes current user.
//	@Description				Compares provided password with original password hash and deletes user forever!
//	@Tags						me
//	@Accept						json
//	@Produce					json
//	@Param						authorization	header		string				true	"access token"
//	@Param						password		body		dto.DeleteMeRequest	true	"password"
//	@Success					200				{array}		dto.UserResponse
//	@Failure					401				{object}	api.errorResponse
//	@Failure					500				{object}	api.errorResponse
//	@SecurityDefinitions.basic	BasicAuth
//	@Router						/me [delete]
func (h *UserHandler) DeleteMe(c *fiber.Ctx) error {
	req := new(dto.DeleteMeRequest)
	if err := c.BodyParser(req); err != nil {
		return &domain.Error{Code: domain.CodeInvalidArgument, Msg: "Body data should be in json format"}
	}

	if err := h.userUsecase.DeleteUser(
		c.Context(),
		c.Locals(middleware.AuthPayloadKey).(*domain.Payload).Username,
		req.Password); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
