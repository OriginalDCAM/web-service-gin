package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"database/sql"
)

type Game struct {
	ID        string   `json:"id"`
	Title     string  `json:"title"`
	Developer string  `json:"developer"`
	Price     float64 `json:"price"`
}

func GetAllGames(c *gin.Context, db *sql.DB) {
	var games []Game

	rows, err := db.Query("SELECT * FROM game")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Price); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		games = append(games, game)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, games)
}

func GetGameByID(c *gin.Context, db *sql.DB) {
	var game Game
	id := c.Param("id")

	row := db.QueryRow("SELECT * FROM game WHERE id = $1", id)
	if err := row.Scan(&game.ID, &game.Title, &game.Developer, &game.Price); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Game not found"})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unkown error"})
		return
	}
	c.IndentedJSON(http.StatusOK, game)
}

func CreateNewGame(c *gin.Context, db *sql.DB) {
	var newGame Game

	if err := c.BindJSON(&newGame); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
        return
    }

	c.IndentedJSON(http.StatusCreated, newGame)
}