package models

type Rating struct {
	ID uint64 `gorm:"primary_key; auto_increment" json:"id"`
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