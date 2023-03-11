# Card Deck API
This is a simple Go-based API for a card deck game. It is built using the following:

- Go
- Fiber web framework

## How to Run

To run this application locally, you can follow these steps:

1. Clone this repository:

```bash
git clone <https://github.com/galuhpradipta/card-deck-api.git>

```

1. Install the necessary dependencies:

```bash
go get github.com/gofiber/fiber/v2

```

1. Run the application:

```bash
go run main.go

```

The application should be running on `localhost:3000`.

## API Endpoints

### Create a new deck

```bash
POST /decks
Content-Type: application/json

{
    "shuffle": true,
    "cards": ["AS", "2S", "3S", "4S"]
}

```

### Retrieve a deck by ID

```bash
GET /decks/:id

```

### Draw a card from a deck

```bash
POST /decks/:id/draw
Content-Type: application/json

{
    "count": 1
}

```