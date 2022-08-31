package database

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Database struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

func New() *Database {
	return &Database{}
}

func (database *Database) Setup(logger *logrus.Logger) {
	if os.Getenv("DATABASE_NAME") == "" {
		database.Name = "postgres"
		database.Host = "localhost"
		database.Port = "5432"
		database.Username = "postgres"
		database.Password = "root"
	} else {
		database.Name = os.Getenv("DATABASE_NAME")
		database.Host = os.Getenv("DATABASE_HOST")
		database.Port = os.Getenv("DATABASE_PORT")
		database.Username = os.Getenv("DATABASE_USERNAME")
		database.Password = os.Getenv("DATABASE_PASSWORD")
	}
}
