package handlers

import (
	"log"
	"net/http"
	"time"

	BDD "../BDD"
	guuid "github.com/google/uuid"
)

// Permet de créer un cookie. Fonction appelée lorsqu'un utilisateur se connecte.
func CreateCookie(w http.ResponseWriter, r *http.Request, username string) error {
	id := guuid.New()
	start := time.Now()
	end := start.Add(time.Minute * 15)
	c := http.Cookie{
		Name:    "CookieSession",
		Value:   id.String(),
		Expires: end}
	http.SetCookie(w, &c)

	// Ajout de l'UUID dans la base de données.
	return BDD.AddUUID(username, id.String())
}

// Renvoie la valeur du cookie actif. Renvoie une erreur lorsqu'il n'y a plus de cookie.
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

// Permet de savoir si un utilisateur est connecté et son pseudo
func VerifyUserConnected(w http.ResponseWriter, r *http.Request) UserConnectedStruct {
	var userConnected UserConnectedStruct
	// On lit le cookie de la page pour avoir sa valeur
	CookieValue := ReadCookie(w, r, "CookieSession")
	if CookieValue == "" {
		userConnected.PseudoConnected = ""
		userConnected.Connected = false
		return userConnected
	}
	var user_connected string
	db := BDD.OpenDataBase()

	// On parcourt la BDD pour savoir quel utilisateur est lié au cookie de session.
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

	// On retourne le pseudo et le status de la connexion
	userConnected.PseudoConnected = user_connected
	userConnected.Connected = true
	return userConnected
}

// Permet de supprimer un cookie lors de la déconnexion voulue d'un utilisateur
func GetDeconnected(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "CookieSession",
		MaxAge: -1}
	http.SetCookie(w, &c)
	Color(1, "[CONNEXION] : 🟢 Vous êtes bien déconnectés ")
	http.Redirect(w, r, "/connexion", http.StatusSeeOther)
}
