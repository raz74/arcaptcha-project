package model

import "time"

type WebSite struct {
	User_Id     int       `json:"user_id"`
	Site_Key    string    `json:"site_key"`
	Secret_Key  string    `json:"secret_key"`
	Label       string    `json:"label"`
	Alert       bool      `json:"alert"`
	Subdomain   bool      `json:"subdomain"`
	Version     int       `json:"version"`
	Alert_Limit int       `json:"alert_limit"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
}

type CreateWebsiteRequest struct {
	User_Id    int    `json:"user_id"`
	Site_Key   string `json:"site_key"`
	Secret_Key string `json:"secret_key"`
	Label      string `json:"label"`
}

type CreateWebsite_v1_Request struct {
	Website_Id  int    `json:"website_id"`
	Ch_Type     string `json:"ch_type"`
	Level       int    `json:"level"`
	FingerPrint bool   `json:"fingerprint"`
	Brand       bool   `json:"brand"`
}
