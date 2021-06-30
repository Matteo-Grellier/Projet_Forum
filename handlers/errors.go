package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

// Fonction qui exécutera la page 404
func Error404(w http.ResponseWriter, req *http.Request) bool {

	arr := []string{"/", "/connexion", "/likedPosts", "/oneCategory", "/postsActivity", "/topic", "/inscription", "/test", "/categories"}
	compteurURL := 0
	for i := 0; i < len(arr); i++ {
		if req.URL.Path != arr[i] {
			compteurURL++
		} else if req.URL.Path == arr[i] {
			break
		}
	}
	if compteurURL == len(arr) {
		t, _ := template.ParseFiles("./templates/layouts/error404.html", "./templates/layouts/header.html", "./templates/layouts/sidebar.html")
		Color(4, "[SERVER_INFO_PAGE] : 🔴 Page 'Page404' : cette page n'existe pas")
		t.Execute(w, nil)
		return false
	}
	return true
}

// Fonction qui exécutera la page 404 si les données ne sont pas renvoyés
func NoItemsError(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./templates/layouts/error404.html", "./templates/layouts/header.html", "./templates/layouts/sidebar.html")
	Color(4, "[SERVER_INFO_PAGE] : 🔴 Page 'Page404' : item not found")
	t.Execute(w, nil)
}

// Fonction qui exécutera la page 500
func Error500(w http.ResponseWriter, req *http.Request, err error) {
	Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
	fmt.Println(err)
	t, _ := template.ParseFiles("./templates/layouts/error500.html", "./templates/layouts/header.html", "./templates/layouts/sidebar.html")
	t.Execute(w, nil)
}
