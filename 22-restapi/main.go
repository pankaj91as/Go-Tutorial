package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// create movieDB slice variable from Movies struct
var movieDB []Movies

// Movies struct
type Movies struct {
	Id          int      `json:"movie_number" validate:"required"`
	Name        string   `json:"movie_name" validate:"required"`
	Rating      float32  `json:"rating" validate:"required"`
	Genres      []string `json:"genres" validate:"required"`
	ReleaseDate string   `json:"movie_release_date" validate:"required"`
	Poster      string   `json:"movie_banner" validate:"required"`
}

var validate = validator.New()

func main() {
	fmt.Println("Design REST API In GoLang!")

	app := fiber.New()

	app.Get("/", getHomePath)

	// Get list of movies from DB
	app.Get("/movie/list", getMovieList)

	// Insert Movie Into DB
	app.Post("/movie/list", insertMovie)

	// Delete Movie From DB
	app.Delete("/movie/:id", deleteMovie)

	// Update Movie In DB
	app.Patch("/movie/update", updateMovie)

	app.Listen(":4000")
}

func getHomePath(c *fiber.Ctx) error {
	return c.SendString("Welcome to movies library!")
}

func insertMovie(c *fiber.Ctx) error {
	var parsedMovie Movies
	if err := c.BodyParser(&parsedMovie); err != nil {
		panic(err)
	}
	movieDB = append(movieDB, parsedMovie)
	return c.JSON(movieDB)
}

func getMovieList(c *fiber.Ctx) error {
	// Generate Dummy Data
	// moviesData := []Movies{
	// 	{Name: "Terminater 1", Rating: 5, Genres: []string{"action", "animation"}, ReleaseDate: "2024-02-11", Poster: "https://m.media-amazon.com/images/I/71+Nertw2+S._SY741_.jpg"},
	// }
	// return c.JSON(moviesData)

	// Return DB Values
	return c.JSON(movieDB)
}

func deleteMovie(c *fiber.Ctx) error {
	parameters := c.AllParams()
	movieID, _ := strconv.Atoi(parameters["id"])
	for key, data := range movieDB {
		if data.Id == movieID {
			movieDB = append(movieDB[:key], movieDB[key+1:]...)
		}
	}

	return c.JSON(movieDB)
}

func validateMovieJson(body []byte) bool {
	var inputJson Movies
	if err := json.Unmarshal(body, &inputJson); err != nil {
		return false
	}

	if err := validate.Struct(inputJson); err != nil {
		return false
	}

	return true
}

func updateMovie(c *fiber.Ctx) error {
	if !validateMovieJson(c.Body()) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	var updatedMovieData Movies
	if err := c.BodyParser(&updatedMovieData); err != nil {
		log.Error(err)
	}
	movieId := updatedMovieData.Id

	index := -1
	for i, data := range movieDB {
		if data.Id == movieId {
			index = i
		}
	}

	movieDB[index] = updatedMovieData
	return c.JSON(movieDB)
}

/*
Use Below JSON as Request Body to insert Movies

{
    "movie_number": 1,
    "movie_name": "Terminater 1",
    "rating": 5,
    "genres": [
        "action",
        "animation"
    ],
    "movie_release_date": "2024-02-11",
    "movie_banner": "https://m.media-amazon.com/images/I/71+Nertw2+S._SY741_.jpg"
}

Do GET Movies
DO DELETE Movies

Play with it...
*/
