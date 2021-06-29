package handlers

import BDD "../BDD"

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
type UserConnectedStruct struct {
	PseudoConnected string
	Connected       bool
}

var BDDerror error
