package BDD

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Pseudo   string
	Mail     string
	Password string
}
type DataUsed struct {
	Users  []User
	Topics []Topic
}

type Topic struct {
	ID          int
	Title       string
	Content     string
	User_pseudo string
	Category_ID int
}

func AddUser(pseudo string, mail string, password string) {
	db := OpenDataBase()
	createNew, err := db.Prepare("INSERT INTO user (pseudo, mail, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = createNew.Exec(pseudo, mail, password)
	if err != nil {
		log.Fatal(err)
	}
}

// Vérifie si l'élément demandé est déjà dans la base de données
func VerifyBDD(element string, column string) (bool, string) {
	db := OpenDataBase()
	var oneElement string
	var prepareElements *sql.Stmt
	var errorPrepare error
	if column == "pseudo" {
		prepareElements, errorPrepare = db.Prepare("SELECT pseudo FROM user")
	} else if column == "mail" {
		prepareElements, errorPrepare = db.Prepare("SELECT mail FROM user")
	} else if column == "session" {
		prepareElements, errorPrepare = db.Prepare("SELECT user_pseudo FROM session")
	}

	if errorPrepare != nil {
		log.Fatal(errorPrepare)
	}

	allElements, err := prepareElements.Query()
	if err != nil {
		log.Fatalf("%s", err)
	}
	for allElements.Next() {
		allElements.Scan(&oneElement)
		fmt.Println("Pseudo ou mail = ", oneElement)
		if oneElement == element {
			ErrorMessage := column + " déjà dans la base de données."
			allElements.Close()
			return true, ErrorMessage
		}
	}
	allElements.Close()
	return false, ""
}

// vérification de la correspondance entre le mdp de l'utilisateur et le mdp entré
func VerifyPassword(password string, pseudo string) bool {
	db := OpenDataBase()
	var eachPseudo User
	selectPassword, _ := db.Prepare("SELECT password FROM user WHERE pseudo = ?")
	verifPassword, err := selectPassword.Query(pseudo)
	if err != nil {
		log.Fatal(err)
	}
	for verifPassword.Next() {
		verifPassword.Scan(&eachPseudo.Password)
		err := bcrypt.CompareHashAndPassword([]byte(eachPseudo.Password), []byte(password))
		if err != nil {
			log.Println(err)
			verifPassword.Close()
			return false
		} else {
			verifPassword.Close()
			return true
		}
	}
	db.Close()
	return false
}

// Ajout de l'UUID dans la base de données
func AddUUID(pseudo string, UUID string) {
	db := OpenDataBase()
	var actionBDD *sql.Stmt
	var errAction error

	// Vérification de la présence de l'utilisateur dans la table session
	correctPseudo, _ := VerifyBDD(pseudo, "session")
	if correctPseudo {
		actionBDD, errAction = db.Prepare("UPDATE session SET UUID = ? WHERE user_pseudo = ?")
	} else {
		actionBDD, errAction = db.Prepare("INSERT INTO session (UUID, user_pseudo) VALUES(?, ?)")
	}
	if errAction != nil {
		log.Fatal(errAction)
	}

	//Ajout ou update de l'UUID à l'utilisateur connecté
	_, err := actionBDD.Exec(UUID, pseudo)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func Afficher(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./templates/BDD.html")
	DataUsedOK := DataUsed{
		Users:  SelectUsers(),
		Topics: SelectTopics(),
	}

	if req.Method == "POST" {
		if req.FormValue("delete") == "delete" {
			delete()
		} else if req.FormValue("create") == "create" {
			create()
		}
	}
	AjoutCommentaires()
	t.Execute(w, DataUsedOK)
}

func OpenDataBase() *sql.DB {
	/*Ouverture de la base de données*/
	db, err := sql.Open("sqlite3", "./BDD/BDDForum1.db")
	if err != nil {
		fmt.Println("Could Not open Database")
	}
	return db
}

func SelectUsers() []User {
	/*Sélection du pseudo et de l'adresse mail de tous les utilisateurs*/
	db := OpenDataBase()
	var eachUser User
	var tabUsers []User
	entries, err := db.Query("SELECT pseudo, mail FROM user")

	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for entries.Next() {
		entries.Scan(&eachUser.Pseudo, &eachUser.Mail)
		tabUsers = append(tabUsers, eachUser)
	}
	entries.Close()
	return tabUsers
}
func SelectTopics() []Topic {
	db := OpenDataBase()
	var eachTopic Topic
	var tabTopics []Topic
	topics, err := db.Query("SELECT * FROM topic")
	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for topics.Next() {
		topics.Scan(&eachTopic.ID, &eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo, &eachTopic.Category_ID)
		tabTopics = append(tabTopics, eachTopic)
	}
	topics.Close()
	return tabTopics

}

func create() {
	db := OpenDataBase()
	creation, _ := db.Prepare("INSERT INTO user (pseudo, mail, password) VALUES(?, ?, ?)")
	creation.Exec("pseudo", "mail@gmail.com", "password")
	creation.Close()
}

func delete() {
	db := OpenDataBase()
	delete, _ := db.Prepare("DELETE FROM ? WHERE ? = ?")
	_, err := delete.Exec()
	if err != nil {
		log.Fatal(err)
	}
	delete.Close()
}
func AjoutCommentaires() {
	db := OpenDataBase()
	creation, _ := db.Prepare("INSERT INTO comment (user_pseudo, content, post_id) VALUES(?, ?, ?)")
	// _, err := creation.Exec()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	creation.Close()
	// Modifier les ? en fonction de ce qu'on veut supprimer
}
