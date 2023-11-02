package gamestore

import (
	"context"
	"database/sql"
	"errors"
)

type Game struct {
	ID        string   `json:"id"`
	Title     string  `json:"title"`
	Developer string  `json:"developer"`
	Price     float64 `json:"price"`
}

var (
	ctx = context.Background()
)

func CreateGame(db *sql.DB, game Game) error {
    _, err := db.ExecContext(ctx,
        "INSERT INTO games (title, developer, price) VALUES ($1, $2, $3)",
        game.Title, game.Developer, game.Price,
    )

    return err
}

func FetchAllGames(db *sql.DB) ([]Game, error) {
	rows, err := db.Query("SELECT * FROM games")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var games []Game
    for rows.Next() {
        var game Game
        if err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Price); err != nil {
            return nil, err
        }
        games = append(games, game)
    }

    return games, nil
}

func FetchGameByID(db *sql.DB, id int64) (*Game, error) {
	var game Game
    err := db.QueryRow("SELECT * FROM games WHERE id = $1", id).
        Scan(&game.ID, &game.Title, &game.Developer, &game.Price)

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil // Game not found, return nil without error
        }
        return nil, err
    }

    return &game, nil
}

func DeleteGame(db *sql.DB, id int64) error {
    _, err := db.ExecContext(context.Background(), "DELETE FROM games WHERE id = $1", id)
    return err
}

func UpdateGame(db *sql.DB, game Game) error {
    _, err := db.ExecContext(context.Background(),
        "UPDATE games SET title = $2, developer = $3, price = $4 WHERE id = $1",
        game.ID, game.Title, game.Developer, game.Price,
    )

    return err
}