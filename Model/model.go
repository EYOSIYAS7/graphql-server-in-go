package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Id int 	`json:"id"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}
