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
	
	PORT := os.Getenv("PORT")
	fmt.Println("Starting server on the port ", PORT)

	log.Fatal(http.ListenAndServe(":" + PORT, r))
}
