// Seminario GoLang - 02.2 SQL
// https://jmoiron.github.io/sqlx/
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// User is a obcject
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	createSchema(db)
}

func createSchema(db *sqlx.DB) {
	schema := `CREATE TABLE user (
		id integer primary key autoincrement,
		name varchar(56));`
	_, err := db.Exec(schema)

	if err != nil {
		panic(err.Error())
	}

	db.MustExec("INSERT INTO user (name) VALUES (?)", "jane doe")
	rows, err := db.Query("SELECT id, name FROM user")

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		fmt.Println(u)
	}
}
