package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./templates/home.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if req.URL.Path != "/" {
		fmt.Printf("error 404 %s", err)
		t, _ = template.ParseFiles("./templates/layouts/error.html")
		t.Execute(w, nil)
		return
	}
	fmt.Println("Page Home ✅")
	t.Execute(w, nil)
}
