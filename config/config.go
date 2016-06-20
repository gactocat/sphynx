package config

import (
	"database/sql"
	"github.com/gactocat/snowshoe/db"
	"github.com/gactocat/snowshoe/models"
)

type context struct {
	Db db.DatabaseInterface
}

var sharedInstance *context = newContext()

func GetContext() *context {
	return sharedInstance
}

func newContext() *context {
	sqlDb, _ := sql.Open("mysql", "sarasota-usr:sarasota-pwd@tcp(dev-dotmoney-dbm01.amb-stg-295.a4c.jp:3306)/sarasota")
	database := models.NewDatabase(sqlDb, models.Config{})
	return &context{
		Db: database,
	}
}
