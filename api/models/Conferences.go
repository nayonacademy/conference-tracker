package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

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

func (c *Conference) Prepare(){
	c.ID = 0
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.Website = html.EscapeString(strings.TrimSpace(c.Website))
	c.About = html.EscapeString(strings.TrimSpace(c.About))
	c.Phone = html.EscapeString(strings.TrimSpace(c.Phone))
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	c.Address = html.EscapeString(strings.TrimSpace(c.Address))
	c.City = html.EscapeString(strings.TrimSpace(c.City))
	c.ZipCode = html.EscapeString(strings.TrimSpace(c.ZipCode))
	c.Facebook = html.EscapeString(strings.TrimSpace(c.Facebook))
	c.Twitter = html.EscapeString(strings.TrimSpace(c.Twitter))
	c.Instagram = html.EscapeString(strings.TrimSpace(c.Instagram))
}

func (c *Conference) Validate() error{
	if c.Name == ""{
		return errors.New("require name")
	}
	if c.Website == ""{
		return errors.New("require website")
	}
	if c.About == ""{
		return errors.New("require about")
	}
	if c.Phone == ""{
		return errors.New("require phone")
	}
	if c.Address == ""{
		return errors.New("require address")
	}
	if c.City == ""{
		return errors.New("require city")
	}
	if c.ZipCode == ""{
		return errors.New("require zip code")
	}
	if c.Facebook == ""{
		return errors.New("require facebook")
	}
	if c.Twitter == ""{
		return errors.New("require twitter")
	}
	if c.Instagram == ""{
		return errors.New("require instagram")
	}
	return nil
}

func (c *Conference) SaveConference(db *gorm.DB)(*Conference, error){
	var err error
	err = db.Debug().Model(&Conference{}).Create(&c).Error
	if err != nil{
		return &Conference{}, err
	}
	return c, nil
}

func (c *Conference) FindAllConference(db *gorm.DB)(*[]Conference, error){
	var err error
	var conferences []Conference
	err = db.Debug().Model(&Conference{}).Find(&conferences).Error
	if err != nil{
		return &[]Conference{}, err
	}
	return &conferences, err
}

func (c *Conference) FindConferenceByID(db *gorm.DB, pid uint64)(*Conference, error){
	var err error
	err = db.Debug().Model(&Conference{}).Where("id = ?", pid).Take(&c).Error
	if err != nil{
		return &Conference{}, err
	}
	return c, nil
}

func (c *Conference) UpdateConference(db *gorm.DB)(*Conference, error){
	var err error
	err = db.Debug().Model(&Conference{}).Where("id = ?", c.ID).Updates(Conference{
		Name:        c.Name,
		Website:     c.Website,
		About:       c.About,
		Phone:       c.Phone,
		Email:       c.Email,
		Address:     c.Address,
		City:        c.City,
		ZipCode:     c.ZipCode,
		Speakers:    c.Speakers,
		Facebook:    c.Facebook,
		Twitter:     c.Twitter,
		Instagram:   c.Instagram,
		OrganizerID: c.OrganizerID,
		Locations:   c.Locations,
		Rating:      c.Rating,
	}).Error
	if err != nil{
		return &Conference{}, err
	}
	return c, nil
}

func (c *Conference) DeleteConference(db *gorm.DB, pid uint64)(int64, error) {
	db = db.Debug().Model(&Conference{}).Where("id = ?", pid).Take(&Conference{}).Delete(&Conference{})
	if db.Error != nil{
		if gorm.IsRecordNotFoundError(db.Error){
			return 0, errors.New("conference not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}