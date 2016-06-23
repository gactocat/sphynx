package models

import (
	"database/sql"

	"github.com/gactocat/snowshoe/db"
	"github.com/spf13/viper"
)

var database DatabaseInterface

func conn() ConnectionInterface {
	if database == nil {
		// TODO 排他
		database = newDatabase()
	}
	return database.Connection()
}

func newDatabase() DatabaseInterface {
	sqlDb, _ := sql.Open("mysql", viper.GetString("database.name"))
	database := db.NewDatabase(sqlDb)
	Setup(database)
	return database
}

func Setup(database *db.DB) {
	database.TableMap().AddTableWithName(User{}, "user").SetKeys(true, "Id")
}

type DatabaseInterface interface {
	db.DatabaseInterface
}

type ConnectionInterface interface {
	db.ConnectionInterface
}
