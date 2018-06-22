package models

// Book repesents a book object.
type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	AuthorID string `json:"authorId"`
}

// Author represents an author object
type Author struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Books []Book `json:"books"`
}
