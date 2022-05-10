package router

import (
	"api_ticket/config/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"api_ticket/controller/ticketController"
	"api_ticket/repository/ticketRepository"
	"api_ticket/usecase/ticketUsecase"

	"github.com/gin-gonic/gin"
)

func InitRouterTest() *gin.Engine {

	postgresConn := db.DB

	Repository := ticketRepository.GetTicketRepository(postgresConn)
	Usecase := ticketUsecase.GetTicketUsecase(Repository, nil, nil)
	Controller := ticketController.GetTickerController(Usecase, nil)

	router := gin.Default()
	router.Use()

	router.POST("/user", Controller.CreateNewUser)
	router.POST("/login", Controller.UserLogin)
	router.GET("/tickets", Controller.GetAllTickets)
	router.GET("/ticket/:id", Controller.GetTicketById)
	router.POST("/ticket", Controller.CreateNewTicket)
	router.PUT("/ticket/:id", Controller.EditTicket)
	router.DELETE("/ticket/:id", Controller.DeleteTicket)

	router.Run(":8001")

	return router
}

func TestGetAllTickets(t *testing.T) {
	router := InitRouterTest()
	requestBody := strings.NewReader("")
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8001/tickets", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	fmt.Println("sasaassasa", recorder.Result())
	response := recorder.Result()
	assert.Equal(t, "ok", response.Status, "error response get all tickets")
	bodyResult, _ := io.ReadAll(response.Body)
	var responseBody interface{}
	json.Unmarshal(bodyResult, &responseBody)
	assert.NotEqual(t, 0, len(responseBody.([]interface{})))
}
