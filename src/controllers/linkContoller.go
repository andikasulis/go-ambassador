package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-ambassador/src/database"
	"go-ambassador/src/models"
	"strconv"
)

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var links []models.Link

	database.DB.Where("user_id = ?", id).Find(&links)

	return c.JSON(links)
}
