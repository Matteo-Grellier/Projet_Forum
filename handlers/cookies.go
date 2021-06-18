package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	BDD "../BDD"
	guuid "github.com/google/uuid"
)

type UserConnectedStruct struct {
	PseudoConnected string
	Connected       bool
}

// Création du cookie. Renvoie la valeur du cookie.
// Fonction appelée lorsqu'un utilisateur se connecte.
func CreateCookie(w http.ResponseWriter, r *http.Request) string {
	id := guuid.New()
	start := time.Now()
	start2 := start.Add(time.Minute * 1)
	c := http.Cookie{
		Name:    "CookieSession",
		Value:   id.String(),
		Expires: start2}
	http.SetCookie(w, &c)

	return id.String()
}

// Ajout de l'UUID dans la table session
//Fonction appelée lors de la connexion d'un utilisateur
func CreateUUID(username string, newUUID string, db *sql.DB) {
	var usernameFound string
	var user_registered bool
	verifyUser, err := db.Query("SELECT user_pseudo FROM session")
	if err != nil {
		log.Fatal(err)
	}
	for verifyUser.Next() {
		verifyUser.Scan(&usernameFound)
		if usernameFound == username {
			user_registered = true
			break
		}
	}
	verifyUser.Close()

	if user_registered {
		update, err := db.Prepare("UPDATE session SET UUID = ? WHERE user_pseudo = ?")
		if err != nil {
			log.Fatal(err)
		}
		_, err = update.Exec(newUUID, username)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		add, err := db.Prepare("INSERT INTO session (UUID, user_pseudo) VALUES(?, ?)")

		if err != nil {
			log.Fatal(err)
		}
		_, err = add.Exec(newUUID, username)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// Renvoie la valeur du cookie actif. Renvoie une erreur lorsqu'il n'y a plus de cookie.
// Utilisé lors du chargement d'une page
func ReadCookie(w http.ResponseWriter, r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		if err == http.ErrNoCookie {
			return ""
		}
		return ""
	}
	return c.Value
}

// Retourne le nom de l'utilisateur connecté. Utilise la fonction ReadCookie pour savoir sa valeur.
// Parcourt la BDD pour savoir quel utilisateur est lié au cookie de session.
// Utilisé à chaque chargement de page.
func VerifyUserConnected(w http.ResponseWriter, r *http.Request) UserConnectedStruct {
	var userConnected UserConnectedStruct
	CookieValue := ReadCookie(w, r, "CookieSession")
	if CookieValue == "" {
		userConnected.PseudoConnected = ""
		userConnected.Connected = false
		return userConnected
	}

	var user_connected string
	db := BDD.OpenDataBase()
	checkUUID, err := db.Prepare("SELECT user_pseudo FROM session WHERE UUID = ?")
	if err != nil {
		log.Fatal(err)
	}
	check_user, err := checkUUID.Query(CookieValue)
	if err != nil {
		log.Fatal(err)
	}

	for check_user.Next() {
		check_user.Scan(&user_connected)
	}
	checkUUID.Close()
	userConnected.PseudoConnected = user_connected
	userConnected.Connected = true
	return userConnected
}

// Supprime le cookie actif
// Fonction appelée lorsque l'utilisateur souhaite se déconnecter
func GetDeconnected(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "CookieSession",
		MaxAge: -1}
	http.SetCookie(w, &c)
	Color(1, "[CONNEXION] : 🟢 Vous êtes bien déconnectés ")
	http.Redirect(w, r, "/connexion", http.StatusSeeOther)
}
