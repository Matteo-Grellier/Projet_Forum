package handlers

import (
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"
)

type DataComment struct {
	Comments     []Comment
	ErrorMessage string
}

type Comment struct {
	Content     string
	User_pseudo string
}

func Commentaires(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/commentaires.html")

	var DataCommentOK DataComment

	DataCommentOK.ErrorMessage = ""

	if r.Method == "POST" {
		addComment(w, r)
	}

	DataCommentOK.Comments = DisplayComments()

	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'posts_activity'")
	t.Execute(w, DataCommentOK)
}

func DisplayComments() []Comment {
	db := BDD.OpenDataBase()
	var eachComment Comment
	var tabComments []Comment
	comments, err := db.Query("SELECT content, user_pseudo FROM comment")
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for comments.Next() {
		comments.Scan(&eachComment.Content, &eachComment.User_pseudo)
		tabComments = append(tabComments, eachComment)
	}
	return tabComments
}

func addComment(w http.ResponseWriter, r *http.Request) Errors {
	comment := r.FormValue("comment")
	//TEST BRUT
	user := "Roberto04"
	postId := "227"

	var data = []string{comment}

	var ErrorsPost Errors

	if verifyInput(data) {
		db := BDD.OpenDataBase()
		createNew, err3 := db.Prepare("INSERT INTO Comment (content, user_pseudo, post_id) VALUES (?, ?, ?)")
		if err3 != nil {
			Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
			log.Fatal(err3)
		}
		createNew.Exec(comment, user, postId)
	} else {
		ErrorsPost.Error = ErrorMessage
		ErrorMessage = ""
	}
	return ErrorsPost
}
