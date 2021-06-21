package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Commentaires(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/commentaires.html")
	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'posts_activity'")
	t.Execute(w, nil)
}
