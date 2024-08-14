package v1

import (
	pb "api_gateway/genproto/budgeting_service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CreateBudget godoc
// @Security        ApiKeyAuth
// @Router          /budgets/create [post]
// @Description     Creates a new budget
// @Tags            budgets
// @Accept          json
// @Produce         json
// @Param           body body budgeting_service.CreateBudget true "Budget Creation Request"
// @Success         201 {object} budgeting_service.Budget "Budget created successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) CreateBudget(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		return handleResponse(ctx, h.log, "error while getting user info from token", 401, err.Error())
	}

	req := pb.CreateBudget{}
	err = ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", 404, err.Error())
	}
	req.UserId = user.Id

	_, err = time.Parse("04-05-2006", req.StartDate)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing start_date", 404, err.Error())
	}
	_, err = time.Parse("04-05-2006", req.EndDate)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing end_date", 404, err.Error())
	}

	res, err := h.services.BudgetService().Create(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while creating budget", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Budget successfully created", 201, res)
}

// GetBudgetById godoc
// @Security        ApiKeyAuth
// @Router          /budgets/{id} [get]
// @Description     Retrieves a budget by ID
// @Tags            budgets
// @Accept          json
// @Produce         json
// @Param           id path string true "Budget ID"
// @Success         200 {object} budgeting_service.Budget "Budget retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetBudgetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := &pb.PrimaryKey{Id: id}
	res, err := h.services.BudgetService().GetById(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving budget by ID", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Budget successfully retrieved", 200, res)
}

// GetAllBudgets godoc
// @Security        ApiKeyAuth
// @Router          /budgets/all [get]
// @Description     Retrieves all budgets
// @Tags            budgets
// @Accept          json
// @Produce         json
// @Param           page query int true "Page number"
// @Param           limit query int true "Limit number"
// @Param           user_id query string false "User ID"
// @Param           category_id query string false "category_id"
// @Param           period query string false "Name"
// @Param           type query string false "Type"
// @Param           amount query float64 false "Amount"
// @Param           start_date query string false "start_date"
// @Param           end_date query string false "end_date"
// @Success         200 {object} budgeting_service.Budgets "Budgets retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetAllBudgets(ctx *fiber.Ctx) error {
	req := &pb.BudgetFilter{}
	req.Page = int32(ctx.QueryInt("page", 1))
	req.Limit = int32(ctx.QueryInt("limit", 10))
	req.UserId = ctx.Query("user_id")
	req.Amount = ctx.QueryFloat("amount")
	req.UserId = ctx.Query("user_id")
	req.CategoryId = ctx.Query("category_id")
	req.Period = ctx.Query("period")
	req.StartDate = ctx.Query("start_date")
	req.EndDate = ctx.Query("end_date")

	res, err := h.services.BudgetService().GetAll(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving budgets", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Budgets successfully retrieved", 200, res)
}

// UpdateBudget godoc
// @Security        ApiKeyAuth
// @Router          /budgets/{id}/update [put]
// @Description     Updates an existing budget
// @Tags            budgets
// @Accept          json
// @Produce         json
// @Param           id path string true "Budget ID"
// @Param           body body budgeting_service.Budget true "Budget Update Request"
// @Success         200 {object} budgeting_service.Budget "Budget updated successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) UpdateBudget(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := pb.Budget{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", 404, err.Error())
	}
	req.Id = id
	data, err := json.Marshal(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while marshalling request", 404, err.Error())
	}

	err = h.iKafka.ProduceMessage("budget_updated", string(data))
	if err != nil {
		return handleResponse(ctx, h.log, "Error while sending message", http.StatusInternalServerError, err.Error())
	}

	// res, err := h.services.BudgetService().Update(reqCtx, &req)
	// if err != nil {
	// 	return handleResponse(ctx, h.log, "Error while updating budget", http.StatusInternalServerError, err.Error())
	// }

	return handleResponse(ctx, h.log, "Budget successfully updated", 200, nil)
}

// DeleteBudget godoc
// @Security        ApiKeyAuth
// @Router          /budgets/{id}/delete [delete]
// @Description     Deletes a budget by ID
// @Tags            budgets
// @Accept          json
// @Produce         json
// @Param           id path string true "Budget ID"
// @Success         200 {string} string  "Budget deleted successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) DeleteBudget(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req := &pb.PrimaryKey{Id: id}
	_, err := h.services.BudgetService().Delete(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while deleting budget", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Budget successfully deleted", 200, nil)
}
