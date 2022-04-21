package repository

import "api_ticket/models"

type TicketRepositoryInterface interface {
	CreateNewUser(models.User) (*models.User, error)
	CheckLogin(models.Login) (*models.CheckLogin, error)
	GetAllTickets() ([]models.Ticket, error)
	GetSearchTicket(string) ([]models.Ticket, error)
	GetTicketById(string) (*models.Ticket, error)
	CreateNewTicket(models.TicketDB) (*models.TicketDB, error)
	EditTicket(models.TicketDB, string) (*models.TicketDB, error)
	DeleteTicket(string) error
}
