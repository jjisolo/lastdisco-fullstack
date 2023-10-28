package main

import (
	"log"
	"github.com/jjisolo/lastdisco-backend/api"
	"github.com/jjisolo/lastdisco-backend/storage"
)

func main() {
	storage, err := storage.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)	
	}
	storage.Initialize()

	server := api.NewAPIServer(":3000", storage) 
	server.Run()
}