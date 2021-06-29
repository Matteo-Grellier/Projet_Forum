package handlers

import (
	BDD "../BDD"
)

func AddTopic(titre string, content string, categId int, user string) (string, int, error) {

	var data = []string{titre, content}
	if verifyInput(data) {
		err := BDD.AddTopic(titre, content, user, categId)
		if err != nil {
			return "", 0, err
		}
	} else {
		return "Il manque un item.", 0, nil
	}
	topicID, err := BDD.DisplayTopicID()
	if err != nil {
		return "", 0, err
	}
	return "", topicID, nil
}
