package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var t *template.Template
var tErr *template.Template

func InscriptionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/inscription.html", "templates/layouts/sidebar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	fmt.Println("Page Connexion ✅")
	t.Execute(w, nil)
}
