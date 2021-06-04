package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./templates/home.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")

	fmt.Println("Page Home ✅")
	t.Execute(w, nil)
}
