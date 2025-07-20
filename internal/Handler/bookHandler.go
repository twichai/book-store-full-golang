package handler

import (
	"github.com/gofiber/fiber/v2"
	entities "github.com/twichai/book-store-full-golang/internal/domain"
)

type BookHandler struct {
	usecase entities.BookUsecase
}

// @Summary      Create a new book
// @Description  Creates a new book and returns the created book object
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book  body      entities.Book  true  "Book Data"
// @Success      201   {object}  entities.Book
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /books [post]
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

// Delete A Book
// @Summary      Delete a book
// @Description  Deletes a book by ID and returns status
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /books/{id} [delete]
func (b *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}
	if err := b.usecase.Delete(uint(id)); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"error": err})
	}
	return c.SendStatus(fiber.StatusOK)

}

// Get All Books
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

// Get A Book
// @Summary      Get a books
// @Description  Returns a books
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  entities.Book
// @Failure      500  {object}  map[string]string
// @Router       /books/{id} [get]
func (b *BookHandler) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}
	book, err := b.usecase.GetById(uint(id))
	if err != nil {
		c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(book)
}

// @Summary      Update a book
// @Description  Update a book by ID and return the updated book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id    path      int                 true  "Book ID"
// @Param        book  body      entities.Book   true  "Book Data"
// @Success      200   {object}  entities.Book
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /books/{id} [put]
func (b *BookHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	book := new(entities.Book)
	if err = c.BodyParser(&book); err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	book, err = b.usecase.Update(book, uint(id))
	if err != nil {
		c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"error": err})
	}
	return c.JSON(book)
}

func NewBookHandler(usecase entities.BookUsecase) entities.BookHandler {
	return &BookHandler{usecase: usecase}
}
