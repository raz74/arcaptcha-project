package request

import "time"

type CreateUserRequest struct {
	Password     string `json:"password"`
	Name         string `json:"name"`
	ID           int    `json:"id"`
	Email        string `jason:"email"`
	Phone        string `json:"phone"`
	Company_name string `json:"company_name"`
	Job_title    string `json:"job_title"`
	Active       bool   `json:"active"`
}

type SignupRequest struct {
	Password string `json:"password" validate:"required"`
	Name     string `json:"name"`
	Email    string `jason:"email"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Name     string `json:"name"`
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
