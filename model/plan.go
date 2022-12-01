package model

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	ID     int    `json:"id"`
	Title  string `json:"title" gorm:"unique"`
	Custom bool   `json:"custom"`
}

type UserPlan struct {
	ID     int       `json:"id" gorm:"autoIncrement"`
	UserID int       `gorm:"primaryKey"`
	PlanID int       `gorm:"primaryKey"`
	ExTime time.Time `json:"ex_time"`
	gorm.Model
}
