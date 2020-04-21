package models

import "github.com/jinzhu/gorm"

type ConfOwnProfile struct {
	gorm.Model
	Reviews int16
	Reads int16
	Useful int16
	Attend int16
	Favourite string
	Picture string
}