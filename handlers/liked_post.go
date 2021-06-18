package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Liked_Posts(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/liked_posts.html", "templates/layouts/sidebar.html")
	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'liked_post'")
	t.Execute(w, nil)
}
