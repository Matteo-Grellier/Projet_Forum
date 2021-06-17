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

	fmt.Println(req.URL.Path)
	test := string(req.URL.Path)
	if itemExists(arr, test) {
		t, err := template.ParseFiles("./templates/one_category.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/actus.html", "./templates/layouts/all_categories.html")
		if err != nil {
			t, _ = template.ParseFiles("./templates/layouts/error500.html")
			t.Execute(w, nil)
			return
		}
		fmt.Println("Page Home ✅")
		t.Execute(w, nil)
	} else {
		t, _ := template.ParseFiles("./templates/layouts/error404.html")
		t.Execute(w, nil)
		return
	}

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
