package controller

import "github.com/gin-gonic/gin"

type TicketControllerInterface interface {
	CreateNewUser(*gin.Context)
	UserLogin(*gin.Context)
	GetAllTickets(*gin.Context)
	GetTicketById(*gin.Context)
	CreateNewTicket(*gin.Context)
	EditTicket(*gin.Context)
	DeleteTicket(*gin.Context)
}
