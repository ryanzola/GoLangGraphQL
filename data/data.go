package data

import (
	"fmt"

	"github.com/ryanzola/GoLangGraphQL/server/models"
)

// Books object model
var Books = []*models.Book{}

// Authors object model
var Authors = []*models.Author{}

func init() {
	book1 := &models.Book{
		ID:       "5b25981017052433f920f119",
		Name:     "Catcher in the Rye",
		Genre:    "Fiction",
		AuthorID: "5b25935f44b43333c8da8707",
	}
	book2 := &models.Book{
		ID:       "5b25cbd549a7cc37f48b91df",
		Name:     "Harry Potter and the Sorcerer's Stone",
		Genre:    "Fantasy",
		AuthorID: "5b25cad549a7cc37f48b91de",
	}
	author1 := &models.Author{
		ID:   "5b25935f44b43333c8da8707",
		Name: "J.D. Salinger",
		Age:  94,
	}
	author2 := &models.Author{
		ID:   "5b25cad549a7cc37f48b91de",
		Name: "J.K. Rowling",
		Age:  50,
	}

	Books = []*models.Book{
		book1,
		book2,
	}

	Authors = []*models.Author{
		author1,
		author2,
	}
}

//GetBook retrieves a book for a given user ID
func GetBook(id string) (*models.Book, error) {
	for _, book := range Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("Book (id: %v) was not found", id)
}

// GetBooks gets all the books
func GetBooks() ([]*models.Book, error) {
	if len(Books) > 0 {
		return Books, nil
	}
	return nil, fmt.Errorf("There doesn't appear to be any books")
}

//GetAuthor retrieves an author for a given user ID
func GetAuthor(id string) (*models.Author, error) {
	for _, author := range Authors {
		if author.ID == id {
			return author, nil
		}
	}
	return nil, fmt.Errorf("Author (id: %v) was not found", id)
}

// GetAuthors returns all of the authors
func GetAuthors() ([]*models.Author, error) {
	if len(Authors) > 0 {
		return Authors, nil
	}
	return nil, fmt.Errorf("There doesn't appear to be any authors")
}

// GetAllBooksByAuthor returns an array of books by an author
func GetAllBooksByAuthor(authorID string) ([]*models.Book, error) {
	var bookSlice []*models.Book
	for _, book := range Books {
		if book.AuthorID == authorID {
			bookSlice = append(bookSlice, book)
		}
	}
	return bookSlice, nil
}
