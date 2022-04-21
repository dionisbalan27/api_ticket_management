package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	pglogrus "gopkg.in/gemnasium/logrus-postgresql-hook.v1"
	"gorm.io/gorm"
)

type LogCustom struct {
	Logrus *logrus.Logger
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

var instance *LogCustom
var once sync.Once

func NewLogDbCustom(db *gorm.DB) *LogCustom {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	sqlDb, err := db.DB()

	hook := pglogrus.NewAsyncHook(sqlDb, map[string]interface{}{})
	hook.InsertFunc = func(sqlDb *sql.Tx, entry *logrus.Entry) error {
		level := entry.Level.String()
		if level == "info" {
			level = "success"
		}

		_, err = sqlDb.Exec("INSERT INTO log_tables (level,  location, message)VALUES ($1,$2,$3);", level, entry.Data["error_cause"], entry.Message)
		return err
	}
	log.AddHook(hook)

	once.Do(func() {
		instance = &LogCustom{
			Logrus: log,
		}
	})

	return instance
}
func (l *LogCustom) Error(err error, description string) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"error_cause": stFormat,
	}).Error(description)
}

func (l *LogCustom) Success(description string) {
	l.Logrus.Info(description)
}
