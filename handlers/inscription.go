package handlers

import (
	"log"
	"regexp"

	BDD "../BDD"
	"golang.org/x/crypto/bcrypt"
)

// Fonction qui permet de hasher un mot de passe
func HashPassword(password string) string {
	var passByte = []byte(password)

	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.MinCost)
	if err != nil {
		Color(4, "[HASH_INFO] : 🔻 Error function 'HashPassword' : ")
		log.Fatal(err)
	}

	return string(hash)
}

// Enregistre l'utilisateur qui s'inscrit
func GetRegister(pseudo string, email string, password string, confirmPwd string) (Errors, error) {
	var allDataRegister = []string{pseudo, email, password, confirmPwd}
	var newPass = HashPassword(password)
	var DataPageInscription Errors
	DataPageInscription.Mail = email
	DataPageInscription.Pseudo = pseudo

	// On vérifie le mail dans la BDD
	correctEmail, errMail, err := BDD.VerifyBDD(email, "mail")
	if err != nil {
		return DataPageInscription, err
	}
	// On vérifie le pseudo dans la BDD
	correctPseudo, errPseudo, err := BDD.VerifyBDD(pseudo, "pseudo")
	if err != nil {
		return DataPageInscription, err
	}
	correctPassword, errPassword := verifMdp(password)

	// On vérifie que tous les champs sont bien remplis
	if !verifyInput(allDataRegister) {
		DataPageInscription.Error = "Veuillez rentrer tous les champs."
		return DataPageInscription, nil

		// On vérifie que le pseudo n'est pas déjà dans la BDD
	} else if correctPseudo {
		DataPageInscription.Error = errPseudo
		return DataPageInscription, nil

		//On regarde si le mail est sous la bonne forme
	} else if !isValidEmail(email) {
		DataPageInscription.Error = "Mail non valide."
		return DataPageInscription, nil

		// On vérifie que l'email n'est pas déjà dans la BDD
	} else if correctEmail {
		DataPageInscription.Error = errMail
		return DataPageInscription, nil

		// On regarde si les 2 mots de passe correspondent
	} else if !sameMdp(password, confirmPwd) {
		DataPageInscription.Error = "Les mots de passe ne correspondent pas."
		return DataPageInscription, nil

		// On regarde si le mot de passe correspond aux critères demandés
	} else if !correctPassword {
		DataPageInscription.Error = errPassword
		return DataPageInscription, nil
	}

	// Lorsque toutes les conditions sont remplies, on ajoute à la BDD.
	BDDerror = BDD.AddUser(pseudo, email, newPass)
	return DataPageInscription, BDDerror
}

// Vérifie si l'email entré est valide
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`(?mi)[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}`)
	return re.MatchString(email)
}

// Vérifie que tous les inputs sont bien remplis
func verifyInput(label []string) bool {

	for index := 0; index < len(label); index++ {
		if len(label[index]) == 0 {
			return false
		}
	}
	return true
}

// Vérifie tous les critères pour le mot de passe (7 caractères, min 1 Minuscule, min 1 Majuscule, min 1 chiffre)
func verifMdp(mdp string) (bool, string) {
	var maj int = 0
	var min int = 0
	var chiffre int = 0

	for index := 0; index < len(mdp); index++ {
		if mdp[index] >= 65 && mdp[index] <= 90 {
			maj++
		}
		if mdp[index] >= 97 && mdp[index] <= 122 {
			min++
		}
		if mdp[index] >= 48 && mdp[index] <= 57 {
			chiffre++
		}
	}
	if len(mdp) >= 7 {
		if maj >= 1 {
		} else {
			return false, "Mot de passe non valide (au moins 1 majuscule)."
		}
		if min >= 1 {
		} else {
			return false, "Mot de passe non valide (au moins 1 minuscule)."
		}
		if chiffre >= 1 {
		} else {
			return false, "Mot de passe non valide (au moins 1 chiffre)."
		}
	} else {
		return false, "Mot de passe non valide (minimum 7 caractères)."
	}
	return true, ""
}

// Vérification si les mots de passes correspondent dans l'inscription
func sameMdp(firstpwd string, secondpwd string) bool {
	if firstpwd != secondpwd {
		return false
	}
	return true
}
