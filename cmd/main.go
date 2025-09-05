package main

import (
	"log"
	"net/http"

	"github.com/lifebee66-afk/NarusheniamNET/handlers"
	postgresql "github.com/lifebee66-afk/NarusheniamNET/pkg/postgreSQL"
)

func main() {
	db, err := postgresql.SQLinit()
	if err != nil {
		log.Fatal("failed connect data base", err)
	}
	h := handlers.NewHandler(db)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", h.RegisterUserHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
