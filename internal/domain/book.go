package entities

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	title          string    `json:"title,omitempty"`
	author         string    `json:"author,omitempty"`
	published_date time.Time `json:"published_date,omitempty"`
}

type BookRepo interface {
	Create(book Book) error
	GetById(id uint) (Book, error)
	GetAll() ([]Book, error)
	Update(book Book) (Book, error)
	Delete(id uint) error
}

type BookUsecase interface {
	Create(book Book) error
	GetById(id uint) (Book, error)
	GetAll() ([]Book, error)
	Update(book Book, bookId uint) (Book, error)
	Delete(id uint) error
}

type BookHandler interface {
	Create(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
