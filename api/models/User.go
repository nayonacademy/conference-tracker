package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	FullName string `gorm:"size:255;not null" json:"fullname"`
	Email string `gorm:"size:255; not null; unique" json:"email"`
	Password string `gorm:"size:100; not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

func Hash(password string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost), nil
}

func VerifyPassword(hashedPassword, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}

func (u *User) BeforeSave() error{
	hashedPassword, err := Hash(u.Password)
	if err != nil{
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare(){
	u.ID = 0
	u.FullName = html.EscapeString(strings.TrimSpace(u.FullName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdateAt = time.Now()
}

func (u *User) Validate(action string) error{
	switch  strings.ToLower(action) {
	case "update":
		if u.FullName == ""{
			return errors.New("required fullname")
		}
		if u.Email == ""{
			return errors.New("required email")
		}
		if u.Password == ""{
			return errors.New("required password")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil{
			return errors.New("invalid email")
		}
		return nil
	case "login":
		if u.Email == ""{
			return errors.New("required email")
		}
		if u.Password == ""{
			return errors.New("required password")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil{
			return errors.New("invalid email")
		}
		return nil
	default:
		if u.FullName == ""{
			return errors.New("required fullname")
		}
		if u.Email == ""{
			return errors.New("required email")
		}
		if u.Password == ""{
			return errors.New("required password")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil{
			return errors.New("invalid email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB)(*User, error){
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil{
		return &User{}, nil
	}
	return u, nil
}

func (u *User)FindAllUsers(db *gorm.DB)(*[]User, error){
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil{
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32)(*User, error){
	var err error
	err = db.Debug().Model(User{}).Where("id= ?", uid).Take(&u).Error
	if err !=nil{
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err){
		return &User{}, errors.New("user not found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32)(*User, error) {
	return u , nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32)(int64, error){
	db = db.Debug().Model(&User{}).Where("id=?",uid).Take(&User{}).Delete(&User{})
	if db.Error != nil{
		return 0, db.Error
	}
	return db.RowsAffected, nil
}