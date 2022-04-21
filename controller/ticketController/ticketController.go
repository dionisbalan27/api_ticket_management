package ticketController

import (
	"api_ticket/config/db"
	"api_ticket/controller"
	"api_ticket/models"
	"api_ticket/usecase"
	"api_ticket/utils"

	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ticketController struct {
	usecase usecase.TicketUsecaseInterface
	log     *db.LogCustom
}

func GetTickerController(ticketUsecase usecase.TicketUsecaseInterface, dbLog *db.LogCustom) controller.TicketControllerInterface {
	return &ticketController{
		usecase: ticketUsecase,
		log:     dbLog,
	}
}

func (res *ticketController) CreateNewUser(c *gin.Context) {
	request := models.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		res.log.Error(errors.New("binding json error"), "CreateNewUser controller : binding json error")
	}
	response := res.usecase.CreateNewUser(request)
	if response.Status != "ok" {
		res.log.Error(errors.New(response.Status), "CreateNewUser controller : response from usecase error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (res *ticketController) UserLogin(c *gin.Context) {
	request := models.Login{}
	if err := c.ShouldBindJSON(&request); err != nil {
		res.log.Error(errors.New("binding json error"), "UserLogin controller : binding json error")
	}
	response := res.usecase.UserLogin(request)
	if response.Status != "ok" {
		res.log.Error(errors.New(response.Status), "UserLogin controller : response from usecase error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (res *ticketController) GetAllTickets(c *gin.Context) {

	if param := c.Query("s"); param != "" {
		response := res.usecase.SearchTicket(param)
		if response.Status != "ok" {
			res.log.Error(errors.New(response.Status), "controller : search ticket by params error")
			c.JSON(response.StatusCode, response)
			return
		}
		res.log.Success("controller : search ticket by params success")
		c.JSON(http.StatusOK, response)
		return
	}
	response := res.usecase.GetAllTickets()
	// fmt.Printf("%+v", response)
	if response.Status != "ok" {
		res.log.Error(errors.New(response.Status), "controller : get all tickets error")
		c.JSON(response.StatusCode, response)
		return
	}
	res.log.Success("controller : get all tickets success")
	c.JSON(http.StatusOK, response)
}

func (res *ticketController) GetTicketById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.GetTicketById(id)
	if response.StatusCode == http.StatusNotFound {
		res.log.Error(errors.New(response.Status), "GetTicketById controller : not found id ")
		c.JSON(http.StatusOK, response)
		return
	}

	if response.Status != "ok" {
		res.log.Error(errors.New(response.Status), "controller : get ticket by id error ")
		c.JSON(response.StatusCode, response)
		return
	}
	res.log.Success("controller : get ticket by id success")
	c.JSON(http.StatusOK, response)
}

func (res *ticketController) CreateNewTicket(c *gin.Context) {
	var requestbody models.Ticket
	var req2 map[string]interface{}
	dataku, _ := ioutil.ReadAll(c.Request.Body)
	finalData := string(dataku)

	json.Unmarshal([]byte(string(finalData)), &req2)
	if errkua := utils.ValidatorJsonName(req2, requestbody); errkua.Status != "ok" {
		res.log.Error(errors.New(errkua.Status), "CreateNewTicket controller : wrong json key")
		c.JSON(errkua.StatusCode, errkua)
		return
	}

	json.Unmarshal([]byte(string(finalData)), &requestbody)
	if respon := utils.ValidateVal(requestbody); respon.Status != "ok" {
		res.log.Error(errors.New(respon.Status), "CreateNewTicket controller : wrong on json value")
		c.JSON(respon.StatusCode, respon)
		return
	}

	response := res.usecase.CreateNewTicket(requestbody)
	if response.Status != "ok" {
		res.log.Error(errors.New(response.Status), "controller : create new ticket error")
		c.JSON(response.StatusCode, response)
		return
	}
	res.log.Success("controller : create new ticket success")
	c.JSON(http.StatusOK, response)

}

func (res *ticketController) EditTicket(c *gin.Context) {
	var requestbody models.Ticket
	var req2 map[string]interface{}
	dataku, _ := ioutil.ReadAll(c.Request.Body)
	finalData := string(dataku)

	json.Unmarshal([]byte(string(finalData)), &req2)
	if errkua := utils.ValidatorJsonName(req2, requestbody); errkua.Status != "ok" {
		res.log.Error(errors.New(errkua.Status), "EditTicket controller :  wrong json key")
		c.JSON(errkua.StatusCode, errkua)
		return
	}

	json.Unmarshal([]byte(string(finalData)), &requestbody)
	if respon := utils.ValidateVal(requestbody); respon.Status != "ok" {
		res.log.Error(errors.New(respon.Status), "EditTicket controller : wrong on json value")
		c.JSON(respon.StatusCode, respon)
		return
	}

	id := c.Param("id")
	response := res.usecase.EditTicket(requestbody, id)
	if response.Status != "ok" {
		res.log.Error(errors.New(response.Status), "controller : edit ticket error")
		c.JSON(response.StatusCode, response)
		return
	}
	res.log.Success("controller : edit ticket success")
	c.JSON(http.StatusOK, response)
}

func (res *ticketController) DeleteTicket(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteTicket(id)
	if response.Status != "ok" {
		res.log.Error(errors.New(response.Status), "controller : delete ticket error")
		c.JSON(response.StatusCode, response)
		return
	}
	res.log.Success("controller : delete ticket success")
	c.JSON(http.StatusOK, response)
}
