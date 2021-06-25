package main

import (
	handlers "./handlers"

	"net/http"

	BDD "./BDD"
)

func main() {

	handlers.Color(2, "[SERVER_INFO] : Starting local Server...")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/connexion", handlers.ConnexionPage)
	http.HandleFunc("/inscription", handlers.InscriptionPage)
	http.HandleFunc("/categories", handlers.CategoriesPage)
	http.HandleFunc("/oneCategory", handlers.OneCategoryPage)

	// http.HandleFunc("/postsActivity", handlers.Posts_Activity)
	http.HandleFunc("/topic", handlers.TopicPage)
	http.HandleFunc("/BDD", BDD.Afficher)

	// http.HandleFunc("/likedPosts", handlers.Liked_Posts)

	// Fonctions exécutées après une requête
	http.HandleFunc("/deconnexion", handlers.GetDeconnected)

	// For form method post --> action "/addtopic/post"
	http.HandleFunc("/addtopic/post", handlers.GetValue)

	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	handlers.Color(1, "[SERVER_READY] : on http://localhost:8080 ✅ ")
	handlers.Color(3, "[SERVER_INFO] : To stop the program : Ctrl + c")

	http.ListenAndServe(":8080", nil)
}
