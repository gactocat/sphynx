package models

import (
	"database/sql"

	"github.com/gactocat/snowshoe/db"
)

type DatabaseInterface interface {
	db.DatabaseInterface
}

type ConnectionInterface interface {
	db.ConnectionInterface
}

type Config struct {
	DefaultTemplatePath string
}

func NewDatabase(sqlDB *sql.DB, config Config) DatabaseInterface {
	database := db.NewDatabase(sqlDB, db.Config{
		DefaultTemplatePath: config.DefaultTemplatePath,
	})
	Setup(database)

	return database
}

func Setup(database *db.DB) {
	database.TableMap().AddTableWithName(User{}, "user").SetKeys(true, "Id")
}
