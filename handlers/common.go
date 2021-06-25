package handlers

import BDD "../BDD"

type TopicDataUsed struct {
	Topics       []BDD.Topic
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
