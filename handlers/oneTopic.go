package handlers

import (
	"log"

	BDD "../BDD"
)

type TopicDataUsed struct {
	Topics []TopicStruct

	ErrorMessage string
}

type TopicStruct struct {
	ID          int
	Title       string
	Content     string
	User_pseudo string
}

func DisplayOneTopic(idCat int) []TopicStruct {
	db := BDD.OpenDataBase()
	var eachTopic TopicStruct
	var tabTopics []TopicStruct
	searchTopics, err := db.Prepare("SELECT title, content, user_pseudo, ID FROM topic WHERE ID = ?")
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}

	topics, err := searchTopics.Query(idCat)
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for topics.Next() {
		topics.Scan(&eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo, &eachTopic.ID)
		tabTopics = append(tabTopics, eachTopic)
	}
	return tabTopics
}
