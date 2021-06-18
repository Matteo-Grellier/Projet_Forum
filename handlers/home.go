package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func Home(w http.ResponseWriter, req *http.Request) {

	arr := []string{"/", "/connexion", "/likedPosts", "/postsActivity", "/topic", "/inscription", "/test"}
	foo := 102
	for i := 0; i < 150; i++ {
		foo++
		arr = append(arr, "/oneCategory="+strconv.Itoa(foo))
	}
	fmt.Println(arr)

	t, err := template.ParseFiles("./templates/home.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/actus.html", "./templates/layouts/all_categories.html")
	fmt.Println(req.URL.Path)
	test := string(req.URL.Path)
	if itemExists(arr, test) {
		fmt.Println("Page Home ✅")
		t.Execute(w, nil)
	} else {
		t, _ = template.ParseFiles("./templates/layouts/error404.html")
		t.Execute(w, nil)
		return
	}
	if err != nil {
		t, _ = template.ParseFiles("./templates/layouts/error500.html")
		t.Execute(w, nil)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'home'")
	t.Execute(w, nil)
}
func itemExists(arr []string, test string) bool {
	isInArr := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == test {
			isInArr = true
		}
	}
	return isInArr
}
