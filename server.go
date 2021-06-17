package main

import (
	"fmt"
	"net/http"

	BDD "./BDD"
	handlers "./handlers"
)

func main() {

	colorGreen := "\033[32m"
	colorBlue := "\033[34m"
	colorYellow := "\033[33m"

	fmt.Println(string(colorBlue), "[SERVER_INFO] : Starting local Server...")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/connexion", handlers.ConnexionPage)
	http.HandleFunc("/login", handlers.GetLogin)
	http.HandleFunc("/register", handlers.GetRegister)
	http.HandleFunc("/likedPosts", handlers.Liked_Posts)
	http.HandleFunc("/oneCategory", handlers.One_Category)
	/* 	http.HandleFunc("/oneCategory/post", handlers.GetTopic) */
	http.HandleFunc("/postsActivity", handlers.Posts_Activity)
	http.HandleFunc("/topic", handlers.TopicPage)
	http.HandleFunc("/inscription", handlers.InscriptionPage)
	http.HandleFunc("/all_categories", handlers.RetrieveCat)
	http.HandleFunc("/BDD", BDD.Afficher)
	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println(string(colorGreen), "[SERVER_READY] : on http://localhost:8080 ✅ ")
	fmt.Println(string(colorYellow), "[SERVER_INFO] : To stop the program : Ctrl + c")
	http.ListenAndServe(":8080", nil)
}
