package request

import "time"

type CreateUserRequest struct {
	Password     string `json:"password"`
	Name         string `json:"name" validate:"required len gt 2"`
	Email        string `jason:"email" validate:"required,email"`
	Phone        string `json:"phone" validate:"required,phone eq 11"`
	Company_name string `json:"company_name"`
	Job_title    string `json:"job_title"`
	Active       bool   `json:"active"`
}

type SignupRequest struct {
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required len gt 2"`
	Email    string `jason:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,phone eq 11"`
}

type LoginRequest struct {
	Name     string `json:"name" validate:"required len gt 2"`
	Password string `json:"password"`
}

type HTTPError struct {
	Status int       `json:"status"`
	Msg    string    `json:"massage"`
	Date   time.Time `json:"date"`
}

func NewHTTPEror(status int, msg string) HTTPError {
	return HTTPError{
		Status: status,
		Msg:    msg,
		Date:   time.Now(),
	}
}
