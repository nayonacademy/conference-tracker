package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Location struct {
	gorm.Model
	Name string
	Date time.Month
	Time time.Time
}
