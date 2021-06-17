package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type checkConnexion struct {
	Pseudo string
	Log    string
}

func Home(w http.ResponseWriter, req *http.Request) {

	arr := []string{"/", "/connexion", "/likedPosts", "/oneCategory", "/postsActivity", "/topic", "/inscription", "/test"}

	t, err := template.ParseFiles("./templates/home.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	for i := 0; i < len(arr); i++ {
		if req.URL.Path != arr[i] {
			t, _ = template.ParseFiles("./templates/layouts/error404.html")
			t.Execute(w, nil)
			return
		} else if req.URL.Path == arr[i] {
			break
		}
	}
	if err != nil {
		t, _ = template.ParseFiles("./templates/layouts/error500.html")
		t.Execute(w, nil)
		return
	}
	fmt.Println("Page Home ✅")

	// Vérification de l'utilisateur connecté

	pseudo, connected := VerifyUserConnected(w, req)
	var userConnected checkConnexion
	if connected {
		userConnected = checkConnexion{
			Pseudo: pseudo,
			Log:    "Logout",
		}
		fmt.Println(userConnected)
		fmt.Println("Utilisateur connecté")

	} else {
		userConnected = checkConnexion{
			Log: "Login",
		}
		fmt.Println("Pas d'utilisateur connecté")
	}
	t.Execute(w, userConnected)
}
