package models

// movie data from api
type Movie_Api struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
}

// movie data from db
type Detail struct {
	Movie_ID int `json:"movie_id"`
	Title string `json:"title"`
	Overview string `json:"overview"`
	Views    int `json:"views"`
}

// comments from db
type Comment struct {
	Comment_ID int	`json:"comment_id"`
	Movie_ID int    `json:"movie_id"`
	User_ID  int    `json:"user_id"`
	Content  string `json:"content"`
}

// users from db
type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// movie to return
type Movie_Detail struct {
	Movie_ID int    `json:"movie_id"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Views    int    `json:"views"`
	Comments []Comment_View `json:"comments"`
}

// comment to return
type Comment_View struct {
	UserName  string `json:"user_name"`
	Content  string `json:"content"`
}
