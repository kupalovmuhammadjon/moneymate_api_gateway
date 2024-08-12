package v1

import (
	pb "api_gateway/genproto/budgeting_service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CreateCategory godoc
// @Security        ApiKeyAuth
// @Router          /categories/create [post]
// @Description     Creates a new category
// @Tags            categories
// @Accept          json
// @Produce         json
// @Param           body body budgeting_service.CreateCategory true "Category Creation Request"
// @Success         201 {object} budgeting_service.Category "Category created successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) CreateCategory(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		return handleResponse(ctx, h.log, "error while getting user info from token", http.StatusUnauthorized, err.Error())
	}

	req := pb.CreateCategory{}
	err = ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", http.StatusNotFound, err.Error())
	}
	req.UserId = user.Id

	res, err := h.services.CategoryService().Create(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while creating category", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Category successfully created", http.StatusCreated, res)
}

// GetCategoryById godoc
// @Security        ApiKeyAuth
// @Router          /categories/{id} [get]
// @Description     Retrieves a category by ID
// @Tags            categories
// @Accept          json
// @Produce         json
// @Param           id path string true "Category ID"
// @Success         200 {object} budgeting_service.Category "Category retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetCategoryById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := &pb.PrimaryKey{Id: id}
	res, err := h.services.CategoryService().GetById(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving category by ID", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Category successfully retrieved", http.StatusOK, res)
}

// GetAllCategories godoc
// @Security        ApiKeyAuth
// @Router          /categories/all [get]
// @Description     Retrieves all categories
// @Tags            categories
// @Accept          json
// @Produce         json
// @Param           page query int true "Page number"
// @Param           limit query int true "Limit number"
// @Param           user_id query string false "User ID"
// @Param           name query string false "Name"
// @Param           type query string false "Type"
// @Success         200 {object} budgeting_service.Categories "Categories retrieved successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) GetAllCategories(ctx *fiber.Ctx) error {
	req := &pb.CategoryFilter{}
	req.Page = int32(ctx.QueryInt("page", 1))
	req.Limit = int32(ctx.QueryInt("limit", 10))
	req.UserId = ctx.Query("user_id")
	req.Name = ctx.Query("name")
	req.Type = ctx.Query("type")

	res, err := h.services.CategoryService().GetAll(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while retrieving categories", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Categories successfully retrieved", http.StatusOK, res)
}

// UpdateCategory godoc
// @Security        ApiKeyAuth
// @Router          /categories/{id}/update [put]
// @Description     Updates an existing category
// @Tags            categories
// @Accept          json
// @Produce         json
// @Param           id path string true "Category ID"
// @Param           body body budgeting_service.Category true "Category Update Request"
// @Success         200 {object} budgeting_service.Category "Category updated successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) UpdateCategory(ctx *fiber.Ctx) error {
	reqCtx := ctx.Context()

	id := ctx.Params("id")
	if len(id) < 5 {
		return handleResponse(ctx, h.log, "Invalid id", http.StatusBadRequest, nil)
	}

	req := pb.Category{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while parsing body", http.StatusNotFound, err.Error())
	}
	req.Id = id

	res, err := h.services.CategoryService().Update(reqCtx, &req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while updating category", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Category successfully updated", http.StatusOK, res)
}

// DeleteCategory godoc
// @Security        ApiKeyAuth
// @Router          /categories/{id}/delete [delete]
// @Description     Deletes a category by ID
// @Tags            categories
// @Accept          json
// @Produce         json
// @Param           id path string true "Category ID"
// @Success         200 {string} string  "Category deleted successfully"
// @Failure         401 {object} models.Response "Unauthorized"
// @Failure         404 {object} models.Response "Not Found"
// @Failure         500 {object} models.Response "Internal Server Error"
func (h *HandlerV1) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req := &pb.PrimaryKey{Id: id}
	_, err := h.services.CategoryService().Delete(ctx.Context(), req)
	if err != nil {
		return handleResponse(ctx, h.log, "Error while deleting category", http.StatusInternalServerError, err.Error())
	}

	return handleResponse(ctx, h.log, "Category successfully deleted", http.StatusOK, nil)
}
