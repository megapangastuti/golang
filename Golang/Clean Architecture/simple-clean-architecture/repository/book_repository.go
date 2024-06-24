// ToDo :
// 1. Mendeklarasikan nama package pada file book_repository.go
// 2. Mendeklarasikan struct bernama bookRepository
// 3. Mendeklarasikan interface bernama BookRepository
// 4. Membuat method dari interface yang telat dibuat
// 5. Memuat constructor bernama NewBookRepository

package repository

import (
	"database/sql"
	"simple-clean-architecture/model"
)

type bookRepository struct {
	db *sql.DB
}

type BookRepository interface {
	CreateNewBook(book model.Book) (model.Book, error)
	GetAllBook() ([]model.Book, error)
	GetBookById(id int) (model.Book, error)
	UpdateBookById(book model.Book) (model.Book, error)
	DeleteBookById(id int) error
}

func (b *bookRepository) CreateNewBook(book model.Book) (model.Book, error) {
	var bookId int

	err := b.db.QueryRow("INSERT INTO mst_book (title, author, release_year, pages) VALUES ($1, $2, $3, $4) RETURNING id", book.Title, book.Author, book.ReleaseYear, book.Pages).Scan(&bookId)

	if err != nil {
		return model.Book{}, err
	}
	book.Id = bookId

	return book, nil
}

func (b *bookRepository) GetAllBook() ([]model.Book, error) {
	var books []model.Book

	rows, err := b.db.Query("SELECT id, title, author, release_year, pages FROM mst_book")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book model.Book

		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.ReleaseYear, &book.Pages)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (b *bookRepository) GetBookById(id int) (model.Book, error) {
	var book model.Book

	err := b.db.QueryRow("SELECT id, title, author, release_year, pages FROM mst_book WHERE id = $1", id).Scan(&book.Id, &book.Title, &book.Author, &book.ReleaseYear, &book.Pages)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (b *bookRepository) UpdateBookById(book model.Book) (model.Book, error) {
	_, err := b.db.Exec("UPDATE mst_book SET title = $2, author = $3, release_year = $4, pages = $5 WHERE id = $1", book.Id, book.Title, book.Author, book.ReleaseYear, book.Pages)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (b *bookRepository) DeleteBookById(id int) error {
	_, err := b.db.Exec("DELETE FROM mst_book WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}
