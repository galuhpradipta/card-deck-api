package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/galuhpradipta/card-deck-api/handlers"
	"github.com/galuhpradipta/card-deck-api/repositories"
	"github.com/galuhpradipta/card-deck-api/services"
)

var fullCardDecks = []string{
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "TS", "JS", "QS", "KS",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "TD", "JD", "QD", "KD",
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "TC", "JC", "QC", "KC",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "TH", "JH", "QH", "KH",
}

func main() {
	deckRepository := repositories.NewDeckRepository(fullCardDecks)
	deckService := services.NewDeckService(deckRepository)
	deckHandler := handlers.NewDeckHandler(deckService)

	app := fiber.New()
	app.Get("/decks/:id", deckHandler.GetByID)
	app.Post("/decks", deckHandler.Create)
	app.Post("/decks/:id/draw", deckHandler.Draw)

	log.Fatal(app.Listen(":3000"))
}
