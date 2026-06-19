package main

import (
	"fmt"
	"hamlib_rest_api/apihandlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	apihandlers.RegisterRoutes(mux)

	fmt.Println("Hamlib Go-API läuft auf http://localhost:8080 ...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Schwerwiegender Server-Fehler: %v\n", err)
	}
}
