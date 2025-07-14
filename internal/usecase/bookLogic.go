package usecase

import entities "github.com/twichai/book-store-full-golang/internal/domain"

type BookUsecase struct {
	repo entities.BookRepo
}

// Delete implements entities.BookUsecase.
func (b *BookUsecase) Delete(id uint) error {
	panic("unimplemented")
}

// GetAll implements entities.BookUsecase.
func (b *BookUsecase) GetAll() ([]entities.Book, error) {
	return b.repo.GetAll()
}

// GetById implements entities.BookUsecase.
func (b *BookUsecase) GetById(id uint) (entities.Book, error) {
	panic("unimplemented")
}

// Update implements entities.BookUsecase.
func (b *BookUsecase) Update(book entities.Book, bookId uint) (entities.Book, error) {
	panic("unimplemented")
}

// Create implements entities.BookUsecase.
func (b *BookUsecase) Create(book entities.Book) error {
	return b.repo.Create(book)
}

func NewBookUsecase(repo entities.BookRepo) entities.BookUsecase {
	return &BookUsecase{repo: repo}
}
