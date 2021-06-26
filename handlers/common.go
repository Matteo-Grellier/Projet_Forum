package handlers

import BDD "../BDD"

type TopicDataUsed struct {
	Topic        BDD.Topic
	Posts        []BDD.Post
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
