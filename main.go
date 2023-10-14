package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	engine "github.com/OrignalDCAM/web-service-gin/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Game struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	Developer string  `json:"developer"`
	Price     float64 `json:"price"`
}

func getAllGames(c *gin.Context) {
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
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		games = append(games, game)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, games)
}

func getGameByID(c *gin.Context) {
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

var db *sql.DB

func main() {
	godotenv.Load()

	connStr := "user=" + os.Getenv("DB_USERNAME") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" host=127.0.0.1 port=5432 sslmode=disable"

	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	engine.Run("8001")

	// router := gin.Default()
	// router.GET("/games", getAllGames)
	// router.GET("/game/:id", getGameByID)

	// router.Run("localhost:8001")
}
