package model

import (
	"time"
)

type User struct {
	Password                string    `json:"password"`
	Name                    string    `json:"name" validate:"required, len gt 2"`
	Id                      int       `json:"id"`
	Email                   string    `jason:"email" validate:"required,email"`
	Phone                   string    `json:"phone" validate:"required,phone len eq 11"`
	Company_name            string    `json:"company_name"`
	Job_title               string    `json:"job_title"`
	Active                  bool      `json:"active"`
	Subscribe_news          string    `json:"subscribe_news"`
	Subscribe_notifications bool      `json:"subscribe_notifications"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	WebSites                []Website
	Plan                    []Plan `gorm:"many2many:UserPlan"`
}
