package models 

import "gorm.io/gorm"

type Movie struct {
    gorm.Model
    Title       string  `json:"title" validate:"required"`
    Description string  `json:"description" validate:"required"`
    Rating      float64 `json:"rating" validate:"required,gte=0,lte=10"`
    MovieDRT    string  `json:"movieDRT" validate:"required"`
    Genre       string  `json:"genre" validate:"required"`
    Poster      string  `json:"poster" validate:"required"`
    Background  string  `json:"background" validate:"required"`
}
