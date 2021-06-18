package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/home.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/all_categories.html", "./templates/layouts/actus.html")

	userConnected := VerifyUserConnected(w, req)
	fmt.Println(userConnected)

	arr := []string{"/", "/connexion", "/likedPosts", "/oneCategory", "/postsActivity", "/topic", "/inscription", "/test"}

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

	t.Execute(w, userConnected)
}
