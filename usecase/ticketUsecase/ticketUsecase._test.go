package ticketUsecase

import (
	"api_ticket/config/db"
	"api_ticket/models"
	"api_ticket/repository/ticketRepository"
	"api_ticket/utils"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

func TestCreateNewUser(t *testing.T) {

	newUser := models.User{
		Name:            "Wangdu",
		Personal_number: "9999",
		Email:           "wangdu@gmail",
		Password:        "password",
	}
	passwordBase64 := utils.Base64(newUser.Password, "enc")
	paswordHash, _ := utils.HashPassword(passwordBase64)
	newUser.Password = paswordHash

	_, err1 := ticketRepository.GetTicketRepository(db.DB).CreateNewUser(newUser)
	assert.Equal(t, nil, err1, "Error create new user to DB")

}

func TestUserLogin(t *testing.T) {

	newUser := models.Login{
		Personal_number: "9999",
		Password:        "password",
	}

	_, err1 := ticketRepository.GetTicketRepository(db.DB).CheckLogin(newUser)
	assert.Equal(t, nil, err1, "login data not found in DB")

}

func TestGetAllTickets(t *testing.T) {

	_, err1 := ticketRepository.GetTicketRepository(db.DB).GetAllTickets()
	assert.Equal(t, errors.Is(err1, gorm.ErrRecordNotFound), err1, "get all tickets not found")
	assert.Equal(t, nil, err1, "Internal server error when get all tickets")
}
func TestGetTicketById(t *testing.T) {
	id := "bd59cdb6-ac9b-4369-80b3-8335dc7fdb41"

	_, err1 := ticketRepository.GetTicketRepository(db.DB).GetTicketById(id)
	assert.Equal(t, errors.Is(err1, gorm.ErrRecordNotFound), err1, "get ticket by id not found")
	assert.Equal(t, nil, err1, "Internal server error when get ticket by id")
}

func TestSearchTicketByParams(t *testing.T) {
	params := "Super"

	_, err1 := ticketRepository.GetTicketRepository(db.DB).GetSearchTicket(params)
	assert.Equal(t, errors.Is(err1, gorm.ErrRecordNotFound), err1, "search ticket by params not found")
	assert.Equal(t, nil, err1, "Internal server error when search ticket by params")
}

func TestCreateNewTicket(t *testing.T) {

	newTicket := models.TicketDB{
		Title_movie: "Naruto the movie",
		Studio:      "2",
		Name:        "Rani",
		Seat:        "28C",
	}

	_, err1 := ticketRepository.GetTicketRepository(db.DB).CreateNewTicket(newTicket)
	assert.Equal(t, nil, err1, "Error create new ticket to DB")
}

func TestEditTicket(t *testing.T) {

	id := "bd59cdb6-ac9b-4369-80b3-8335dc7fdb41"
	newTicket := models.TicketDB{
		Title_movie: "Superman 2",
		Studio:      "3",
		Name:        "Greg",
		Seat:        "21F",
	}

	_, err1 := ticketRepository.GetTicketRepository(db.DB).EditTicket(newTicket, id)
	assert.Equal(t, nil, err1, "Error create new ticket to DB")
}
