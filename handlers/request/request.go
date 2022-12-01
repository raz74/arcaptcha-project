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

type CreateWebsiteRequest struct {
	UserId      int    `json:"user_id"`
	SiteKey     string `json:"site_key"`
	SecretKey   string `json:"secret_key"`
	Label       string `json:"label"`
	ChType      string `json:"ch_type"`
	Level       int    `json:"level"`
	FingerPrint bool   `json:"fingerprint"`
	Brand       bool   `json:"brand"`
}

type UpdateWebsiteRequest struct {
	UserId    int    `json:"user_id"`
	SiteKey   string `json:"site_key"`
	SecretKey string `json:"secret_key"`
	Label     string `json:"label"`
	Alert     bool   `json:"alert"`
	Subdomain bool   `json:"subdomain"`
}

type PlanRequest struct {
	UserId int       `json:"user_id"`
	PlanId int       `json:"plan_id"`
	ExTime time.Time `json:"ex_Time"`
}

