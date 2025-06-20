package movies

import (
    "github.com/gofiber/fiber/v2"
    "github.com/go-playground/validator/v10"
    "gorm.io/gorm"
    "github.com/dishan1223/CodeName-A-Backend/types"
)

var DB *gorm.DB

func Init(db *gorm.DB) {
    DB = db
}

// controllers
// respond with the full movie database
func GetMovies(c *fiber.Ctx) error {
    var movies []models.Movie

    if err := DB.Find(&movies).Error; err != nil{
        return c.Status(500).SendString("Could not retrieve movies")
    }

    return c.JSON(movies)
}

// Respond with a status code for creating a new movie
func AddMovie(c *fiber.Ctx) error {
    var movie models.Movie

    if err := c.BodyParser(&movie); err != nil{
        return c.Status(400).SendString("Could not parse JSON")
    }

    validate := validator.New()
    if err := validate.Struct(&movie); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    } 
    
    if err := DB.Create(&movie).Error; err != nil{
        return c.Status(500).SendString("Could not create movie")
    }

    return c.JSON(movie)
}


// Respond with a status code for updating a movie
func UpdateMovie(c *fiber.Ctx) error {
    id := c.Params("id")

    var movie models.Movie
    if err := DB.First(&movie, id).Error; err != nil {
        return c.Status(404).SendString("Movie not found")
    }

    var updatedData models.Movie
    if err := c.BodyParser(&updatedData); err != nil {
        return c.Status(400).SendString("Failed to parse JSON")
    }

    // GORM will only update non-zero fields unless you use `Select` or `map`
    if err := DB.Model(&movie).Updates(updatedData).Error; err != nil {
        return c.Status(500).SendString("Could not update movie")
    }

    return c.JSON(movie)
}

// Respond with a status code for deleting a movie
func DeleteMovie(c *fiber.Ctx) error {
    id := c.Params("id")

    var movie models.Movie
    if err := DB.First(&movie, id).Error; err != nil {
        return c.Status(404).SendString("Movie not found")
    }

    if err := DB.Delete(&movie).Error; err != nil {
        return c.Status(500).SendString("Could not delete movie")
    }

    return c.JSON(fiber.Map{
        "message": "Movie deleted successfully",
        "status" : "success",
    })
}
