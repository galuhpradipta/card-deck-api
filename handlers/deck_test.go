package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/galuhpradipta/card-deck-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

const timeout = 5000

func Test_deckHandler_Create(t *testing.T) {
	tests := []struct {
		desc         string
		route        string
		method       string
		payload      string
		expectedCode int
	}{
		{
			desc:         "Test deckHandler.Create with valid request",
			route:        "/decks",
			method:       fiber.MethodPost,
			payload:      ``,
			expectedCode: fiber.StatusOK,
		},
	}

	s := services.NewMockDeckService()
	h := NewDeckHandler(s)

	app := fiber.New()
	app.Post("/decks", h.Create)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.route, nil)
			resp, _ := app.Test(req, timeout)
			assert.Equal(t, test.expectedCode, resp.StatusCode)
		})
	}
}

func Test_deckHandler_GetByID(t *testing.T) {
	tests := []struct {
		desc         string
		route        string
		method       string
		payload      string
		expectedCode int
	}{
		{
			desc:         "Test deckHandler.GetByID with valid request",
			route:        "/decks/1",
			method:       fiber.MethodGet,
			payload:      ``,
			expectedCode: fiber.StatusOK,
		},
	}

	s := services.NewMockDeckService()
	h := NewDeckHandler(s)

	app := fiber.New()
	app.Get("/decks/:id", h.GetByID)

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.route, nil)
			resp, _ := app.Test(req, timeout)
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
		})
	}
}

func Test_deckHandler_Draw(t *testing.T) {
	tests := []struct {
		desc         string
		route        string
		method       string
		payload      drawDeckRequest
		expectedCode int
	}{
		{
			desc:         "Test deckHandler.Draw with valid request",
			route:        "/decks/1/draw",
			method:       fiber.MethodPost,
			payload:      drawDeckRequest{Count: 1},
			expectedCode: fiber.StatusOK,
		},
	}

	s := services.NewMockDeckService()
	h := NewDeckHandler(s)

	app := fiber.New()
	app.Post("/decks/:id/draw", h.Draw)

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			payload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(tt.method, tt.route, bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req, timeout)
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
		})
	}
}
