package ticketRepository

import (
	"api_ticket/config/db"
	"api_ticket/models"
	"api_ticket/utils"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {

	newUser := models.User{
		ID:              uuid.New().String(),
		Name:            "Wangdu",
		Personal_number: "9999",
		Email:           "wangdu@gmail",
		Password:        "password",
	}

	errKueriDB := db.DB.Create(&newUser).Error

	assert.Equal(t, nil, errKueriDB, "Error create new user to DB")

}

func TestCheckLogin(t *testing.T) {
	data := models.CheckLogin{}
	user := models.Login{
		Personal_number: "1111",
		Password:        "password",
	}

	errKueriDB := db.DB.Model(&models.User{}).Where("users.personal_number = ?", user.Personal_number).Select("users.id, users.name, users.password").Scan(&data).Error

	passwordBase64 := utils.Base64(user.Password, "enc")
	_, errCheckPass := utils.CheckPasswordHash(passwordBase64, data.Password)

	assert.Equal(t, nil, errKueriDB, "Error get User id, name, and password from DB")
	assert.Equal(t, nil, errCheckPass, "Error check password hash")
	assert.Equal(t, "adi", data.Name, "Wrong User name")
	assert.Equal(t, "password", data.Password, "Wrong user password")
	assert.NotNil(t, data.ID, "User ID not found")
	assert.NotNil(t, data.Name, "User name not found")
	assert.NotNil(t, data.ID, "User ID not found")
}

func TestGetTicketById(t *testing.T) {
	id := "bd59cdb6-ac9b-4369-80b3-8335dc7fdb41"
	tickets := models.Ticket{}

	errKueriDB := db.DB.Model(&models.TicketDB{}).Where("ticket_dbs.ID = ?", id).Select("ticket_dbs.title_movie, ticket_dbs.studio, ticket_dbs.name, ticket_dbs.seat").Scan(&tickets).Error
	assert.Equal(t, nil, errKueriDB, "Error get ticket by id")
}
