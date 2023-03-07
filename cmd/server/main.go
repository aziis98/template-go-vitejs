package main

import (
	"log"

	"aziis98.com/template-go-vitejs/backend/database"
	"aziis98.com/template-go-vitejs/backend/server"
)

func main() {
	srv := server.New(server.Dependencies{
		DB: &database.Mem{},
	})

	log.Fatal(srv.Listen(":4000"))
}
