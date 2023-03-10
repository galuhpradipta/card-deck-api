package handlers

import "github.com/gofiber/fiber/v2"

type DeckHandler interface {
	Create(c *fiber.Ctx) error
}
