package main

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

type Movie struct {
    gorm.Model
    Title string `json:"title"`
    Description string `json:"description"`
    Rating float64 `json:"rating"`
    MovieDRT string `json:"movieDRT"`
    Genre string `json:"genre"`
    Poster string `json:"poster"`
    Background string `json:"background"`
}

// controllers
func getMovies(c *fiber.Ctx) error {
    var movies []Movie

    if err := DB.Find(&movies).Error; err != nil{
        return c.Status(500).SendString("Could not retrieve movies")
    }

    return c.JSON(movies)
}

func AddMovie(c *fiber.Ctx) error {
    var movie Movie

    if err := c.BodyParser(&movie); err != nil{
        return c.Status(400).SendString("Could not parse JSON")
    }

    if err := DB.Create(&movie).Error; err != nil{
        return c.Status(500).SendString("Could not create movie")
    }

    return c.JSON(movie)
}


func main(){
    app := fiber.New()

    var err error
    DB, err = gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    if err := DB.AutoMigrate(&Movie{}); err != nil{
        panic("failed to migrate database")
    }

    
    // Routes
    app.Get("/api/movies", getMovies)
    app.Post("/api/movies/new", AddMovie)
    log.Fatal(app.Listen(":3000"))
}
