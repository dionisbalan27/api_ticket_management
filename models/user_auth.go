package models

import "time"

type User struct {
	ID              string `json:"id"`
	Name            string `json:"name" binding:"required"`
	Personal_number string `json:"personalNumber"  binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	CreateAt        time.Time
}

type Login struct {
	Personal_number string `json:"personalNumber"  binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type CheckLogin struct {
	// gorm.Model
	ID       string `json:"id" gorm:"primaryKey, type:varchar(50)"`
	Name     string `json:"-" binding:"required"`
	Password string `json:"-" binding:"required"`
}
