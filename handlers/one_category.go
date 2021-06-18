package handlers

import (
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"
)

type DataUsed struct {
	Topics       []Topic
	ErrorMessage string
}

type Topic struct {
	Title       string
	Content     string
	User_pseudo string
}

func One_Category(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/one_category.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html")

	var DataUsedOK DataUsed

	DataUsedOK.ErrorMessage = ""

	if r.Method == "POST" {
		DataUsedOK.ErrorMessage = GetTopic(w, r).Error
	}

	DataUsedOK.Topics = DisplayTopics()

	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'one_category'")
	t.Execute(w, DataUsedOK)
}

func DisplayTopics() []Topic {
	db := BDD.OpenDataBase()
	var eachTopic Topic
	var tabTopics []Topic
	topics, err := db.Query("SELECT title, content, user_pseudo FROM topic")
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for topics.Next() {
		topics.Scan(&eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo)
		tabTopics = append(tabTopics, eachTopic)
	}
	return tabTopics
}
