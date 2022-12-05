package model

import (
	"Arc/handlers/response"
	"time"
)

type User struct {
	Password                string    `json:"password"`
	Name                    string    `json:"name"`
	Id                      int       `json:"id"`
	Email                   string    `jason:"email"`
	Phone                   string    `json:"phone"`
	CompanyName             string    `json:"company_name"`
	JobTitle                string    `json:"job_title"`
	Active                  bool      `json:"active"`
	Subscribe_news          bool      `json:"subscribe_news"`
	Subscribe_notifications bool      `json:"subscribe_notifications"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	WebSites                []Website `gorm:"constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Plan                    []Plan    `gorm:"many2many:UserPlan ; constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}

func (u *User) ToResponse() *response.UserResponse {
	return &response.UserResponse{
		Name:        u.Name,
		Email:       u.Email,
		Phone:       u.Phone,
		CompanyName: u.CompanyName,
		JobTitle:    u.JobTitle,
		Active:      u.Active,
	}
}
