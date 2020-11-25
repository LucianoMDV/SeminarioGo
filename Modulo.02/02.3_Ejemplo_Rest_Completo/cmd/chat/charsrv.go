// Seminario GoLang - 02.3 Ejemplo Rest Completo (parte 1)
// Seminario GoLang - 02.4 Ejemplo Rest Completo (parte 2)
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LucianoMDV/SeminarioGo/Modulo.02/02.3_Ejemplo_Rest_Completo/internal/config"
	"github.com/LucianoMDV/SeminarioGo/Modulo.02/02.3_Ejemplo_Rest_Completo/internal/service/chat"
)

func main() {

	cfg := readConfig()
	//go run cmd/chat/charsrv.go -config ./config/config.yaml
	// fmt.Println(cfg.DB.Driver)
	// fmt.Println(cfg.Version)

	service, _ := chat.New(cfg)
	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
