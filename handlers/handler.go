package handlers

import "github.com/gofiber/fiber/v2"

type DeckHandler interface {
	Create(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Draw(c *fiber.Ctx) error
}
