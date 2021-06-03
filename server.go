package main

import (
	"fmt"
	"net/http"

	handlers "./handlers"
	all_categories "./src/all_categories"
	connexion "./src/connection"
	inscription "./src/inscription"
	likedPosts "./src/liked_posts"
	one_category "./src/one_category"
	posts_activity "./src/posts_activity"
	topic "./src/topic"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/connexion", connexion.ConnexionPage)
	http.HandleFunc("/likedPosts", likedPosts.Liked_Posts)
	http.HandleFunc("/oneCategory", one_category.One_Category)
	http.HandleFunc("/postsActivity", posts_activity.Posts_Activity)
	http.HandleFunc("/topic", topic.TopicPage)
	http.HandleFunc("/inscription", inscription.InscriptionPage)
	http.HandleFunc("/all_categories", all_categories.All_Categories)
	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
