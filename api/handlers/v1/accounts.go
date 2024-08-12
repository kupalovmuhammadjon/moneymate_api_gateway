package v1

import (
	pb "api_gateway/genproto/budgeting_service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CreateAccount godoc
// @Security        ApiKeyAuth
// @Router          /accounts/create [post]
// @Description     Creates an account
// @Tags            accounts
// @Accept          json
// @Produce         json
// @Param           body body budgeting_service.CreateAccount true "Account Creation Request"
// @Success         201 {object} budgeting_service.Account "Account created successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) CreateAccount(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		return handleResponse(ctx, h.log, "error while getting user info from token", 401, err.Error())
	}

	req := pb.CreateAccount{}

	err = ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", 404, err.Error())
	}
	req.UserId = user.Id

	res, err := h.services.AccountService().Create(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while creating account", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Account successfully created", 201, res)
}

// GetById godoc
// @Security        ApiKeyAuth
// @Router          /accounts/{id} [get]
// @Description     Retrieves an account by ID
// @Tags            accounts
// @Accept          json
// @Produce         json
// @Param           id path string true "Account ID"
// @Success         200 {object} budgeting_service.Account "Account retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetAccountById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := &pb.PrimaryKey{Id: id}
	res, err := h.services.AccountService().GetById(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving account by ID", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Account successfully retrieved", 200, res)
}

// GetAll godoc
// @Security        ApiKeyAuth
// @Router          /accounts/all [get]
// @Description     Retrieves all accounts
// @Tags            accounts
// @Accept          json
// @Produce         json
// @Param           page query int true "Page number"
// @Param           limit query int true "Limit number"
// @Param           user_id query string false "User ID"
// @Param           name query string false "Name"
// @Param           type query string false "Type"
// @Param           balance_from query float64 false "Balance from"
// @Param           balance_to query float64 false "Balance to"
// @Param           currency query string false "Currency"
// @Success         200 {object} budgeting_service.Accounts "Accounts retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetAllAccounts(ctx *fiber.Ctx) error {
	req := &pb.AccountFilter{}
	req.Page = int32(ctx.QueryInt("page", 1))
	req.Limit = int32(ctx.QueryInt("limit", 10))
	req.UserId = ctx.Query("user_id")
	req.Name = ctx.Query("name")
	req.Type = ctx.Query("type")
	req.BalanceFrom = ctx.QueryFloat("balance_from", 0)
	req.BalanceTo = ctx.QueryFloat("balance_to", 0)
	req.Currency = ctx.Query("currency")

	res, err := h.services.AccountService().GetAll(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving accounts", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Accounts successfully retrieved", 200, res)
}

// Update godoc
// @Security        ApiKeyAuth
// @Router          /accounts/{id}/update [put]
// @Description     Updates an existing account
// @Tags            accounts
// @Accept          json
// @Produce         json
// @Param           id path string  true "Account id"
// @Param           body body budgeting_service.Account true "Account Update Request"
// @Success         200 {object} budgeting_service.Account "Account updated successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) UpdateAccount(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()
	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := pb.Account{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", 404, err.Error())
	}
	req.Id = id

	res, err := h.services.AccountService().Update(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while updating account", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Account successfully updated", 200, res)
}

// Delete godoc
// @Security        ApiKeyAuth
// @Router          /accounts/{id}/delete [delete]
// @Description     Deletes an account by ID
// @Tags            accounts
// @Accept          json
// @Produce         json
// @Param           id path string true "Account ID"
// @Success         200 {string} string  "Account deleted successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) DeleteAccount(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}
	req := &pb.PrimaryKey{Id: id}
	_, err := h.services.AccountService().Delete(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while deleting account", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Account successfully deleted", 200, nil)
}
