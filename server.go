package main

import (
	"fmt"
	"net/http"

	handlers "./handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/connexion", handlers.ConnexionPage)
	http.HandleFunc("/login", handlers.GetLogin)
	http.HandleFunc("/register", handlers.GetRegister)
	http.HandleFunc("/likedPosts", handlers.Liked_Posts)
	http.HandleFunc("/oneCategory", handlers.One_Category)
	http.HandleFunc("/postsActivity", handlers.Posts_Activity)
	http.HandleFunc("/topic", handlers.TopicPage)
	http.HandleFunc("/inscription", handlers.InscriptionPage)
	http.HandleFunc("/all_categories", handlers.All_Categories)
	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
