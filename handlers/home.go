package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./templates/home.html", "./templates/layouts/actus.html", "./templates/layouts/sidebar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	fmt.Println("Page Home ✅")
	t.Execute(w, nil)
}
