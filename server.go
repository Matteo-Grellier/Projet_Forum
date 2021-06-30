package main

import (
	handlers "./handlers"

	"net/http"
)

func main() {

	handlers.Color(2, "[SERVER_INFO] : Starting local Server...")

	// Déclaration des URL et des fonctions associées
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/connexion", handlers.ConnexionPage)
	http.HandleFunc("/inscription", handlers.InscriptionPage)
	http.HandleFunc("/categories", handlers.CategoriesPage)
	http.HandleFunc("/oneCategory", handlers.OneCategoryPage)
	http.HandleFunc("/topic", handlers.OneTopicPage)
	http.HandleFunc("/likes", handlers.LikesPage)

	// Fonction exécutée lors d'une déconnexion
	http.HandleFunc("/deconnexion", handlers.GetDeconnected)

	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	handlers.Color(1, "[SERVER_READY] : on http://localhost:8080 ✅ ")
	handlers.Color(3, "[SERVER_INFO] : To stop the program : Ctrl + c")

	http.ListenAndServe(":8080", nil)
}
