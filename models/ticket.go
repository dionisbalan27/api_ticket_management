package models

import (
	"gorm.io/gorm"
)

type TicketDB struct {
	gorm.Model
	Title_movie string `json:"titleMovie" binding:"required"`
	Studio      string `json:"studio" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Seat        string `json:"seat" binding:"required"`
}

type Ticket struct {
	Title_movie string `json:"titleMovie" binding:"required"`
	Studio      string `json:"studio" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Seat        string `json:"seat" binding:"required"`
}
