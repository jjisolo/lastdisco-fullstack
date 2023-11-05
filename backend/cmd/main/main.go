package main

import (
	"github.com/jjisolo/lastdisco-backend/config"
	"github.com/jjisolo/lastdisco-backend/internal/storage/pgstorage"
	"github.com/jjisolo/lastdisco-backend/internal/transport/rest"
	"log"
)

func main() {
	storage, err := pgstorage.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = storage.Initialize()
	if err != nil {
		log.Fatal(err)
		return
	}

	if config.TESTING {
		server := rest.NewAPIServer(":3000", storage)
		server.Run()
	} else {
		server := rest.NewAPIServer(":80", storage)
		server.Run()
	}
}
