package data

import (
	"backend/domain"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Birthdate time.Time `json:"birthdate" validate:"required"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:        int(u.ID),
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Birthdate: u.Birthdate,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User

	res.Firstname = data.Firstname
	res.Lastname = data.Lastname
	res.Username = data.Username
	res.Email = data.Email
	res.Password = data.Password
	res.Birthdate = data.Birthdate

	return res
}
