package main

import (
	"github.com/LucianoMDV/SeminarioGo/internal/config"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
}
