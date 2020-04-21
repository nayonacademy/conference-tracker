package models

import "github.com/jinzhu/gorm"

type Speaker struct {
	gorm.Model
	Name string
	Position string
}
