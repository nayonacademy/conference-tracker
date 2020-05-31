package conference

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ConfOwnProfile struct {
	gorm.Model
	Reviews int16
	Reads int16
	Useful int16
	Attend int16
	Favourite string
	Picture string
}

type Conference struct {
	gorm.Model
	Name string `gorm:"size:255; not null" json:"name"`
	Website string	`gorm:"size:255; not null" json:"website"`
	About string	`gorm:"size:255; not null" json:"about"`
	Phone string	`gorm:"size:255; not null" json:"phone"`
	Email string	`gorm:"size:255; not null" json:"email"`
	Address string	`gorm:"size:255; not null" json:"address"`
	City string	`gorm:"size:255; not null" json:"city"`
	ZipCode string	`gorm:"size:255; not null" json:"zip_code"`
	Speakers []Speaker	`gorm:"size:255; not null" json:"speakers"`
	Facebook string	`gorm:"size:255; not null" json:"facebook"`
	Twitter string	`gorm:"size:255; not null" json:"twitter"`
	Instagram string	`gorm:"size:255; not null" json:"instagram"`
	OrganizerID int16	`gorm:"size:255; not null" json:"organizer_id"`
	Locations []Location `gorm:"size:255; not null" json:"locations"`
	Rating Rating	`gorm:"size:255; not null" json:"rating"`
}

type Location struct {
	gorm.Model
	Name string
	Date time.Month
	Time time.Time
}

type Rating struct {
	gorm.Model
	Rate int16 `gorm:"not null" json:"rate"`
	Comment string `gorm:"size:255; not null" json:"comment"`
	Image *string `gorm:"column:image"`
	Caption string `gorm:"size:255; not null" json:"caption"`
	AttendQ bool `json:"attend_q" sql:"DEFAULT:true;index" gorm:"not null"`
	EnjoyQ bool `json:"enjoy_q" sql:"DEFAULT:true;index" gorm:"not null"`
	LocationQ bool `json:"location_q" sql:"DEFAULT:true;index" gorm:"not null"`
	ConnectPeerQ bool `json:"connect_peer_q" sql:"DEFAULT:true;index" gorm:"not null"`
	Awesome int16 `gorm:"not null" json:"awesome"`
	Great int16 `gorm:"not null" json:"great"`
	Average int16 `gorm:"not null" json:"average"`
	Poor int16 `gorm:"not null" json:"poor"`
	Terrible int16 `gorm:"not null" json:"terrible"`
	Favorite bool `json:"favorite" sql:"DEFAULT:true;index" gorm:"not null"`
	Like bool `json:"like" sql:"DEFAULT:true;index" gorm:"not null"`
}
type Report struct {
	gorm.Model
	Offensive bool
	Violence bool
	Spam bool
	InAppropriate bool
}

type Speaker struct {
	gorm.Model
	Name string `gorm:"size:255; not null" json:"name"`
	Position string `gorm:"size:255; not null" json:"position"`
}