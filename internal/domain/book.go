package entities

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model    `swaggerignore:"true"`
	Title         string    `json:"title,omitempty"`
	Price         float32   `json:"price,omitempty"`
	Stock         int       `json:"stock,omitempty"`
	PublishedDate time.Time `json:"published_date"`
	PublisherID   uint      `json:"publisher_id,omitempty"`
	Publicsher    Publisher `json:"publicsher" gorm:"foreignKey:PublisherID;references:ID" swaggerignore:"true"`
}

type BookRepo interface {
	Create(book Book) error
	GetById(id uint) (*Book, error)
	GetAll() ([]*Book, error)
	Update(book Book) (*Book, error)
	Delete(id uint) error
}

type BookUsecase interface {
	Create(book Book) error
	GetById(id uint) (*Book, error)
	GetAll() ([]*Book, error)
	Update(book Book, bookId uint) (*Book, error)
	Delete(id uint) error
}

type BookHandler interface {
	Create(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
