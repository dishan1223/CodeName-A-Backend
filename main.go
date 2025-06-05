package main

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    // coming from types/models.go 
    "github.com/dishan1223/CodeName-A-Backend/types"
    // coming from controller/movies.go
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
    v1 := app.Group("/api/v1/movies")
    del := app.Group("api/del")

    v1.Get("/", movies.GetMovies)
    v1.Post("/new", movies.AddMovie)
    v1.Put("/update/:id", movies.UpdateMovie)
    del.Delete("/:id", movies.DeleteMovie)

    log.Fatal(app.Listen(":3000"))
}
