package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	ImdbID      string `json:"imdbID"`
	Title       string `json:"Title"`
	Year        string `json:"Year"`
	Ratting     string `json:"ratting"`
	IsSuperHero string `json:"isSuperHero"`
}

var movies = []Movie{}

func getAllMovies(c echo.Context) error {
	return c.JSON(http.StatusOK, movies)
}

func main() {
	e := echo.New()

	e.GET("/api/movies", getAllMovies)

	e.Logger.Fatal(e.Start(":8080"))
}
