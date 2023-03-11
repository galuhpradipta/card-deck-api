package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/galuhpradipta/card-deck-api/handlers"
	"github.com/galuhpradipta/card-deck-api/repositories"
	"github.com/galuhpradipta/card-deck-api/services"
	"github.com/galuhpradipta/card-deck-api/shared"
)

func main() {
	deckRepository := repositories.NewDeckRepository(shared.FullCardDecks)
	deckService := services.NewDeckService(deckRepository)
	deckHandler := handlers.NewDeckHandler(deckService)

	app := fiber.New()
	app.Get("/decks/:id", deckHandler.GetByID)
	app.Post("/decks", deckHandler.Create)
	app.Post("/decks/:id/draw", deckHandler.Draw)

	log.Fatal(app.Listen(":3000"))
}
