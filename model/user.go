package model

import (
	"time"
	// "role.go"
)

type User struct {
	Username                 string    `json:"username"`
	Password                 string    `json:"password"`
	Name                     string    `json:"name"`
	ID                       int       `json:"id"`
	// Role                     Role      `json:"role"`
	Email                    string    `jason:"email"`
	Phone                    string       `json:"phone"`
	Company_name             string    `json:"company_name"`
	Job_title                string    `json:"job_title"`
	Active                   bool      `json:"active"`
	Subscribe_news          string    `json:"subscribe_news"`
	// Subscribe_notifications []string  `json:"subscribe_notifications"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}


type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Email    string `json:"email"`
}


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}