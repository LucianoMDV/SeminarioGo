// Seminario GoLang - 02.3 Ejemplo Rest Completo (parte 1)
// Seminario GoLang - 02.4 Ejemplo Rest Completo (parte 2)
// Seminario GoLang - 02.5 Ejemplo Rest Completo (parte 3)
// Seminario GoLang - 02.6 Ejemplo Rest Completo (parte 4)
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LucianoMDV/SeminarioGo/Modulo.02/02.3_Ejemplo_Rest_Completo/internal/config"
	"github.com/LucianoMDV/SeminarioGo/Modulo.02/02.3_Ejemplo_Rest_Completo/internal/database"
	chat "github.com/LucianoMDV/SeminarioGo/Modulo.02/02.3_Ejemplo_Rest_Completo/internal/service/producto"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := readConfig()
	// go run cmd/chat/chatsrv.go -config ./config/config.yaml
	// go run cmd/chat/chatsrv.go

	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	/**
	* linea de comando "rm test/data.db" si hay problemas eliminar data.db
	* y con esta funcion se genera de nuevo la BD
	**/
	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := chat.New(db, cfg)
	httpService := chat.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()

	// for _, m := range service.FindAll() {
	// 	fmt.Println(m)
	// }
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config/config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS producto (
		id integer primary key autoincrement,
		nombre varchar,
		precio integer
		);`

	//execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	//or, you can use MustExec, which panics on error
	// insertMessage := `INSERT INTO messages (text) VALUES (?)`
	// s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	// db.MustExec(insertMessage, s)
	return nil
}
