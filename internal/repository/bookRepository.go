package repository

import (
	entities "github.com/twichai/book-store-full-golang/internal/domain"
	"gorm.io/gorm"
)

type GormBookOrderRepository struct {
	db *gorm.DB
}

// Delete implements entities.BookRepo.
func (g *GormBookOrderRepository) Delete(id uint) error {
	panic("unimplemented")
}

// GetAll implements entities.BookRepo.
func (g *GormBookOrderRepository) GetAll() ([]entities.Book, error) {
	books := new([]entities.Book)
	if result := g.db.Find(&books); result.Error != nil {
		return nil, result.Error
	}
	return *books, nil
}

// GetById implements entities.BookRepo.
func (g *GormBookOrderRepository) GetById(id uint) (entities.Book, error) {
	panic("unimplemented")
}

// Update implements entities.BookRepo.
func (g *GormBookOrderRepository) Update(book entities.Book) (entities.Book, error) {
	panic("unimplemented")
}

// Create implements entities.BookRepo.
func (g GormBookOrderRepository) Create(book entities.Book) error {
	result := g.db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewBookRepository(db *gorm.DB) entities.BookRepo {
	return &GormBookOrderRepository{db: db}
}
