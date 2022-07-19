package models

//Book --> Cada livro terá um ID, titulo e um autor.
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
