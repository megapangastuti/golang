package usecase

import (
	"fmt"

	"simple-clean-architecture/model"
	"simple-clean-architecture/repository"
)

type BookUseCase interface {
	CreateNewBook(book model.Book) (model.Book, error)
	GetAllBook() ([]model.Book, error)
	GetBookById(id int) (model.Book, error)
	UpdateBookById(book model.Book) (model.Book, error)
	DeleteBookById(id int) error
}

type bookUseCase struct {
	repo repository.BookRepository
}

func (b *bookUseCase) CreateNewBook(book model.Book) (model.Book, error) {
	return b.repo.CreateNewBook(book)
}

func (b *bookUseCase) GetAllBook() ([]model.Book, error) {
	return b.repo.GetAllBook()
}

func (b *bookUseCase) GetBookById(id int) (model.Book, error) {
	return b.repo.GetBookById(id)
}

func (b *bookUseCase) UpdateBookById(book model.Book) (model.Book, error) {
	_, err := b.repo.GetBookById(book.Id)
	if err != nil {
		return model.Book{}, fmt.Errorf("book with ID %d not found", book.Id)
	}
	return b.repo.UpdateBookById(book)
}

func (b *bookUseCase) DeleteBookById(id int) error {
	_, err := b.repo.GetBookById(id)
	if err != nil {
		return fmt.Errorf("book with ID %d not found", id)
	}
	return b.repo.DeleteBookById(id)
}

func NewBookUseCase(repo repository.BookRepository) BookUseCase {
	return &bookUseCase{repo: repo}
}
