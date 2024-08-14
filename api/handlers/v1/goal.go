package v1

import (
	pb "api_gateway/genproto/budgeting_service"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CreateGoal godoc
// @Security        ApiKeyAuth
// @Router          /goals/create [post]
// @Description     Creates a new goal
// @Tags            goals
// @Accept          json
// @Produce         json
// @Param           body body budgeting_service.CreateGoal true "Goal Creation Request"
// @Success         201 {object} budgeting_service.Goal "Goal created successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) CreateGoal(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		return handleResponse(ctx, h.log, "error while getting user info from token", http.StatusUnauthorized, err.Error())
	}

	req := pb.CreateGoal{}
	err = ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", http.StatusBadRequest, err.Error())
	}
	req.UserId = user.Id

	res, err := h.services.GoalService().Create(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while creating goal", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Goal successfully created", http.StatusCreated, res)
}

// GetGoalById godoc
// @Security        ApiKeyAuth
// @Router          /goals/{id} [get]
// @Description     Retrieves a goal by ID
// @Tags            goals
// @Accept          json
// @Produce         json
// @Param           id path string true "Goal ID"
// @Success         200 {object} budgeting_service.Goal "Goal retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetGoalById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := &pb.PrimaryKey{Id: id}
	res, err := h.services.GoalService().GetById(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving goal by ID", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Goal successfully retrieved", http.StatusOK, res)
}

// GetAllGoals godoc
// @Security        ApiKeyAuth
// @Router          /goals/all [get]
// @Description     Retrieves all goals
// @Tags            goals
// @Accept          json
// @Produce         json
// @Param           page query int true "Page number"
// @Param           limit query int true "Limit number"
// @Param           user_id query string false "User ID"
// @Param           name query string false "Name"
// @Param           target_amount query float64 false "Target Amount"
// @Param           current_amount query float64 false "Current Amount"
// @Param           deadline query string false "Deadline"
// @Param           status query string false "Status"
// @Success         200 {object} budgeting_service.Goals "Goals retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetAllGoals(ctx *fiber.Ctx) error {
	req := &pb.GoalFilter{}
	req.Page = int32(ctx.QueryInt("page", 1))
	req.Limit = int32(ctx.QueryInt("limit", 10))
	req.UserId = ctx.Query("user_id")
	req.Name = ctx.Query("name")
	req.TargetAmount = ctx.QueryFloat("target_amount", 0)
	req.CurrentAmount = ctx.QueryFloat("current_amount", 0)
	req.Deadline = ctx.Query("deadline")
	req.Status = ctx.Query("status")

	res, err := h.services.GoalService().GetAll(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving goals", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Goals successfully retrieved", http.StatusOK, res)
}

// UpdateGoal godoc
// @Security        ApiKeyAuth
// @Router          /goals/{id}/update [put]
// @Description     Updates an existing goal
// @Tags            goals
// @Accept          json
// @Produce         json
// @Param           id path string true "Goal ID"
// @Param           body body budgeting_service.Goal true "Goal Update Request"
// @Success         200 {object} budgeting_service.Goal "Goal updated successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) UpdateGoal(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := pb.Goal{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", http.StatusBadRequest, err.Error())
	}
	req.Id = id
	data, err := json.Marshal(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while marshalling request", 404, err.Error())
	}

	err = h.iKafka.ProduceMessage("goal_progress_updated", string(data))
	if err != nil {
		return handleResponse(ctx, h.log, "Error while sending message", http.StatusInternalServerError, err.Error())
	}

	// res, err := h.services.GoalService().Update(reqCtx, &req)
	// if err != nil {
	// 	return handleResponse(ctx, h.log, "Error while updating goal", http.StatusInternalServerError, err.Error())
	// }

	return handleResponse(ctx, h.log, "Goal successfully updated", http.StatusOK, nil)
}

// DeleteGoal godoc
// @Security        ApiKeyAuth
// @Router          /goals/{id}/delete [delete]
// @Description     Deletes a goal by ID
// @Tags            goals
// @Accept          json
// @Produce         json
// @Param           id path string true "Goal ID"
// @Success         200 {string} string  "Goal deleted successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) DeleteGoal(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req := &pb.PrimaryKey{Id: id}
	_, err := h.services.GoalService().Delete(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while deleting goal", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Goal successfully deleted", http.StatusOK, nil)
}
