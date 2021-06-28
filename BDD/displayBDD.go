package BDD

import (
	"fmt"
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
		categoryFound.Close()
		return "nil"
	}
	categoryFound.Close()

	return nameElement
}

func DisplayOneTopic(idTopic int) Topic {
	db := OpenDataBase()
	var eachTopic Topic
	searchTopics, err := db.Prepare("SELECT title, content, user_pseudo, ID, category_id FROM topic WHERE ID = ?")
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}

	topics, err := searchTopics.Query(idTopic)
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for topics.Next() {
		topics.Scan(&eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo, &eachTopic.ID, &eachTopic.Category_ID)
	}
	eachTopic.Category_name = DisplayCategory(eachTopic.Category_ID)
	topics.Close()
	return eachTopic
}

func DisplayPosts(idTopic int) []Post {
	db := OpenDataBase()
	var eachPost Post
	var tabPosts []Post
	searchPosts, err := db.Prepare("SELECT ID, content, user_pseudo FROM post WHERE topic_id = ?")
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}

	Posts, err := searchPosts.Query(idTopic)
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for Posts.Next() {
		Posts.Scan(&eachPost.ID, &eachPost.Content, &eachPost.User_pseudo)
		eachPost.Comments, eachPost.NumberComments = DisplayComments(eachPost.ID)
		tabPosts = append(tabPosts, eachPost)
	}
	fmt.Println("--------------------")
	fmt.Println(tabPosts)
	return tabPosts
}

func DisplayComments(postId int) ([]Comment, int) {
	db := OpenDataBase()
	var eachComment Comment
	var counter int
	var tabComments []Comment
	searchComments, err := db.Prepare("SELECT ID, content, user_pseudo FROM comment WHERE post_id = ?")
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	comments, err := searchComments.Query(postId)
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for comments.Next() {
		comments.Scan(&eachComment.ID, &eachComment.Content, &eachComment.User_pseudo)
		counter++
		tabComments = append(tabComments, eachComment)
	}
	comments.Close()
	return tabComments, counter
}

func DisplayLikes() Likes {
	db := OpenDataBase()
	var like Likes
	searchLikes, err := db.Query("SELECT liked FROM Like WHERE ID = 100")
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	if err != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for searchLikes.Next() {
		searchLikes.Scan(&like.liked)
	}
	searchLikes.Close()
	fmt.Println("like !!!!!!!!!!")
	fmt.Println()
	return like
}
