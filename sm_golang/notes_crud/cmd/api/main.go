package main

import (
	"fmt"
	"log"

	"github.com/itz-prashant/notes-api/internal/config"
	"github.com/itz-prashant/notes-api/internal/db"
	"github.com/itz-prashant/notes-api/internal/server"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Config error", err)
	}

	client, database , err := db.Connect(cfg)

	if err != nil {
		log.Fatal("Db error", err)
	}

	defer func() {
		if err := db.DisConnect(client); err != nil {
			log.Fatal("Mongo disconnected error", err)
		}
	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatal("Server failed")
	}
}
