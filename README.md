# Card Deck API
This is a simple Go-based API for a card deck game. It is built using the following:

- Go
- [Fiber](https://github.com/gofiber/fiber)

## How to Run

To run this application locally, you can follow these steps:

## Clone this repository:

```bash 
$ git clone https://github.com/galuhpradipta/card-deck-api

```
## Run the application:

```bash
$ go run main.go

```

The application should be running on `localhost:3000`.

## API Endpoints

### Create a new deck
Parameters:
| name    | type     | default | 
| --------| -------- | --------|
| shuffle | boolean  | false   |
| cards   | []string | empty array provided will build full deck of cards (52)|

Response Fields:
| name      | type    |
| --------  | --------|
| deck_id   | uuid    |
| shuffled  | boolean |
| remaining | int     |

### Example
```bash
POST /decks
Content-Type: application/json
Request 
{
    "shuffle": true,
    "cards": ["AS", "2S", "3S", "4S"]
}


Response 
{
    "deck_id": "bb2836c7-b513-432a-bfaf-63c348ed94ed",
    "shuffled": false,
    "remaining": 4
}
```

### Retrieve a deck by ID
Response Fields:
| name      | type    |
| --------  | --------|
| deck_id   | uuid    |
| shuffled  | boolean |
| remaining | int     |
| cards     | int     |

Example
```bash
GET /decks/:id

Response 
{
    "deck_id": "bb2836c7-b513-432a-bfaf-63c348ed94ed",
    "shuffled": false,
    "remaining": 4,
    "cards": [
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
        {
            "value": "2",
            "suit": "SPADES",
            "code": "2S"
        },
        {
            "value": "3",
            "suit": "SPADES",
            "code": "3S"
        },
        {
            "value": "4",
            "suit": "SPADES",
            "code": "4S"
        }
    ]
}
```

### Draw a card from a deck
Parameters:
| name    | type     |
| --------| -------- |
| count   | int      |


Response Fields:
| name      | type    |
| --------  | --------|
| cards     | []Card  |


Example
```bash
POST /decks/:id/draw
Content-Type: application/json

Request
{
    "count": 2
}

Response
[
    "cards": {
        "value": "ACE",
        "suit": "SPADES",
        "code": "AS"
    },
    {
        "value": "2",
        "suit": "SPADES",
        "code": "2S"
    }
]

```

### Running Tests
```bash
$ go test ./... -v 
```
