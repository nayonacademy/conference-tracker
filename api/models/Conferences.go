package models

import "github.com/jinzhu/gorm"

type Conference struct {
	gorm.Model
	Name string
	Website string
	About string
	Phone string
	Email string
	Address string
	City string
	ZipCode string
	Speakers []Speaker
	Facebook string
	Twitter string
	Instagram string
	OrganizerID int16
	Locations []Location
}

func (c *Conference) Prepare(){

}

func (c *Conference) Validate(){

}

func (c *Conference) SaveConference(){

}

func (c *Conference) FindAllConference(){

}

func (c *Conference) FindConferenceByID(){

}

func (c *Conference) UpdateConference(){

}

func (c *Conference) DeleteConference() {

}