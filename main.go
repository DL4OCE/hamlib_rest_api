package main

import (
	"embed"
	"fmt"
	"hamlib_rest_api/apihandlers"
	"io/fs"
	"log"
	"net/http"
)

//go:embed gui/*
var webFiles embed.FS

func main() {
	mux := http.NewServeMux()
	apihandlers.RegisterRoutesRigctld(mux)
	apihandlers.RegisterRoutesRotctld(mux)

	publicFiles, err := fs.Sub(webFiles, "gui")
	if err != nil {
		log.Fatalf("Failed to create subtree for web files: %v", err)
	}
	fileServer := http.FileServer(http.FS(publicFiles))
	mux.Handle("/gui/", http.StripPrefix("/gui/", fileServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/gui/", http.StatusFound)
			return
		}
		http.NotFound(w, r)
	})

	fmt.Println("Hamlib REST API running on http://localhost:8080 ...")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Severe server error: %v\n", err)
	}
}
