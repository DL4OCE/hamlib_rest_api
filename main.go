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
	apihandlers.RegisterRoutes(mux)

	publicFiles, err := fs.Sub(webFiles, "gui")
	if err != nil {
		log.Fatalf("Failed to create subtree for web files: %v", err)
	}
	fileServer := http.FileServer(http.FS(publicFiles))
	mux.Handle("/gui/", http.StripPrefix("/gui/", fileServer))

	// mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.URL.Path == "/" {
	// 		http.Redirect(w, r, "/gui/", http.StatusFound)
	// 		return
	// 	}
	// 	http.NotFound(w, r)
	// })

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Da "/" im alten Stil alles matcht, müssen wir hier
		// strikt prüfen, ob wirklich NUR die Root aufgerufen wurde
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/gui/", http.StatusFound)
			return
		}
		// Alles andere, was nicht von RegisterRoutes oder /gui/ abgefangen wurde,
		// läuft hier rein und bekommt ein sauberes 404
		http.NotFound(w, r)
	})

	// log.Println("Starting server on :8080...")
	// log.Fatal(http.ListenAndServe(":8080", mux))

	fmt.Println("Hamlib REST API running on http://localhost:8080 ...")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Severe server error: %v\n", err)
	}

	// // GUI
	// publicFiles, _ := fs.Sub(webFiles, "web")
	// fileServer := http.FileServer(http.FS(publicFiles))
	// err := http.ListenAndServe(":8081", mux)
	// if err != nil {
	// 	fmt.Printf("Severe server error: %v\n", err)
	// }
	// fmt.Println("Hamlib GUI running on http://localhost:8081 ...")
}
