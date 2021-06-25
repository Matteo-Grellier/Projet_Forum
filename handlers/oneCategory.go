package handlers

import (
	"fmt"

	BDD "../BDD"
)

type DataPageCategory struct {
	Topics     []BDD.Topic
	Category   string
	CategoryID int
	Error      string
}

func AddTopic(titre string, content string, categId int, user string) (string, int) {

	var data = []string{titre, content}
	fmt.Println(verifyInput(data))
	if verifyInput(data) {
		BDD.AddTopic(titre, content, user, categId)
	} else {
		return "Il manque un item.", 0
	}
	return "", BDD.DisplayTopicID()
}
