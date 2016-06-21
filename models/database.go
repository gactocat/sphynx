package models

import (
	"database/sql"

	"github.com/gactocat/snowshoe/db"
)

var database DatabaseInterface = newDatabase()

func conn() ConnectionInterface {
	return database.Connection()
}

func newDatabase() DatabaseInterface {
	sqlDb, _ := sql.Open("mysql", "sarasota-usr:sarasota-pwd@tcp(dev-dotmoney-dbm01.amb-stg-295.a4c.jp:3306)/sarasota")
	return NewDatabase(sqlDb, Config{})
}

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
