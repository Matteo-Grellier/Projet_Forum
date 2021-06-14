package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

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
	t.Execute(w, nil)
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Println(start)
	start2 := start.Add(time.Second*20)
	fmt.Println(start2)
	c := http.Cookie{
			Name:   "ithinkidroppedacookie",
			Value:  "thedroppedcookiehasgoldinit",
			Expires: start2}
	http.SetCookie(w, &c)

	expire := time.Now().Add(20 * time.Minute) // Expires in 20 minutes
	cookie := http.Cookie{Name: "username", Value: "nonsecureuser", Path: "/", Expires: expire, MaxAge: 86400}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "secureusername", Value: "secureuser", Path: "/", Expires: expire, MaxAge: 86400, HttpOnly: true, Secure: true}
	http.SetCookie(w, &cookie)	

	fmt.Println(c)
}