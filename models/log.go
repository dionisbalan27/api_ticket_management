package models

import (
	"time"
)

type LogTable struct {
	ID        uint      `json:"id" gorm:"column:id"`
	Level     string    `json:"level" gorm:"column:level"`
	Location  string    `json:"location"`
	Message   string    `json:"message" gorm:"column:message"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
