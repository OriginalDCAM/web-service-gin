package handlers

import (
	"fmt"
	"strconv"

	"github.com/OrignalDCAM/web-service-gin/services"
	"github.com/gin-gonic/gin"

	"database/sql"
	"net/http"
)

func GetAllGames(c *gin.Context, db *sql.DB) {
	results, err := gamestore.FetchAllGames(db);

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, results)
}

func GetGameByID(c *gin.Context, db *sql.DB) {
	Paramid := c.Param("id")

	id, err := strconv.ParseInt(Paramid, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Id is an invalid character"})
	}

	result, err := gamestore.FetchGameByID(db, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Game not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func CreateNewGame(c *gin.Context, db *sql.DB) {
	var newGame gamestore.Game

	if err := c.BindJSON(&newGame); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
        return
    }

	err := gamestore.CreateGame(db, newGame)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong inserting data into the database"})
		return
	}

	c.IndentedJSON(http.StatusCreated, "created")
}