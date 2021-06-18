package main

import (
	"net/http"

	BDD "./BDD"
	handlers "./handlers"
)

func main() {
<<<<<<< HEAD
=======

	handlers.Color(2, "[SERVER_INFO] : Starting local Server...")

>>>>>>> 5f25b74fa4d4959cb5742afbcc115b32a25c5ca1
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/connexion", handlers.ConnexionPage)
	http.HandleFunc("/login", handlers.GetLogin)
	http.HandleFunc("/register", handlers.GetRegister)
	http.HandleFunc("/likedPosts", handlers.Liked_Posts)
	//http.HandleFunc("/oneCategory", handlers.One_Category)
	/* 	http.HandleFunc("/oneCategory/post", handlers.GetTopic) */
	http.HandleFunc("/postsActivity", handlers.Posts_Activity)
	http.HandleFunc("/topic", handlers.TopicPage)
	http.HandleFunc("/inscription", handlers.InscriptionPage)
	http.HandleFunc("/all_categories", handlers.RetrieveCat)
	http.HandleFunc("/BDD", BDD.Afficher)
	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	handlers.Color(1, "[SERVER_READY] : on http://localhost:8080 ✅ ")
	handlers.Color(3, "[SERVER_INFO] : To stop the program : Ctrl + c")

	http.ListenAndServe(":8080", nil)
}
