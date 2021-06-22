package main

import (
	"net/http"

	BDD "./BDD"
	handlers "./handlers"
)

func main() {

	handlers.Color(2, "[SERVER_INFO] : Starting local Server...")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/connexion", handlers.ConnexionPage)
	http.HandleFunc("/login", handlers.GetLogin)
	http.HandleFunc("/register", handlers.GetRegister)
	http.HandleFunc("/deconnexion", handlers.GetDeconnected)
	http.HandleFunc("/likedPosts", handlers.Liked_Posts)
	http.HandleFunc("/oneCategory", handlers.One_Category)
	http.HandleFunc("/postsActivity", handlers.Posts_Activity)
	http.HandleFunc("/topic", handlers.TopicPage)
	http.HandleFunc("/inscription", handlers.InscriptionPage)
	http.HandleFunc("/all_categories", handlers.RetrieveCat)
	http.HandleFunc("/BDD", BDD.Afficher)

	// 2 HandleFunc for addPost
	http.HandleFunc("/addtopic", handlers.Post)
	// For form method post --> action "/addtopic/post"
	http.HandleFunc("/addtopic/post", handlers.GetValue)

	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	handlers.Color(1, "[SERVER_READY] : on http://localhost:8080 ✅ ")
	handlers.Color(3, "[SERVER_INFO] : To stop the program : Ctrl + c")

	http.ListenAndServe(":8080", nil)
}
