package handlers

import (
	"github.com/ESPEDUZA/go-IBC4/tree/main/fiberTest/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	h.DB.Find(&users)
	return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	h.DB.Create(&user)
	return c.JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	h.DB.First(&user, id)

	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	h.DB.Save(&user)
	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}
	h.DB.Delete(&user)
	return c.SendStatus(fiber.StatusNoContent)
}
