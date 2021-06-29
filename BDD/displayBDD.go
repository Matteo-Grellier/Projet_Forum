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

func DisplayPosts(idTopic int, userConnected string) []Post {
	db := OpenDataBase()
	var eachPost Post
	var tabPosts []Post
	var userLiked int
	searchPosts, err := db.Prepare("SELECT ID, content, user_pseudo, topic_id FROM post WHERE topic_id = ?")
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
		Posts.Scan(&eachPost.ID, &eachPost.Content, &eachPost.User_pseudo, &eachPost.Topic_ID)
		eachPost.Comments, eachPost.NumberComments = DisplayComments(eachPost.ID)
		eachPost.NumberLikes, eachPost.NumberDislikes, userLiked = DisplayLikes(eachPost.ID, userConnected)

		if userLiked == 1 {
			eachPost.UserLiked = true
			eachPost.UserDisliked = false
		} else if userLiked == -1 {
			eachPost.UserDisliked = true
			eachPost.UserLiked = false
		} else if userLiked == 0 {
			eachPost.UserDisliked = false
			eachPost.UserLiked = false

		}

		tabPosts = append(tabPosts, eachPost)
	}

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

func DisplayLikes(postID int, user_pseudo string) (int, int, int) {
	// postID := 227
	var counterLikes int
	var counterDislikes int

	db := OpenDataBase()
	var eachLike Likes
	var statusLike int
	searchLikes, err := db.Prepare("SELECT liked, user_pseudo, ID FROM like WHERE post_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	foundLikes, err := searchLikes.Query(postID)
	if err != nil {
		log.Fatal(err)
	}
	for foundLikes.Next() {
		foundLikes.Scan(&eachLike.Status, &eachLike.User_Pseudo, &eachLike.ID)
		if eachLike.User_Pseudo == user_pseudo {
			statusLike = eachLike.Status
		}
		if eachLike.Status == 1 {
			counterLikes++
		} else if eachLike.Status == -1 {
			counterDislikes++
		}
	}
	foundLikes.Close()
	return counterLikes, counterDislikes, statusLike
}
