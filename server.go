package main

import (
	"fmt"
	"net/http"
	handlers "./handlers"
)

func main(){

	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/", handlers.Home)

	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}