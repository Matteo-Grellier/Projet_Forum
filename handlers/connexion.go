package handlers

import (
	"net/http"

	BDD "../BDD"
)

func GetLogin(w http.ResponseWriter, r *http.Request, pseudo string, password string) Errors {
	var DataPageConnexion Errors
	DataPageConnexion.Pseudo = pseudo

	correctPseudo, _ := BDD.VerifyBDD(pseudo, "pseudo")
	if !correctPseudo {
		DataPageConnexion.Error = "Ce pseudonyme n'existe pas dans notre base de donnée"
		return DataPageConnexion
	}

	correctPassword := BDD.VerifyPassword(password, pseudo)
	if !correctPassword {
		DataPageConnexion.Error = "Mot de passe incorrect !"
		return DataPageConnexion
	}
	return DataPageConnexion

}
