package usecase

import (
	"errors"
	entities "github.com/twichai/book-store-full-golang/internal/domain"
)

type BookUsecase struct {
	repo entities.BookRepo
}

// Delete implements entities.BookUsecase.
func (b *BookUsecase) Delete(id uint) error {
	return b.repo.Delete(id)
}

// GetAll implements entities.BookUsecase.
func (b *BookUsecase) GetAll() ([]*entities.Book, error) {
	return b.repo.GetAll()
}

// GetById implements entities.BookUsecase.
func (b *BookUsecase) GetById(id uint) (*entities.Book, error) {
	return b.repo.GetById(id)
}

// Update implements entities.BookUsecase.
func (b *BookUsecase) Update(book *entities.Book, bookId uint) (*entities.Book, error) {
	book.ID = bookId
	if book.Price < 0 {
		return nil, errors.New("Price can't be negative")
	}
	if book.Stock < 0 {
		return nil, errors.New("Stock can't be negative")
	}
	return b.repo.Update(book, bookId)
}

// Create implements entities.BookUsecase.
func (b *BookUsecase) Create(book *entities.Book) error {
	return b.repo.Create(book)
}

func NewBookUsecase(repo entities.BookRepo) entities.BookUsecase {
	return &BookUsecase{repo: repo}
}
