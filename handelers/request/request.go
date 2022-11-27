package handelers

import "time"


type UserRequest struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	ID       int    `json:"id"`
	// Role                     Role      `json:"role"`
	Email        string `jason:"email"`
	Phone        string `json:"phone"`
	Company_name string `json:"company_name"`
	Job_title    string `json:"job_title"`
	Active       bool   `json:"active"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	ID       int    `json:"id"`
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

// build a constructor  for the error
