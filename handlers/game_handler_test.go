package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllGames(t *testing.T) {
	// Set mode to test
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Create mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	r.GET("/games", func (c *gin.Context) {
		GetAllGames(c, db)
	})

	rows := sqlmock.NewRows([]string{"id", "title", "developer", "price"}).
							AddRow(1, "Game 1", "Developer 1", 29.99).
							AddRow(2, "Game 2", "Developer 2", 39.99)

	mock.ExpectQuery("SELECT \\* FROM games").WillReturnRows(rows)

	req, _ := http.NewRequest("GET", "/games", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.NoError(t, mock.ExpectationsWereMet())
}