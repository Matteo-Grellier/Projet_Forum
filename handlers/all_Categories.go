package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func All_Categories(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/all_categories.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'all_catégories'")
	t.Execute(w, nil)
}
