package handlers

import (
	"log"
	"net/http"
	"regexp"
	"text/template"

	BDD "../BDD"
	"golang.org/x/crypto/bcrypt"
)

var ErrorMessage string

type Errors struct {
	Error  string
	Pseudo string
	Mail   string
}

func HashPassword(password string) string {
	var passByte = []byte(password)

	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.MinCost)
	if err != nil {
		Color(4, "[HASH_INFO] : 🔻 Error function 'HashPassword' : ")
		log.Fatal(err)
	}

	return string(hash)
}

func GetRegister(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/inscription.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	err2 := r.ParseForm()
	if err2 != nil {
		Color(4, "[PARSE_FORM_INFO] : 🔻 Error function 'GetRegister' : ")
		log.Fatal(err2)
	}

	pseudo := r.FormValue("Pseudo")
	email := r.FormValue("Email")
	password := r.FormValue("Password")
	confirmPwd := r.FormValue("ConfirmPassword")

	var allDataRegister = []string{pseudo, email, password, confirmPwd}

	var newPass = HashPassword(password)

	if verifyInput(allDataRegister) && isValidEmail(email) && verifyBDD(email, "mail") && verifyBDD(pseudo, "pseudo") && verifMdp(password) && sameMdp(password, confirmPwd) {
		db := BDD.OpenDataBase()
		createNew, err := db.Prepare("INSERT INTO user (pseudo, mail, password) VALUES (?, ?, ?)")
		if err != nil {
			Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
			log.Fatal(err)
		}
		_, err = createNew.Exec(pseudo, email, newPass)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
	} else {
		ErrorsInscriptions := Errors{
			Error:  ErrorMessage,
			Pseudo: pseudo,
			Mail:   email,
		}
		ErrorMessage = ""
		t.Execute(w, ErrorsInscriptions)
	}
}

func isValidEmail(email string) bool {
	// Vérifie si l'email entré est valide
	var re = regexp.MustCompile(`(?mi)[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}`)
	if re.MatchString(email) {
		return true
	}
	ErrorMessage = "Email non valide."

	return false
}

func verifyBDD(element string, column string) bool {
	// Vérifie si l'élément est déjà dans la base de donnée ou pas.
	db := BDD.OpenDataBase()
	var oneElement string
	allElements, err := db.Query("SELECT " + column + " FROM user")
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatalf("%s", err)
	}
	for allElements.Next() {
		allElements.Scan(&oneElement)
		if oneElement == element {
			ErrorMessage = column + " déjà dans la base de données."
			return false
		}
	}
	return true
}

func verifyInput(label []string) bool {

	for index := 0; index < len(label); index++ {
		if len(label[index]) == 0 {
			ErrorMessage = "Veuillez renseigner tous les champs"
			return false
		}
	}
	return true

}

func verifMdp(mdp string) bool {
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
			ErrorMessage = "Mot de passe non valide (au moins 1 majuscule)."
			return false
		}
		if min >= 1 {
		} else {
			ErrorMessage = "Mot de passe non valide (au moins 1 minuscule)."
			return false
		}
		if chiffre >= 1 {
		} else {
			ErrorMessage = "Mot de passe non valide (au moins 1 chiffre)."
			return false
		}
	} else {
		ErrorMessage = "Mot de passe non valide (minimum 7 caractères)."
		return false
	}
	return true
}

func sameMdp(firstpwd string, secondpwd string) bool {
	if firstpwd != secondpwd {
		ErrorMessage = "les mots de passe ne correspondent pas."
		return false
	}
	return true
}
