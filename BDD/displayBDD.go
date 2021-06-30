package BDD

// Fonction permettant d'afficher les catégories venant de la base de donnée
func DisplayCategories() ([]Category, error) {
	db := OpenDataBase()
	var eachCategory Category
	var tabCategories []Category

	// On lance notre requête dans la table "Comment"
	entries, err := db.Query("SELECT name,ID FROM category")

	if err != nil {
		db.Close()
		return nil, err
	}
	for entries.Next() {
		// On stocke les données trouvées dans une variable puis on ajoute dans un tableau
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

	// On prépare notre requête à la table "topic"avec les colonnes qui nous intéressent.
	searchTopics, err := db.Prepare("SELECT ID, title, content, user_pseudo FROM topic WHERE category_id = ?")

	if err != nil {
		db.Close()
		return nil, err
	}
	// On lance notre requête à la catégorie correspondante
	topics, err := searchTopics.Query(idCat)
	if err != nil {
		db.Close()
		return nil, err
	}
	for topics.Next() {
		// On stocke les données trouvées dans une variable puis on ajoute dans un tableau
		topics.Scan(&eachTopic.ID, &eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo)
		tabTopics = append(tabTopics, eachTopic)
	}

	db.Close()
	return tabTopics, nil
}

// Fonction permettant d'afficher le dernier topic ajouté dans la base de données
func DisplayLastTopic() (int, error) {
	var topicID int
	db := OpenDataBase()
	// On lance notre requête dans la table "topic"
	searchIDTopic, err := db.Query("SELECT ID FROM topic ORDER BY rowid DESC LIMIT 1")
	if err != nil {
		db.Close()
		return 0, err
	}
	for searchIDTopic.Next() {
		// On stocke les données trouvées dans une variable
		searchIDTopic.Scan(&topicID)
	}
	db.Close()
	return topicID, nil
}

// Fonction permettant d'afficher la catégorie venant de la base de donnée
func DisplayCategory(idCat int) (string, error) {
	var nameElement string
	db := OpenDataBase()

	// On prépare notre requête à la table "category"avec les colonnes qui nous intéressent.
	searchName, err := db.Prepare("SELECT name FROM category WHERE ID = ?")
	if err != nil {
		db.Close()
		return "", err
	}
	// On demande le nom de la catégorie à l'ID correspondant
	categoryFound, err := searchName.Query(idCat)
	if err != nil {
		db.Close()
		return "", err
	}
	for categoryFound.Next() {
		// On stocke le nom trouvé dans une variable
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

	// On prépare notre requête à la table "topic" avec les colonnes qui nous intéressent.
	searchTopics, err := db.Prepare("SELECT title, content, user_pseudo, ID, category_id FROM topic WHERE ID = ?")
	if err != nil {
		db.Close()
		return eachTopic, err
	}

	// On lance notre requête à la catégorie correspondante
	topics, err := searchTopics.Query(idTopic)
	if err != nil {
		db.Close()
		return eachTopic, err
	}
	for topics.Next() {
		// On stocke les données trouvées dans la variable
		topics.Scan(&eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo, &eachTopic.ID, &eachTopic.Category_ID)
	}
	// On va chercher le nom de la catégorie correspondant au topic trouvé
	eachTopic.Category_name, BDDerror = DisplayCategory(eachTopic.Category_ID)
	if BDDerror != nil {
		db.Close()
		return eachTopic, BDDerror
	}
	topics.Close()
	return eachTopic, nil
}

// Fonction permettant d'afficher les 3 derniers posts ajoutés dans la BDD
func DisplayPostsActus() ([]Post, error) {
	db := OpenDataBase()
	var eachPost Post
	var tabPosts []Post
	// On lance notre requête dans la table "post"
	searchPosts, err := db.Prepare("SELECT content, user_pseudo, topic_id FROM post ORDER BY rowid DESC LIMIT 3")
	if err != nil {
		db.Close()
		return nil, err
	}
	Posts, err := searchPosts.Query()
	if err != nil {
		db.Close()
		return nil, err
	}
	for Posts.Next() {
		// On stocke les données trouvées dans une variable
		Posts.Scan(&eachPost.Content, &eachPost.User_pseudo, &eachPost.Topic_ID)
		tabPosts = append(tabPosts, eachPost)
	}
	return tabPosts, nil
}

// Fonction permettant d'afficher les posts venant de la base de donnée
func DisplayPosts(idTopic int, userConnected string) ([]Post, error) {
	db := OpenDataBase()
	var eachPost Post
	var tabPosts []Post
	var userLiked int

	// On prépare notre requête à la table "post" avec les colonnes qui nous intéressent.
	searchPosts, err := db.Prepare("SELECT ID, content, user_pseudo, topic_id FROM post WHERE topic_id = ?")
	if err != nil {
		db.Close()
		return nil, err
	}
	// On lance notre requête au topic correspondant
	Posts, err := searchPosts.Query(idTopic)
	if err != nil {
		db.Close()
		return nil, err
	}
	for Posts.Next() {
		// On stocke les données trouvées dans une variable puis on ajoute dans un tableau
		Posts.Scan(&eachPost.ID, &eachPost.Content, &eachPost.User_pseudo, &eachPost.Topic_ID)

		// On ajoute les commentaires du post
		eachPost.Comments, eachPost.NumberComments, err = DisplayComments(eachPost.ID)
		if err != nil {
			db.Close()
			return nil, err
		}
		// On ajoute les likes du post
		eachPost.NumberLikes, eachPost.NumberDislikes, userLiked, err = DisplayLikes(eachPost.ID, userConnected)
		if err != nil {
			db.Close()
			return nil, err
		}

		// On ajoute la condition pour afficher le like
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

	// On prépare notre requête à la table "comment" avec les colonnes qui nous intéressent.
	searchComments, err := db.Prepare("SELECT ID, content, user_pseudo FROM comment WHERE post_id = ?")
	if err != nil {
		db.Close()
		return nil, 0, err
	}
	// On lance notre requête au post correspondant
	comments, err := searchComments.Query(postId)
	if err != nil {
		db.Close()
		return nil, 0, err
	}
	for comments.Next() {
		// On stocke les données trouvées dans une variable puis on ajoute dans un tableau
		comments.Scan(&eachComment.ID, &eachComment.Content, &eachComment.User_pseudo)
		// Le compteur des commentaires prend +1 à chaque commentaire
		counter++
		tabComments = append(tabComments, eachComment)
	}
	comments.Close()
	return tabComments, counter, nil
}

// Fonction permettant d'afficher les likes venant de la base de donnée
func DisplayLikes(postID int, user_pseudo string) (int, int, int, error) {
	var counterLikes int
	var counterDislikes int

	db := OpenDataBase()
	var eachLike Likes
	var statusLike int

	// On prépare notre requête à la table "like" avec les colonnes qui nous intéressent.
	searchLikes, err := db.Prepare("SELECT liked, user_pseudo, ID FROM like WHERE post_id = ?")
	if err != nil {
		db.Close()
		return 0, 0, 0, err
	}
	// On lance notre requête au post correspondant
	foundLikes, err := searchLikes.Query(postID)
	if err != nil {
		db.Close()
		return 0, 0, 0, err
	}
	for foundLikes.Next() {
		// On stocke les données trouvées dans une variable
		foundLikes.Scan(&eachLike.Status, &eachLike.User_Pseudo, &eachLike.ID)
		// On regarde si l'utilisateur a liké ou non le post
		if eachLike.User_Pseudo == user_pseudo {
			statusLike = eachLike.Status
		}
		// S'il a agit dessus, on regarde si c'est un like ou un dislike
		if eachLike.Status == 1 {
			counterLikes++
		} else if eachLike.Status == -1 {
			counterDislikes++
		}
	}
	foundLikes.Close()
	return counterLikes, counterDislikes, statusLike, nil
}

// Fonction permettant l'affichage des posts likés par l'utilisateur
func DisplayLikedPosts(user_pseudo string) ([]Post, error) {
	db := OpenDataBase()
	var eachPostID int
	var statusLike int
	var eachPost Post
	var tabPosts []Post

	// On prépare notre requête à la table "like" avec les colonnes qui nous intéressent.
	searchLikes, err := db.Prepare("SELECT liked, post_id FROM like WHERE user_pseudo = ?")
	if err != nil {
		db.Close()
		return nil, err
	}
	// On lance notre requête à l'utilisateur concerné
	foundLikes, err := searchLikes.Query(user_pseudo)
	if err != nil {
		db.Close()
		return nil, err
	}
	for foundLikes.Next() {
		foundLikes.Scan(&statusLike, &eachPostID)
		// On vérifie le status des posts likés
		if statusLike == 1 {
			// On va chercher les informations utiles du post liké
			searchPosts, err := db.Prepare("SELECT content, user_pseudo, topic_id FROM post WHERE ID = ?")
			if err != nil {
				db.Close()
				return nil, err
			}

			Posts, err := searchPosts.Query(eachPostID)
			if err != nil {
				db.Close()
				return nil, err
			}
			for Posts.Next() {
				// On ajoute dans un tableau les données du post liké par l'utilisateur
				Posts.Scan(&eachPost.Content, &eachPost.User_pseudo, &eachPost.Topic_ID)
				tabPosts = append(tabPosts, eachPost)
			}
		}
	}
	return tabPosts, nil
}
