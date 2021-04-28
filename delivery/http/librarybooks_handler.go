package http

import (
	"context"
	"echo-mongo-project2/config/dbiface"
	librarybooksUcase "echo-mongo-project2/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewLybraryBooksHandler(e *echo.Echo, libocol dbiface.CollectionAPI) {
	h3 := Handler{Col: libocol}
	e.GET("/library/books", h3.GetLibraryBooksEndpoint)
}

func (h *Handler) GetLibraryBooksEndpoint(c echo.Context) error {
	librarys, err := librarybooksUcase.FindLbraryBooksAllData(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, librarys, "  ") //c.JSON(http.StatusOK, librarys)
}
