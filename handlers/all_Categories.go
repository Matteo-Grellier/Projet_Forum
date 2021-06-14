package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func All_Categories(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/all_categories.html", "templates/layouts/sidebar.html", "templates/layouts/header.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	fmt.Println("Page All Catégories ✅")
	t.Execute(w, nil)
}
