package ticketController

import (
	"api_ticket/config/db"
	"api_ticket/models"
	"api_ticket/repository/ticketRepository"
	"api_ticket/usecase/ticketUsecase"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {

	request := models.User{
		Name:            "Wangdu",
		Personal_number: "9999",
		Email:           "wangdu@gmail",
		Password:        "password"}

	Repository := ticketRepository.GetTicketRepository(db.DB)
	Usecase := ticketUsecase.GetTicketUsecase(Repository, nil, nil)

	response := Usecase.CreateNewUser(request)
	assert.Equal(t, "ok", response.Status, "CreateNewUser controller : response from usecase error")
}

func TestUserLogin(t *testing.T) {

	request := models.Login{
		Personal_number: "9999",
		Password:        "password"}

	Repository := ticketRepository.GetTicketRepository(db.DB)
	Usecase := ticketUsecase.GetTicketUsecase(Repository, nil, nil)

	response := Usecase.UserLogin(request)
	assert.Equal(t, "ok", response.Status, "error response from user login ")
}

func TestGetAllTickets(t *testing.T) {

	Repository := ticketRepository.GetTicketRepository(db.DB)
	Usecase := ticketUsecase.GetTicketUsecase(Repository, nil, nil)

	response := Usecase.GetAllTickets()
	assert.Equal(t, "ok", response.Status, "error response get all tickets")
}

func TestGetTicketById(t *testing.T) {

	id := "bd59cdb6-ac9b-4369-80b3-8335dc7fdb41"

	Repository := ticketRepository.GetTicketRepository(db.DB)
	Usecase := ticketUsecase.GetTicketUsecase(Repository, nil, nil)

	response := Usecase.GetTicketById(id)
	assert.NotEqual(t, http.StatusNotFound, response.Status, "not found get ticket by id")
	assert.Equal(t, "ok", response.Status, "error response get ticket by id")
}

func TestSearchTicketByParams(t *testing.T) {
	params := "Super"

	Repository := ticketRepository.GetTicketRepository(db.DB)
	Usecase := ticketUsecase.GetTicketUsecase(Repository, nil, nil)

	response := Usecase.SearchTicket(params)
	assert.Equal(t, "ok", response.Status, "error get ticket by params")
}
