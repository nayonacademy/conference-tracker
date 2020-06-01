package conference

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
	"context"
)

type ConfOwnProfile struct {
	gorm.Model
	Reviews int16	`json:"reviews"`
	Reads int16	`json:"reads"`
	Useful int16 `json:"useful"`
	Attend int16	`json:"attend"`
	Favourite string	`json:"favourite"`
	Picture string	`json:"picture"`
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
	Name string	`json:"name"`
	Date time.Month	`json:"date"`
	Time time.Time	`json:"time"`
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
	Offensive bool	`json:"offensive"`
	Violence bool	`json:"violence"`
	Spam bool	`json:"spam"`
	InAppropriate bool	`json:"in_appropriate"`
}

type Speaker struct {
	gorm.Model
	Name string `gorm:"size:255; not null" json:"name"`
	Position string `gorm:"size:255; not null" json:"position"`
}

type Repository interface {
	CreateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile) error
	GetConfOwnProfile(ctx context.Context, id string)(interface{}, error)
	UpdateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteConfOwnProfile(ctx context.Context, id string)(string, error)

	CreateConference(ctx context.Context, conference Conference) (string, error)
	GetConference(ctx context.Context, id string)(interface{}, error)
	UpdateCreateConference(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateConference(ctx context.Context, id string)(string, error)

	CreateLocation(ctx context.Context, location Location) error
	GetLocation(ctx context.Context, id string)(interface{}, error)
	UpdateCreateLocation(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateLocation(ctx context.Context, id string)(string, error)

	CreateRating(ctx context.Context, rating Rating) error
	GetRating(ctx context.Context, id string)(interface{}, error)
	UpdateCreateRating(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateRating(ctx context.Context, id string)(string, error)

	CreateReport(ctx context.Context, report Report) error
	GetReport(ctx context.Context, id string)(interface{}, error)
	UpdateCreateReport(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateReport(ctx context.Context, id string)(string, error)

	CreateSpeaker(ctx context.Context, speaker Speaker) error
	GetSpeaker(ctx context.Context, id string)(interface{}, error)
	UpdateCreateSpeaker(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateSpeaker(ctx context.Context, id string)(string, error)
}

func GetAllUsers () []*Conference{
	accs := []*Conference{}
	err := errors.New("write query here")
	if err != nil{
		panic("error")

	}
	return accs
}