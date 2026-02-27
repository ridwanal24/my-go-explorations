package main

import (
	"authentication-feature/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes := routes.SetupRoutes()

	fmt.Println("Serve to port 8000")
	err := http.ListenAndServe(":8000", routes)

	if err != nil {
		log.Println(err.Error())
	}
}
