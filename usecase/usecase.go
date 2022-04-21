package usecase

import (
	"api_ticket/models"

	"github.com/golang-jwt/jwt"
)

type TicketUsecaseInterface interface {
	CreateNewUser(models.User) models.MsgRes
	UserLogin(models.Login) models.MsgRes
	GetAllTickets() models.MsgRes
	SearchTicket(string) models.MsgRes
	GetTicketById(string) models.MsgRes
	CreateNewTicket(models.Ticket) models.MsgRes
	EditTicket(models.Ticket, string) models.MsgRes
	DeleteTicket(string) models.MsgRes
}

type JwtUsecase interface {
	GenerateToken(string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	ValidateTokenAndGetUserId(string) (string, error)
}
