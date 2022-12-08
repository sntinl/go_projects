package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Price       float64
}
