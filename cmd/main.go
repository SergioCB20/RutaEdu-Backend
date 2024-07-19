package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/SergioCB20/RutaEdu-Backend/pkg/routes"
)

func main() {
	router := routes.SetRoutes()
	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000",router))
}