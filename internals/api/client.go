package api

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/rpstvs/steamprice-api/internals/database"
)

type Client struct {
	httpClient http.Client
	DB         *database.Queries
}

func NewClient(timeout time.Duration) Client {

	dbURL := os.Getenv("DATABASE_URL")
	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	return Client{
		httpClient: http.Client{Timeout: timeout},
		DB:         dbQueries,
	}
}
