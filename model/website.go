package model

import "time"

type Website struct {
	ID         int       `json:"id"`
	UserId     int       `json:"user_id"`
	SiteKey    string    `json:"site_key"`
	SecretKey  string    `json:"secret_key"`
	Label      string    `json:"label"`
	Alert      bool      `json:"alert"`
	Subdomain  bool      `json:"subdomain"`
	Version    int       `json:"version"`
	AlertLimit int       `json:"alert_limit"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	WebsiteV1  WebsiteV1
}

type WebsiteV1 struct {
	WebsiteID   int    `json:"website_id" gorm:"primaryKey"`
	ChType      string `json:"chType"`
	Level       int    `json:"level"`
	FingerPrint bool   `json:"fingerprint"`
	Brand       bool   `json:"brand"`
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
	UserId     int       `json:"user_id"`
	SiteKey    string    `json:"site_key"`
	SecretKey  string    `json:"secret_key"`
	Label      string    `json:"label"`
	Alert      bool      `json:"alert"`
	Subdomain  bool      `json:"subdomain"`
}