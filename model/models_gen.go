// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
	Author    *User  `json:"author"`
}

type User struct {
	ID    string  `json:"id"`
	Email *string `json:"email"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}
