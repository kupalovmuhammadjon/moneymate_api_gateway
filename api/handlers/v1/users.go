package v1

import (
	"api_gateway/api/handlers/models"
	pb "api_gateway/genproto/users"
	"api_gateway/pkg/logger"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserProfile  	godoc
// @Router 			/users/profile [get]
// @Summary 		Get User Profile
// @Description 	getting user profile by user id
// @Accept 			json
// @Produce 		json
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) GetUserProfile(ctx *gin.Context) {

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.UsersService().GetUserProfile(ctx, &pb.PrimaryKey{Id: user.Id})
	if err != nil {
		handleResponse(ctx, h.log, "error while using GetUserProfile method of users service", http.StatusInternalServerError, logger.Error(err))
		return
	}

	handleResponse(ctx, h.log, "", http.StatusOK, resp)
}

// UpdateUserProfile  	godoc
// @Router 			/users/update [put]
// @Summary 		Update User Profile
// @Description 	updating user profile
// @Accept 			json
// @Produce 		json
// @Param 			user body models.UpdateUser true "user"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) UpdateUserProfile(ctx *gin.Context) {

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
		return
	}

	req := models.UpdateUser{}
	if err = json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		handleResponse(ctx, h.log, "error while getting request body", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.UsersService().UpdateUserProfile(ctx, &pb.UpdateUser{
		Id:             user.Id,
		FullName:       req.FullName,
		NativeLanguage: req.NativeLanguage,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error while using UpdateUserProfile method of users service", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "", http.StatusOK, resp)
}

// ChangePassword  	godoc
// @Router 			/users/password [put]
// @Summary 		Change User Password
// @Description 	changing user password
// @Accept 			json
// @Produce 		json
// @Param 			change_password body models.ChangePassword true "change_password"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) ChangePassword(ctx *gin.Context) {

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
		return
	}

	req := models.ChangePassword{}
	if err = json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		handleResponse(ctx, h.log, "error while getting request body", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.UsersService().ChangePassword(ctx, &pb.ChangePassword{
		CurrentPassword: req.CurrentPassword,
		NewPassword:     req.NewPassword,
		UserId:          user.Id,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error while using ChangePassword method of users service", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "", http.StatusOK, resp)
}
