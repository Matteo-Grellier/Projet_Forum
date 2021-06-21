package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/home.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/bouton_all_categories.html", "./templates/layouts/actus.html")

	userConnected := VerifyUserConnected(w, req)

	// arr := []string{"/", "/connexion", "/likedPosts", "/oneCategory", "/postsActivity", "/topic", "/inscription", "/test"}

	if req.URL.Path != "/" {
		t, _ = template.ParseFiles("./templates/layouts/error404.html")
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Page404'")
		t.Execute(w, nil)
		return
	}
	if err != nil {
		t, _ = template.ParseFiles("./templates/layouts/error500.html")
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Page500'")
		t.Execute(w, nil)
		return
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'home'")
	t.Execute(w, userConnected)
}
