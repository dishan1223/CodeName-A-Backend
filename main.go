package main

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "github.com/dishan1223/CodeName-A-Backend/types"
    "github.com/dishan1223/CodeName-A-Backend/controllers"
)

var DB *gorm.DB

func main(){
    app := fiber.New()

    var err error
    DB, err = gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    if err := DB.AutoMigrate(&models.Movie{}); err != nil{
        panic("failed to migrate database")
    }

    movies.Init(DB)
    
    // Routes
    app.Get("/api/movies", movies.GetMovies)
    app.Post("/api/movies/new", movies.AddMovie)
    log.Fatal(app.Listen(":3000"))
}
