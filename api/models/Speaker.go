package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Speaker struct {
	gorm.Model
	Name string `gorm:"size:255; not null" json:"name"`
	Position string `gorm:"size:255; not null" json:"position"`
}

func (s *Speaker) Prepare(){
	s.ID = 0
	s.Name = html.EscapeString(strings.TrimSpace(s.Name))
	s.Position = html.EscapeString(strings.TrimSpace(s.Position))
}
func (s *Speaker) Validate() error{
	if s.Name == ""{
		return errors.New("required name")
	}
	if s.Position == ""{
		return errors.New("required position")
	}
	return nil
}
func (s *Speaker) SaveSpeaker(db *gorm.DB)(*Speaker, error){
	var err error
	err = db.Debug().Model(&Speaker{}).Create(&s).Error
	if err != nil{
		return &Speaker{}, err
	}
	return s, nil
}
func (s *Speaker) FindAllSpeaker(db *gorm.DB)(*[]Speaker, error){
	var err error
	speakers := []Speaker{}
	err = db.Debug().Model(&Speaker{}).Limit(100).Find(&speakers).Error
	if err != nil{
		return &[]Speaker{}, err
	}
	return &speakers, nil
}
func (s *Speaker) FindSpeakerByID(db *gorm.DB, pid uint64)(*Speaker, error){
	var err error
	err = db.Debug().Model(&Speaker{}).Where("id = ?", pid).Take(&s).Error
	if err != nil{
		return &Speaker{}, err
	}
	return s, nil
}
func (s *Speaker) UpdateSpeaker(db *gorm.DB)(*Speaker, error){
	var err error
	err = db.Debug().Model(&Speaker{}).Where("id = ?", s.ID).Updates(Speaker{
		Name:     s.Name,
		Position: s.Position,
	}).Error
	if err != nil{
		return &Speaker{}, err
	}
	return s, nil
}
func (s *Speaker) DeleteSpeaker(db *gorm.DB, pid uint64)(int64, error){
	db = db.Debug().Model(&Speaker{}).Where("id = ?", pid).Take(&Speaker{}).Delete(&Speaker{})
	if db.Error != nil{
		if gorm.IsRecordNotFoundError(db.Error){
			return 0, errors.New("speaker not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}