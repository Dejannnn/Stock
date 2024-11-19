package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Dejannnn/Stock.git/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	port := os.Getenv("PORT")
	if err != nil {
		log.Fatalf("Port not found %v", port)
	}
	r := router.Router()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Server listen on port %v", port)
	log.Fatal(server.ListenAndServe())
}
