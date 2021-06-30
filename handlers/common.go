package handlers

import BDD "../BDD"

// Toutes les structures de données des pages utilisées
type DataPageCategories struct {
	Categories []BDD.Category
	User       UserConnectedStruct
}

type TopicDataUsed struct {
	Topic         BDD.Topic
	Posts         []BDD.Post
	ErrorMessage  string
	UserConnected UserConnectedStruct
}
type DataPageCategory struct {
	Topics        []BDD.Topic
	Category      string
	CategoryID    int
	Error         string
	UserConnected UserConnectedStruct
}

type DataPageHome struct {
	UserConnected UserConnectedStruct
	Posts         []BDD.Post
}

type DataPageLikes struct {
	UserConnected UserConnectedStruct
	Posts         []BDD.Post
}
type UserConnectedStruct struct {
	PseudoConnected string
	Connected       bool
}

// Stocke l'erreur de BDD lors d'une requête pour la renvoyer au serveur.
var BDDerror error

// Structure de données pour les formulaires d'inscription et de connexion
// Permet de laisser affiché le pseudo et le mail s'il y a eu une erreur lors de l'inscription/connexion
type Errors struct {
	Error  string
	Pseudo string
	Mail   string
}

// Stocke le message d'erreur qui s'affichera sur la page
var ErrorMessage string
