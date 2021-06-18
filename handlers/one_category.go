package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func One_Category(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/one_category.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	fmt.Println("Page Connexion ✅")
	t.Execute(w, nil)
}
