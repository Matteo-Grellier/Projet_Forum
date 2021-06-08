package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Liked_Posts(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/liked_posts.html", "templates/layouts/sidebar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	fmt.Println("Page Connexion ✅")
	t.Execute(w, nil)
}
