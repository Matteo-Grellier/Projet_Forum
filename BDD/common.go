package BDD

type User struct {
	Pseudo   string
	Mail     string
	Password string
}
type DataUsed struct {
	Users  []User
	Topics []Topic
}

type Topic struct {
	ID          int
	Title       string
	Content     string
	User_pseudo string
	Category_ID int
}

type Category struct {
	Name string
	Id   string
}
