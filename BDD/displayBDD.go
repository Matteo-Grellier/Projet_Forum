package BDD

import (
	"log"
)

func DisplayCategories() []Category {
	db := OpenDataBase()
	var eachCategory Category
	var tabCategories []Category
	entries, err := db.Query("SELECT name,ID FROM category")

	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
		// return
	}
	for entries.Next() {
		entries.Scan(&eachCategory.Name, &eachCategory.Id)
		tabCategories = append(tabCategories, eachCategory)
	}
	db.Close()
	return tabCategories
}

func DisplayTopics(idCat int) []Topic {
	db := OpenDataBase()
	var eachTopic Topic
	var tabTopics []Topic
	searchTopics, err := db.Prepare("SELECT ID, title, content, user_pseudo FROM topic WHERE category_id = ?")
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}

	topics, err := searchTopics.Query(idCat)
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for topics.Next() {
		topics.Scan(&eachTopic.ID, &eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo)
		tabTopics = append(tabTopics, eachTopic)
	}
	return tabTopics
}
func DisplayTopicID() int {
	var topicID int
	db := OpenDataBase()
	searchIDTopic, err := db.Query("SELECT ID FROM topic ORDER BY rowid DESC LIMIT 1")
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for searchIDTopic.Next() {
		searchIDTopic.Scan(&topicID)
	}
	return topicID
}

func DisplayCategory(idCat int) string {
	var nameElement string
	db := OpenDataBase()
	searchName, err := db.Prepare("SELECT name FROM category WHERE ID = ?")
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	categoryFound, err := searchName.Query(idCat)
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for categoryFound.Next() {
		categoryFound.Scan(&nameElement)
	}
	if nameElement == "" {
		return "nil"
	}

	return nameElement
}
