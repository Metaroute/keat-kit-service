package http

import (
	"github.com/gimmickless/keat-kit-service/internal/app"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type HTTPHandler struct {
	logger    *zap.SugaredLogger
	catgsrv   *app.CategoryService
	ingredsrv *app.IngredientService
	kitsrv    *app.KitService
}

// NewHTTPHandler constructs a new HTTPHandler.
func NewHTTPHandler(
	logger *zap.SugaredLogger,
	catgsrv *app.CategoryService,
	ingredsrv *app.IngredientService,
	kitsrv *app.KitService,
) *HTTPHandler {
	return &HTTPHandler{logger, catgsrv, ingredsrv, kitsrv}
}

// Category handlers
func (h *HTTPHandler) GetCategories(c *fiber.Ctx) error {
	return c.JSON([]categoryResp{})
}

func (h *HTTPHandler) GetCategory(c *fiber.Ctx) error {
	return c.JSON(categoryResp{})
}

func (h *HTTPHandler) SuggestCategory(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": "1",
	})
}

func (h *HTTPHandler) UpdateCategory(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(categoryResp{})
}

func (h *HTTPHandler) UploadCategoryImage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(categoryResp{})
}

func (h *HTTPHandler) DeleteCategory(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// Ingredient handlers
func (h *HTTPHandler) GetIngredients(c *fiber.Ctx) error {
	return c.JSON([]ingredientResp{})
}

func (h *HTTPHandler) GetIngredient(c *fiber.Ctx) error {
	return c.JSON(ingredientResp{})
}

func (h *HTTPHandler) SuggestIngredient(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": "1",
	})
}

func (h *HTTPHandler) UpdateIngredient(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(ingredientResp{})
}

func (h *HTTPHandler) UploadIngredientImage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(categoryResp{})
}

func (h *HTTPHandler) DeleteIngredient(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// Kit handlers
func (h *HTTPHandler) GetKits(c *fiber.Ctx) error {
	return c.JSON([]kitResp{})
}

func (h *HTTPHandler) GetKit(c *fiber.Ctx) error {
	return c.JSON(kitResp{})
}

func (h *HTTPHandler) CreateKit(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": "1",
	})
}

func (h *HTTPHandler) UpdateKit(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(kitResp{})
}

func (h *HTTPHandler) DeleteKit(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
