package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"server/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 5000...")

	PORT := os.Getenv("DB_COLLECTION_NAME")

	log.Fatal(http.ListenAndServe(PORT, r))
}
