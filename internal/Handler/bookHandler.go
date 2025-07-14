package handler

import (
	"github.com/gofiber/fiber/v2"
	entities "github.com/twichai/book-store-full-golang/internal/domain"
)

type BookHandler struct {
	usecase entities.BookUsecase
}

// Create implements entities.BookHandler.
func (b *BookHandler) Create(c *fiber.Ctx) error {
	newBook := new(entities.Book)
	if err := c.BodyParser(newBook); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	if err := b.usecase.Create(*newBook); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusCreated)
}

// Delete implements entities.BookHandler.
func (b *BookHandler) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAll godoc
// @Summary      Get all books
// @Description  Returns a list of books
// @Tags         book
// @Accept       json
// @Produce      json
// @Success      200  {array}   entities.Book
// @Failure      500  {object}  map[string]string
// @Router       /books [get]
func (b *BookHandler) GetAll(c *fiber.Ctx) error {
	books, err := b.usecase.GetAll()
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)
}

// GetById implements entities.BookHandler.
func (b *BookHandler) GetById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements entities.BookHandler.
func (b *BookHandler) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewBookHandler(usecase entities.BookUsecase) entities.BookHandler {
	return &BookHandler{usecase: usecase}
}
