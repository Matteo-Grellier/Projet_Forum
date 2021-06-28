package handlers

import (
	"log"
	"net/http"
	"time"

	BDD "../BDD"
	guuid "github.com/google/uuid"
)

// Création du cookie. Renvoie la valeur du cookie.
// Fonction appelée lorsqu'un utilisateur se connecte.
func CreateCookie(w http.ResponseWriter, r *http.Request, username string) {
	id := guuid.New()
	start := time.Now()
	start2 := start.Add(time.Minute * 15)
	c := http.Cookie{
		Name:    "CookieSession",
		Value:   id.String(),
		Expires: start2}
	http.SetCookie(w, &c)

	// Ajout de l'UUID dans la base de données.
	BDD.AddUUID(username, id.String())
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
