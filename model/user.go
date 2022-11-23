package model

import "time"

type User struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	ID        int       `json:"id"`
	Email     string    `jason:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}