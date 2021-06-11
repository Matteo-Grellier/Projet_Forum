package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/connection.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	CreateCookie(w, r)
	fmt.Println("Page Connexion ✅")
	t.Execute(w, nil)
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Println(start)
	start2 := start.Add(time.Second*20)
	fmt.Println(start2)
	// expire := time.Now().Add(time.Hour * 1)
	c := http.Cookie{
			Name:   "ithinkidroppedacookie",
			Value:  "thedroppedcookiehasgoldinit",
		Expires: start2}
	http.SetCookie(w, &c)
	fmt.Println(c)
	// w.Write([]byte("new cookie created!\n"))
}


