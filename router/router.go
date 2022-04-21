package router

import (
	"api_ticket/config/db"

	"api_ticket/controller/ticketController"
	"api_ticket/repository/ticketRepository"
	"api_ticket/usecase/jwtUsecase"
	"api_ticket/usecase/ticketUsecase"

	"api_ticket/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	postgresConn := db.DB
	dbLog := db.NewLogDbCustom(db.DBlog)

	Repository := ticketRepository.GetTicketRepository(postgresConn)
	jwtUsecase := jwtUsecase.GetJwtUsecase(Repository)
	Usecase := ticketUsecase.GetTicketUsecase(Repository, jwtUsecase, dbLog)
	Controller := ticketController.GetTickerController(Usecase, dbLog)
	defaultCors := middleware.CORSMiddleware()

	router := gin.Default()
	router.Use(defaultCors)

	router.POST("/user", Controller.CreateNewUser)
	router.POST("/login", Controller.UserLogin)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.JWTAuth(jwtUsecase))
	{
		protectedRoutes.GET("/tickets", Controller.GetAllTickets)
		protectedRoutes.GET("/ticket/:id", Controller.GetTicketById)
		protectedRoutes.POST("/ticket", Controller.CreateNewTicket)
		protectedRoutes.PUT("/ticket/:id", Controller.EditTicket)
		protectedRoutes.DELETE("/ticket/:id", Controller.DeleteTicket)
	}

	router.Run(":8001")

	return router
}
