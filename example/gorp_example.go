package main

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	dbMap := initDb()
	defer dbMap.Db.Close()

	// err := dbMap.TruncateTables();
	// checkErr(err, "TruncateTables failed")

	u1 := newUser("sasuke")
	u2 := newUser("mimi")

	err := dbMap.Insert(&u1, &u2)
	checkErr(err, "Insert failed")

	sq.Select("count(*)").From("user").ToSql()
	count, err := dbMap.SelectInt("select count(*) from user")
	checkErr(err, "select count(*) failed")
	log.Println("Rows after inserting:", count)

	u2.Name = "mimisuke"
	count, err = dbMap.Update(&u2)
	checkErr(err, "Update failed")
	log.Println("Rows updated:", count)

	err = dbMap.SelectOne(&u2, "select * from user where id = ?", u2.Id)
	checkErr(err, "SelectOne failed")
	log.Println("u2 row:", u2)

	var users []User
	_, err = dbMap.Select(&users, "select * from user order by id")
	checkErr(err, "Select failed")
	log.Println("All row:")
	for i, u := range users {
		log.Printf("    %d: %v\n", i, u)
	}
}

type User struct {
	Id   uint
	Name string
}

func newUser(name string) User {
	return User{
		Name: name,
	}
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "sarasota-usr:sarasota-pwd@tcp(dev-dotmoney-dbm01.amb-stg-295.a4c.jp:3306)/sarasota")
	checkErr(err, "sql.Open failed")

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	dbMap.AddTableWithName(User{}, "user").SetKeys(true, "Id")

	// err = dbmap.CreateTablesIfNotExists()
	// checkErr(err, "Create tables failed")

	return dbMap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
