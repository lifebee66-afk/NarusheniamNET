package main

import (
	"net/http"

	"github.com/lifebee66-afk/NarusheniamNET/handlers"
)

func main() {
	h := handlers.NewHandler()
	mux := http.NewServeMux()
	mux.Handle("POST")
}
