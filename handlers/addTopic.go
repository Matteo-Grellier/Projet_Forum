package handlers

import (
	BDD "../BDD"
)

func AddTopic(titre string, content string, categId int, user string) (string, int) {

	var data = []string{titre, content}
	if verifyInput(data) {
		BDD.AddTopic(titre, content, user, categId)
	} else {
		return "Il manque un item.", 0
	}
	return "", BDD.DisplayTopicID()
}
