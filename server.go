package main

import (
	"fmt"
	"net/http"

	handlers "./handlers"
	connexion "./src/connection"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/connexion", connexion.ConnexionPage)
	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
