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

type Category struct {
	Name string
	Id   string
}

type Topic struct {
	ID            int
	Title         string
	Content       string
	Like          int
	User_pseudo   string
	Category_ID   int
	Category_name string
}

type Post struct {
	ID             int
	Content        string
	User_pseudo    string
	Topic_ID       string
	Comments       []Comment
	NumberComments int
	NumberLikes    int
}

type Comment struct {
	ID          int
	User_pseudo string
	Content     string
	Post_ID     int
}

type Likes struct {
	Status      int
	User_Pseudo string
	ID          int
}
