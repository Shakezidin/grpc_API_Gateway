package main

import (
	"log"

	"github.com/shakezidin/pkg/admin"
	cnfg "github.com/shakezidin/pkg/config"
	"github.com/shakezidin/pkg/server"
)

func main() {
	config, err := cnfg.LoadConfigure()
	if err != nil {
		log.Printf("Error Loading Config Files, error: %v", err)
	}

	server := server.Server()
	admin.NewAdminRoute(server.R, config)
	server.StartServer(config.APIPORT)

}
