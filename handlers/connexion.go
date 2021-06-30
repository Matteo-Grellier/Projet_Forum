package handlers

import (
	"net/http"

	BDD "../BDD"
)

// Fonction qui vérifie si l'utilisateur est dans la base de donnée
func GetLogin(w http.ResponseWriter, r *http.Request, pseudo string, password string) (Errors, error) {
	var DataPageConnexion Errors
	DataPageConnexion.Pseudo = pseudo

	correctPseudo, _, err := BDD.VerifyBDD(pseudo, "pseudo")
	if err != nil {
		return DataPageConnexion, err
	}
	if !correctPseudo {
		DataPageConnexion.Error = "Ce pseudonyme n'existe pas dans notre base de donnée"
		return DataPageConnexion, nil
	}

	correctPassword, err := BDD.VerifyPassword(password, pseudo)
	if err != nil {
		return DataPageConnexion, err
	}
	if !correctPassword {
		DataPageConnexion.Error = "Mot de passe incorrect !"
		return DataPageConnexion, nil
	}
	return DataPageConnexion, nil

}
