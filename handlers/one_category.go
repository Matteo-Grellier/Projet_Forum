package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"
)

type DataUsed struct {
	Topics []Topic
}

type Topic struct {
	Title       string
	Content     string
	User_pseudo string
}

func One_Category(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/one_category.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html")
	DataUsedOK := DataUsed{
		Topics: DisplayTopics(),
	}

	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	fmt.Println("Page Connexion ✅")
	t.Execute(w, DataUsedOK)
}

func DisplayTopics() []Topic {
	db := BDD.OpenDataBase()
	var eachTopic Topic
	var tabTopics []Topic
	topics, err := db.Query("SELECT title, content, user_pseudo FROM topic")
	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for topics.Next() {
		topics.Scan(&eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo)
		tabTopics = append(tabTopics, eachTopic)
	}
	return tabTopics
}
