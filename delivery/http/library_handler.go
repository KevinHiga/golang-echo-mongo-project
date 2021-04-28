package http

import (
	"context"
	"echo-mongo-project2/config/dbiface"
	models "echo-mongo-project2/models"
	libraryUcase "echo-mongo-project2/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewLibraryHandler(e *echo.Echo, licol dbiface.CollectionAPI) {
	h2 := Handler{Col: licol}
	e.GET("/library/:id", h2.GetLybraryEndpoint)
	e.DELETE("/library/:id", h2.DeleteLybraryEndpoint)
	e.PUT("/library/:id", h2.UpdateLybraryEndpoint)
	e.POST("/library", h2.CreateLybrarysEndpoint)
	e.GET("/library", h2.GetLybrarysEndpoint)
}

func (h *Handler) GetLybrarysEndpoint(c echo.Context) error {
	librarys, err := libraryUcase.FindLibraryAllData(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, librarys, "  ")
}

func (h *Handler) GetLybraryEndpoint(c echo.Context) error {
	library, err := libraryUcase.FindLibraryOneData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, library, "  ")
}

func (h *Handler) UpdateLybraryEndpoint(c echo.Context) error {
	library, err := libraryUcase.UpdateLibraryData(context.Background(), c.Param("id"), c.Request().Body, h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, library, "  ")
}

func (h *Handler) CreateLybrarysEndpoint(c echo.Context) error {
	var librarys []models.Library
	fmt.Println(librarys)
	if err := c.Bind(&librarys); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	IDs, err := libraryUcase.CreateLibraryData(context.Background(), librarys, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}

func (h *Handler) DeleteLybraryEndpoint(c echo.Context) error {
	delCount, err := libraryUcase.DeleteLibraryData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, delCount)
}
