package handlers

import (
	"net/http"

	BDD "../BDD"
)

// Fonction qui vérifie que l'utilisateur peut se connecter
func GetLogin(w http.ResponseWriter, r *http.Request, pseudo string, password string) (Errors, error) {
	var DataPageConnexion Errors
	DataPageConnexion.Pseudo = pseudo

	// On vérifie que l'utilisateur fait bien partie de la base de données
	correctPseudo, _, err := BDD.VerifyBDD(pseudo, "pseudo")
	if err != nil {
		return DataPageConnexion, err
	}
	if !correctPseudo {
		DataPageConnexion.Error = "Ce pseudonyme n'existe pas dans notre base de donnée"
		return DataPageConnexion, nil
	}

	// On vérifie que le mot de passe entré et le mot de passe dans la base de données correspondent
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
