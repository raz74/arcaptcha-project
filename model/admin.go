package model

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required len gt 2"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
}
