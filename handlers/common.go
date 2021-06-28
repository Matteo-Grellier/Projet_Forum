package handlers

import BDD "../BDD"

type TopicDataUsed struct {
	Topic BDD.Topic
	Posts []BDD.Post
	Like  BDD.Likes
	/* Likes        []BDD.Like */
	ErrorMessage string
}
type DataPageCategory struct {
	Topics     []BDD.Topic
	Category   string
	CategoryID int
	Error      string
}
type UserConnectedStruct struct {
	PseudoConnected string
	Connected       bool
}
