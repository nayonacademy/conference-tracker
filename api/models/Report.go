package models

import "github.com/jinzhu/gorm"

type Report struct {
	gorm.Model
	Offensive bool
	Violence bool
	Spam bool
	InAppropriate bool
}