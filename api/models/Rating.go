package models

type Rating struct {
	ID uint64 `gorm:"primary_key; auto_increment" json:"id"`
	Rate int16 `gorm:"size:255; not null" json:"rate"`
	Comment string `gorm:"size:255; not null" json:"comment"`
	Image *string `gorm:"column:image"`
	Caption string `gorm:"size:255; not null" json:"caption"`
	AttendQ bool `json:"attend_q" sql:"DEFAULT:true;index" gorm:"not null"`
	EnjoyQ bool `json:"enjoy_q" sql:"DEFAULT:true;index" gorm:"not null"`
	LocationQ bool `json:"location_q" sql:"DEFAULT:true;index" gorm:"not null"`
	ConnectPeerQ bool `json:"connect_peer_q" sql:"DEFAULT:true;index" gorm:"not null"`
	Awesome int16
	Great int16
	Average int16
	Poor int16
	Terrible int16
	Favorite bool
	Like bool
}

func (r *Rating) Prepared(){

}

func (r *Rating) Validate(){

}

func (r *Rating) SaveRating(){

}

func (r *Rating) FindAllRating(){

}

func (r *Rating) FindRatingByID(){

}

func (r *Rating) UpdateRating(){

}

func (r *Rating) DeleteRating(){

}