package ticketUsecase

import (
	"api_ticket/config/db"
	"api_ticket/models"
	"api_ticket/repository"
	"api_ticket/usecase"
	"api_ticket/utils"
	"errors"

	"gorm.io/gorm"
)

type ticketUsecase struct {
	ticketRepo repository.TicketRepositoryInterface
	jwtUsecase usecase.JwtUsecase
	log        *db.LogCustom
}

func GetTicketUsecase(ticketRepository repository.TicketRepositoryInterface, jwtUsecase usecase.JwtUsecase, dbLog *db.LogCustom) usecase.TicketUsecaseInterface {
	return &ticketUsecase{
		ticketRepo: ticketRepository,
		jwtUsecase: jwtUsecase,
		log:        dbLog,
	}
}

func (t *ticketUsecase) CreateNewUser(newUser models.User) models.MsgRes {
	passwordBase64 := utils.Base64(newUser.Password, "enc")
	paswordHash, _ := utils.HashPassword(passwordBase64)
	newUser.Password = paswordHash

	userData, err := t.ticketRepo.CreateNewUser(newUser)

	if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}

	return utils.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID}, 200)
}

func (t *ticketUsecase) UserLogin(newUser models.Login) models.MsgRes {

	res, err := t.ticketRepo.CheckLogin(newUser)
	if err != nil {
		return utils.ResponseError("Data not found", err, 404)
	}

	token, err := t.jwtUsecase.GenerateToken(res.ID)
	if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}

	return utils.ResponseSuccess("ok", nil, map[string]interface{}{
		"token": token, "name": res.Name}, 200)

}

func (t *ticketUsecase) GetAllTickets() models.MsgRes {
	ticketlist, err := t.ticketRepo.GetAllTickets()

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.Error(errors.New("Data not found"), "GetAllTickets usecase :  Data not found")
		return utils.ResponseError("Data not found", err, 404)
	} else if err != nil {
		t.log.Error(errors.New("Internal server error"), "GetAllTickets usecase : Internal server error")
		return utils.ResponseError("Internal server error", err, 500)
	}
	return utils.ResponseSuccess("ok", nil, ticketlist, 200)
}

func (t *ticketUsecase) SearchTicket(param string) models.MsgRes {
	ticketlist, err := t.ticketRepo.GetSearchTicket(param)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.Error(errors.New("Data not found"), "SearchTicket usecase :  Data not found")
		return utils.ResponseError("Data not found", err, 404)
	} else if err != nil {
		t.log.Error(errors.New("Internal server error"), "SearchTicket usecase :  Internal server error")
		return utils.ResponseError("Internal server error", err, 500)
	}
	return utils.ResponseSuccess("ok", nil, ticketlist, 200)
}

func (t *ticketUsecase) GetTicketById(id string) models.MsgRes {
	ticket, err := t.ticketRepo.GetTicketById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.Error(errors.New("Data not found"), "GetTicketById usecase :  Data not found")
		return utils.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		t.log.Error(errors.New("Internal server error"), "GetTicketById usecase :  Internal server error")
		return utils.ResponseError("Internal server error", err.Error(), 500)
	}

	return utils.ResponseSuccess("ok", nil, ticket, 200)
}

func (t *ticketUsecase) CreateNewTicket(newTicket models.Ticket) models.MsgRes {

	ticketInsert := models.TicketDB{
		Title_movie: newTicket.Title_movie,
		Studio:      newTicket.Studio,
		Name:        newTicket.Name,
		Seat:        newTicket.Seat,
	}

	ticketData, err := t.ticketRepo.CreateNewTicket(ticketInsert)

	if err != nil {
		if err.Error() == "Personal number already registered" {
			t.log.Error(errors.New("Conflict"), "CreateNewTicket usecase :  Conflict")
			return utils.ResponseError("Conflict", err.Error(), 409)
		} else {
			t.log.Error(errors.New("Internal server error"), "CreateNewTicket usecase :  Internal server error")
			return utils.ResponseError("Internal server error", err.Error(), 500)
		}

	}
	return utils.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": ticketData.ID}, 201)
}

func (t *ticketUsecase) EditTicket(ticketUpdate models.Ticket, id string) models.MsgRes {

	ticketInsert := models.TicketDB{
		Title_movie: ticketUpdate.Title_movie,
		Studio:      ticketUpdate.Studio,
		Name:        ticketUpdate.Name,
		Seat:        ticketUpdate.Seat,
	}

	_, err := t.ticketRepo.EditTicket(ticketInsert, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.Error(errors.New("Data not found"), "EditTicket usecase :  Data not found")
		return utils.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		if err.Error() == "Personal number already taken" {
			t.log.Error(errors.New("Confilct"), "EditTicket usecase :  Confilct")
			return utils.ResponseError("Confilct", err.Error(), 409)
		}
		t.log.Error(errors.New("Internal server error"), "EditTicket usecase :  Internal server error")
		return utils.ResponseError("Internal server error", err.Error(), 500)
	}
	return utils.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (t *ticketUsecase) DeleteTicket(id string) models.MsgRes {

	err := t.ticketRepo.DeleteTicket(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.Error(errors.New("Data not found"), "DeleteTicket usecase :  Data not found")
		return utils.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		t.log.Error(errors.New("Internal server error"), "DeleteTicket usecase :  Internal server error")
		return utils.ResponseError("Internal server error", err.Error(), 500)
	}
	return utils.ResponseSuccess("ok", nil, nil, 200)
}
