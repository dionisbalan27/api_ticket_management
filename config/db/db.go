package db

import (
	"fmt"
	"log"
	"time"

	"api_ticket/config/env"
	"api_ticket/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBlog *gorm.DB

var config = env.Config

func init() {

	//Config DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.PostgresConfig.Host,
		config.PostgresConfig.User,
		config.PostgresConfig.Password,
		config.PostgresConfig.Name,
		config.PostgresConfig.Port)

	//Open DB Connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("config/db: cannot connect DB, ", err.Error())
	}
	DB = db

	//Config DB Log
	dsn2 := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.PostgresConfig.Host,
		config.PostgresConfig.User,
		config.PostgresConfig.Password,
		config.DbNameLog,
		config.PostgresConfig.Port)

	//Open DB Connection Log
	dblog, err2 := gorm.Open(postgres.Open(dsn2), &gorm.Config{})
	if err2 != nil {
		log.Println("config/db: cannot connect DB, ", err.Error())
	}
	DBlog = dblog

	// Get generic database object sql.DB to use its functions
	postgresDB, _ := db.DB()

	//Database Connection Pool
	postgresDB.SetMaxIdleConns(10)
	postgresDB.SetMaxOpenConns(100)
	postgresDB.SetConnMaxLifetime(time.Hour)

	//Auto Migrate Struct
	AutoMigrate(db)
	AutoMigrateLog(dblog)

	err = postgresDB.Ping()

	//
	if err != nil {
		log.Fatal(err, "config/db: can't ping the db", nil)
	} else {
		go doEvery(10*time.Minute, pingDb, DB)
		return
	}
}

func AutoMigrate(db *gorm.DB) {
	err2 := db.Debug().AutoMigrate(&models.TicketDB{}, &models.User{}, &models.LogTable{})
	if err2 != nil {
		log.Fatal(err2, "config/db: can't migrate db", nil)
		return
	}
}
func AutoMigrateLog(db *gorm.DB) {
	err2 := db.Debug().AutoMigrate(&models.LogTable{})
	if err2 != nil {
		log.Fatal(err2, "config/db: can't migrate db", nil)
		return
	}
}

func doEvery(d time.Duration, f func(*gorm.DB), x *gorm.DB) {
	for range time.Tick(d) {
		f(x)
	}
}

func pingDb(db *gorm.DB) {
	s, err := db.DB()
	if err != nil {
		log.Println("config/db: can't ping the db, massage", err.Error())
		return
	}
	err = s.Ping()
	if err != nil {
		log.Println("config/db: can't ping the db, massage", err.Error())
	}
}
