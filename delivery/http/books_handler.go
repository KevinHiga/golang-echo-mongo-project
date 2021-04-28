package http

import (
	"context"
	"echo-mongo-project2/config/dbiface"
	models "echo-mongo-project2/models"
	booksUcase "echo-mongo-project2/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewBooksHandler(e *echo.Echo, bocol dbiface.CollectionAPI) {
	h := Handler{Col: bocol}
	e.GET("/books/:id", h.GetBookEndpoint)
	e.DELETE("/books/:id", h.DeleteBookEndpoint)
	e.PUT("/books/:id", h.UpdateBookEndpoint)
	e.POST("/books", h.CreateBooksEndpoint)
	e.GET("/books", h.GetBooksEndpoint)
}

func (h *Handler) GetBooksEndpoint(c echo.Context) error {
	books, err := booksUcase.FindBooksAllData(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, books, "  ")
}

func (h *Handler) GetBookEndpoint(c echo.Context) error {
	book, err := booksUcase.FindBooksOneData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, book, "  ")
}

func (h *Handler) UpdateBookEndpoint(c echo.Context) error {
	product, err := booksUcase.UpdateBooksData(context.Background(), c.Param("id"), c.Request().Body, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, product)
}

func (h *Handler) CreateBooksEndpoint(c echo.Context) error {
	var books []models.Book
	fmt.Println(books)
	if err := c.Bind(&books); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	IDs, err := booksUcase.CreateBooksData(context.Background(), books, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}

func (h *Handler) DeleteBookEndpoint(c echo.Context) error {
	delCount, err := booksUcase.DeleteBooksData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, delCount)
}
