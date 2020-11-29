package database

import (
	"errors"
	"fmt"

	"github.com/LucianoMDV/SeminarioGo/Modulo.02/02.3_Ejemplo_Rest_Completo/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" //adding sqlite driver support
)

// NewDatabase is a ...
func NewDatabase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type {
	case "sqlite3":

		fmt.Println("DRIVER: " + conf.DB.Driver)
		fmt.Println("CONN: " + conf.DB.Conn)

		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)

		// fmt.Println(db) //aca da el problema me parece...

		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil
	default:
		return nil, errors.New("invalid database type")
	}
}
