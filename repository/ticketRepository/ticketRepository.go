package ticketRepository

import (
	"api_ticket/models"
	"api_ticket/repository"
	"api_ticket/utils"
	_ "api_ticket/utils"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ticketRepository struct {
	mysqlConnection *gorm.DB
}

func GetTicketRepository(mysqlConn *gorm.DB) repository.TicketRepositoryInterface {
	return &ticketRepository{
		mysqlConnection: mysqlConn,
	}
}

func (repo *ticketRepository) CreateNewUser(user models.User) (*models.User, error) {
	user.ID = uuid.New().String()

	if err := repo.mysqlConnection.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *ticketRepository) CheckLogin(user models.Login) (*models.CheckLogin, error) {
	data := models.CheckLogin{}

	err := repo.mysqlConnection.Model(&models.User{}).Where("users.personal_number = ?", user.Personal_number).Select("users.id, users.name, users.password").Scan(&data).Error
	if err != nil {
		return nil, err
	}

	match, err := utils.CheckPasswordHash(user.Password, data.Password)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return &data, err
	}

	return &data, nil
}

func (repo *ticketRepository) GetAllTickets() ([]models.Ticket, error) {

	tickets := []models.Ticket{}
	err := repo.mysqlConnection.Model(&models.TicketDB{}).Select("ticket_dbs.Title_movie, ticket_dbs.Studio, ticket_dbs.Name, ticket_dbs.Seat").Scan(&tickets).Error
	if err != nil {
		return nil, err
	}

	if len(tickets) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return tickets, nil
}

func (repo *ticketRepository) GetSearchTicket(param string) ([]models.Ticket, error) {
	param = "%" + param + "%"
	tickets := []models.Ticket{}
	err := repo.mysqlConnection.Model(&models.TicketDB{}).
		Where("ticket_dbs.title_movie LIKE ? OR ticket_dbs.studio LIKE ? OR ticket_dbs.name LIKE ? OR ticket_dbs.seat LIKE ? ", param, param, param, param).
		Select("ticket_dbs.Title_movie, ticket_dbs.Studio, ticket_dbs.Name, ticket_dbs.Seat").Scan(&tickets).Error
	if err != nil {
		return nil, err
	}

	if len(tickets) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return tickets, nil
}

func (repo *ticketRepository) GetTicketById(id string) (*models.Ticket, error) {
	tickets := models.Ticket{}

	err := repo.mysqlConnection.Model(&models.TicketDB{}).Where("ticket_dbs.ID = ?", id).Select("ticket_dbs.title_movie, ticket_dbs.studio, ticket_dbs.name, ticket_dbs.seat").Scan(&tickets).Error
	if err != nil {
		return nil, err
	}

	return &tickets, nil
}

func (repo *ticketRepository) CreateNewTicket(ticket models.TicketDB) (*models.TicketDB, error) {
	ticket.ID = uuid.New().String()

	if err := repo.mysqlConnection.Create(&ticket).Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (repo *ticketRepository) EditTicket(ticket models.TicketDB, id string) (*models.TicketDB, error) {

	if err := repo.mysqlConnection.Model(&ticket).Where("ID = ?", id).Updates(map[string]interface{}{
		"title_movie": ticket.Title_movie,
		"studio":      ticket.Studio,
		"name":        ticket.Name,
		"seat":        ticket.Seat,
	}).Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (repo *ticketRepository) DeleteTicket(id string) error {

	result := repo.mysqlConnection.Where("ID = ?", id).Delete(&models.TicketDB{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
