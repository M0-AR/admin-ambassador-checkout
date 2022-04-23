package controllers

import (
	"admin-ambassador-checkout/src/database"
	"admin-ambassador-checkout/src/models"
	"github.com/gofiber/fiber/v2"
)

func Ambassador(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}
