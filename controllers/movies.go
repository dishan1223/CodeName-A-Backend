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
func GetMovies(c *fiber.Ctx) error {
    var movies []models.Movie

    if err := DB.Find(&movies).Error; err != nil{
        return c.Status(500).SendString("Could not retrieve movies")
    }

    return c.JSON(movies)
}


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
