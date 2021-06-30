package BDD

// Fonction permettant d'afficher les catégories venant de la base de donnée
func DisplayCategories() ([]Category, error) {
	db := OpenDataBase()
	var eachCategory Category
	var tabCategories []Category
	entries, err := db.Query("SELECT name,ID FROM category")

	if err != nil {
		db.Close()
		return nil, err
	}
	for entries.Next() {
		entries.Scan(&eachCategory.Name, &eachCategory.Id)
		tabCategories = append(tabCategories, eachCategory)
	}
	db.Close()
	return tabCategories, nil
}

// Fonction permettant d'afficher les topics venant de la base de donnée
func DisplayTopics(idCat int) ([]Topic, error) {
	db := OpenDataBase()
	var eachTopic Topic
	var tabTopics []Topic
	searchTopics, err := db.Prepare("SELECT ID, title, content, user_pseudo FROM topic WHERE category_id = ?")

	if err != nil {
		db.Close()
		return nil, err
	}

	topics, err := searchTopics.Query(idCat)
	if err != nil {
		db.Close()
		return nil, err
	}
	for topics.Next() {
		topics.Scan(&eachTopic.ID, &eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo)
		tabTopics = append(tabTopics, eachTopic)
	}

	db.Close()
	return tabTopics, nil
}

// Fonction permettant d'afficher les topics suivant les ID venant de la base de donnée
func DisplayTopicID() (int, error) {
	var topicID int
	db := OpenDataBase()
	searchIDTopic, err := db.Query("SELECT ID FROM topic ORDER BY rowid DESC LIMIT 1")
	if err != nil {
		db.Close()
		return 0, err
	}
	for searchIDTopic.Next() {
		searchIDTopic.Scan(&topicID)
	}
	db.Close()
	return topicID, nil
}

// Fonction permettant d'afficher la catégorie venant de la base de donnée
func DisplayCategory(idCat int) (string, error) {
	var nameElement string
	db := OpenDataBase()
	searchName, err := db.Prepare("SELECT name FROM category WHERE ID = ?")
	if err != nil {
		db.Close()
		return "", err
	}
	categoryFound, err := searchName.Query(idCat)
	if err != nil {
		db.Close()
		return "", err
	}
	for categoryFound.Next() {
		categoryFound.Scan(&nameElement)
	}
	if nameElement == "" {
		categoryFound.Close()
		return "nil", nil
	}
	categoryFound.Close()

	return nameElement, nil
}

// Fonction permettant d'afficher un topic suivant l'ID venant de la base de donnée
func DisplayOneTopic(idTopic int) (Topic, error) {
	db := OpenDataBase()
	var eachTopic Topic
	var BDDerror error
	searchTopics, err := db.Prepare("SELECT title, content, user_pseudo, ID, category_id FROM topic WHERE ID = ?")
	if err != nil {
		db.Close()
		return eachTopic, err
	}

	topics, err := searchTopics.Query(idTopic)
	if err != nil {
		db.Close()
		return eachTopic, err
	}
	for topics.Next() {
		topics.Scan(&eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo, &eachTopic.ID, &eachTopic.Category_ID)
	}
	eachTopic.Category_name, BDDerror = DisplayCategory(eachTopic.Category_ID)
	if BDDerror != nil {
		db.Close()
		return eachTopic, BDDerror
	}
	topics.Close()
	return eachTopic, nil
}

// Fonction permettant d'afficher les posts venant de la base de donnée
func DisplayPosts(idTopic int, userConnected string) ([]Post, error) {
	db := OpenDataBase()
	var eachPost Post
	var tabPosts []Post
	var userLiked int
	searchPosts, err := db.Prepare("SELECT ID, content, user_pseudo, topic_id FROM post WHERE topic_id = ?")
	if err != nil {
		db.Close()
		return nil, err
	}

	Posts, err := searchPosts.Query(idTopic)
	if err != nil {
		db.Close()
		return nil, err
	}
	for Posts.Next() {
		Posts.Scan(&eachPost.ID, &eachPost.Content, &eachPost.User_pseudo, &eachPost.Topic_ID)

		eachPost.Comments, eachPost.NumberComments, err = DisplayComments(eachPost.ID)
		if err != nil {
			db.Close()
			return nil, err
		}
		eachPost.NumberLikes, eachPost.NumberDislikes, userLiked, err = DisplayLikes(eachPost.ID, userConnected)
		if err != nil {
			db.Close()
			return nil, err
		}

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

	return tabPosts, nil
}

// Fonction permettant d'afficher les commentaires venant de la base de donnée
func DisplayComments(postId int) ([]Comment, int, error) {
	db := OpenDataBase()
	var eachComment Comment
	var counter int
	var tabComments []Comment
	searchComments, err := db.Prepare("SELECT ID, content, user_pseudo FROM comment WHERE post_id = ?")
	if err != nil {
		db.Close()
		return nil, 0, err
	}
	comments, err := searchComments.Query(postId)
	if err != nil {
		db.Close()
		return nil, 0, err
	}
	for comments.Next() {
		comments.Scan(&eachComment.ID, &eachComment.Content, &eachComment.User_pseudo)
		counter++
		tabComments = append(tabComments, eachComment)
	}
	comments.Close()
	return tabComments, counter, nil
}

// Fonction permettant d'afficher les likes venant de la base de donnée
func DisplayLikes(postID int, user_pseudo string) (int, int, int, error) {
	// postID := 227
	var counterLikes int
	var counterDislikes int

	db := OpenDataBase()
	var eachLike Likes
	var statusLike int
	searchLikes, err := db.Prepare("SELECT liked, user_pseudo, ID FROM like WHERE post_id = ?")
	if err != nil {
		db.Close()
		return 0, 0, 0, err
	}
	foundLikes, err := searchLikes.Query(postID)
	if err != nil {
		db.Close()
		return 0, 0, 0, err
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
	return counterLikes, counterDislikes, statusLike, nil
}
