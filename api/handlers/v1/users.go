package v1

import (
	"api_gateway/api/handlers/models"
	pb "api_gateway/genproto/users"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetUserProfile  	godoc
// @Security ApiKeyAuth
// @Router 			/users/profile [get]
// @Summary 		Get User Profile
// @Description 	getting user profile by user id
// @Tags users
// @Accept 			json
// @Produce 		json
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) GetUserProfile(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()
	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		return handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
	}

	resp, err := h.services.UsersService().GetUserProfile(reqCtx, &pb.PrimaryKey{Id: user.Id})
	if err != nil {
		return handleResponse(ctx, h.log, "error while using GetUserProfile method of users service", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "", http.StatusOK, resp)
}

// UpdateUserProfile  	godoc
// @Security ApiKeyAuth
// @Router 			/users/update [put]
// @Summary 		Update User Profile
// @Description 	updating user profile
// @Tags users
// @Accept 			json
// @Produce 		json
// @Param 			user body models.UpdateUser true "user"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) UpdateUserProfile(ctx *fiber.Ctx) error {

	reqCtx := ctx.Context()
	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		return handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
	}

	req := pb.UpdateUser{}
	if err = ctx.BodyParser(&req); err != nil {
		return handleResponse(ctx, h.log, "error while getting request body", http.StatusBadRequest, err.Error())
	}
	req.Id = user.Id

	resp, err := h.services.UsersService().UpdateUserProfile(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "error while using UpdateUserProfile method of users service", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "", http.StatusOK, resp)
}

// ChangePassword  	godoc
// @Security ApiKeyAuth
// @Router 			/users/password [put]
// @Summary 		Change User Password
// @Description 	changing user password
// @Tags users
// @Accept 			json
// @Produce 		json
// @Param 			change_password body models.ChangePassword true "change_password"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) ChangePassword(ctx *fiber.Ctx) error {

	reqCtx := ctx.Context()
	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		return handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
	}

	req := models.ChangePassword{}
	if err = ctx.BodyParser(&req); err != nil {
		return handleResponse(ctx, h.log, "error while getting request body", http.StatusBadRequest, err.Error())
	}

	resp, err := h.services.UsersService().ChangePassword(reqCtx, &pb.ChangePassword{
		CurrentPassword: req.CurrentPassword,
		NewPassword:     req.NewPassword,
		UserId:          user.Id,
	})
	if err != nil {
		return handleResponse(ctx, h.log, "error while using ChangePassword method of users service", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "", http.StatusOK, resp)
}
