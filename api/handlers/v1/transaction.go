package v1

import (
	pb "api_gateway/genproto/budgeting_service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CreateTransaction godoc
// @Security        ApiKeyAuth
// @Router          /transactions/create [post]
// @Description     Creates a new transaction
// @Tags            transactions
// @Accept          json
// @Produce         json
// @Param           body body budgeting_service.CreateTransaction true "Transaction Creation Request"
// @Success         201 {object} budgeting_service.Transaction "Transaction created successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) CreateTransaction(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()

	req := pb.CreateTransaction{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", http.StatusBadRequest, err.Error())
	}

	res, err := h.services.TransactionService().Create(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while creating transaction", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Transaction successfully created", http.StatusCreated, res)
}

// GetTransactionById godoc
// @Security        ApiKeyAuth
// @Router          /transactions/{id} [get]
// @Description     Retrieves a transaction by ID
// @Tags            transactions
// @Accept          json
// @Produce         json
// @Param           id path string true "Transaction ID"
// @Success         200 {object} budgeting_service.Transaction "Transaction retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetTransactionById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := &pb.PrimaryKey{Id: id}
	res, err := h.services.TransactionService().GetById(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving transaction by ID", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Transaction successfully retrieved", http.StatusOK, res)
}

// GetAllTransactions godoc
// @Security        ApiKeyAuth
// @Router          /transactions/all [get]
// @Description     Retrieves all transactions
// @Tags            transactions
// @Accept          json
// @Produce         json
// @Param           page query int true "Page number"
// @Param           limit query int true "Limit number"
// @Param           user_id query string false "User ID"
// @Param           amount query float64 false "Amount"
// @Param           type query string false "Type"
// @Success         200 {object} budgeting_service.Transactions "Transactions retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetAllTransactions(ctx *fiber.Ctx) error {
	req := &pb.TransactionFilter{}
	req.Page = int32(ctx.QueryInt("page", 1))
	req.Limit = int32(ctx.QueryInt("limit", 10))
	req.UserId = ctx.Query("user_id")
	if amount := ctx.QueryInt("amount", 0); amount > 0 {
		req.Amount = float64(amount)
	}
	req.Type = ctx.Query("type")

	res, err := h.services.TransactionService().GetAll(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving transactions", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Transactions successfully retrieved", http.StatusOK, res)
}

// UpdateTransaction godoc
// @Security        ApiKeyAuth
// @Router          /transactions/{id}/update [put]
// @Description     Updates an existing transaction
// @Tags            transactions
// @Accept          json
// @Produce         json
// @Param           id path string true "Transaction ID"
// @Param           body body budgeting_service.Transaction true "Transaction Update Request"
// @Success         200 {object} budgeting_service.Transaction "Transaction updated successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) UpdateTransaction(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()

	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := pb.Transaction{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", http.StatusBadRequest, err.Error())
	}
	req.Id = id

	res, err := h.services.TransactionService().Update(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while updating transaction", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Transaction successfully updated", http.StatusOK, res)
}

// DeleteTransaction godoc
// @Security        ApiKeyAuth
// @Router          /transactions/{id}/delete [delete]
// @Description     Deletes a transaction by ID
// @Tags            transactions
// @Accept          json
// @Produce         json
// @Param           id path string true "Transaction ID"
// @Success         200 {string} string  "Transaction deleted successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) DeleteTransaction(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req := &pb.PrimaryKey{Id: id}
	_, err := h.services.TransactionService().Delete(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while deleting transaction", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Transaction successfully deleted", http.StatusOK, nil)
}
