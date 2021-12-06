package main

import (
	"log"
	"net/http"

	"github.service.anz/shenw/middleware/pkg/server"
)

func main() {
	srv := server.NewServer(3)
	srv.Routes()

	log.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", srv.Router))
}
