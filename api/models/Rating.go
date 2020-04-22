package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"reflect"
	"strings"
)

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

func (r *Rating) Prepared(){
	r.ID = 0
	r.Comment = html.EscapeString(strings.TrimSpace(r.Comment))
	r.Caption = html.EscapeString(strings.TrimSpace(r.Caption))
}

func (r *Rating) Validate() error {
	if reflect.TypeOf(r.Rate).Kind() !=  reflect.Int{
		return errors.New("rate should be int")
	}
	if r.Comment == ""{
		return errors.New("comments should not be empty")
	}
	if r.Caption == ""{
		return errors.New("caption should not be empty")
	}
	if reflect.TypeOf(r.AttendQ).Kind() !=  reflect.Bool{
		return errors.New("attend should be bool")
	}
	if reflect.TypeOf(r.EnjoyQ).Kind() !=  reflect.Bool{
		return errors.New("enjoy should be bool")
	}
	if reflect.TypeOf(r.LocationQ).Kind() !=  reflect.Bool{
		return errors.New("location should be bool")
	}
	if reflect.TypeOf(r.ConnectPeerQ).Kind() !=  reflect.Bool{
		return errors.New("connectpeer should be bool")
	}
	if reflect.TypeOf(r.Favorite).Kind() !=  reflect.Bool{
		return errors.New("favourite should be bool")
	}
	if reflect.TypeOf(r.Like).Kind() !=  reflect.Bool{
		return errors.New("like should be bool")
	}
	if reflect.TypeOf(r.Awesome).Kind() !=  reflect.Int{
		return errors.New("awesome should be int")
	}
	if reflect.TypeOf(r.Great).Kind() !=  reflect.Int{
		return errors.New("great should be int")
	}
	if reflect.TypeOf(r.Average).Kind() !=  reflect.Int{
		return errors.New("average should be int")
	}
	if reflect.TypeOf(r.Poor).Kind() !=  reflect.Int{
		return errors.New("poor should be int")
	}
	if reflect.TypeOf(r.Terrible).Kind() !=  reflect.Int{
		return errors.New("terrible should be int")
	}
	return nil
}

func (r *Rating) SaveRating(db *gorm.DB)(*Rating, error){
	var err error
	err = db.Debug().Model(&Rating{}).Create(&r).Error
	if err != nil{
		return &Rating{}, err
	}
	return r, err
}

func (r *Rating) FindAllRating(db *gorm.DB)(*[]Rating, error){
	var err error
	var ratings []Rating
	err = db.Debug().Model(&Rating{}).Limit(100).Find(&ratings).Error
	if err != nil{
		return &[]Rating{}, err
	}
	return &ratings, err
}

func (r *Rating) FindRatingByID(db *gorm.DB, pid uint64)(*Rating, error){
	var err error
	err = db.Debug().Model(&Rating{}).Where("id = ?", pid).Take(&r).Error
	if err != nil{
		return &Rating{}, err
	}
	return r, err
}

func (r *Rating) UpdateRating(db *gorm.DB)(*Rating, error){
	var err error
	err = db.Debug().Model(&Rating{}).Where("id = ?", r.ID).Updates(Rating{

		Rate:         r.Rate,
		Comment:      r.Comment,
		Image:        r.Image,
		Caption:      r.Caption,
		AttendQ:      r.AttendQ,
		EnjoyQ:       r.EnjoyQ,
		LocationQ:    r.LocationQ,
		ConnectPeerQ: r.ConnectPeerQ,
		Awesome:      r.Awesome,
		Great:        r.Great,
		Average:      r.Average,
		Poor:         r.Poor,
		Terrible:     r.Terrible,
		Favorite:     r.Favorite,
		Like:         r.Like,
	}).Error
	if err != nil{
		return &Rating{}, err
	}
	return r, nil
}

func (r *Rating) DeleteRating(db *gorm.DB, pid uint64)(int64, error){
	db = db.Debug().Model(&Rating{}).Where("id = ?", pid).Take(&Rating{}).Delete(&Rating{})
	if db.Error != nil{
		if gorm.IsRecordNotFoundError(db.Error){
			return 0, errors.New("rating not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}