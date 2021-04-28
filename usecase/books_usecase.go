package usecase

import (
	"context"
	"io"
	"echo-mongo-project2/config/dbiface"
	models "echo-mongo-project2/models"
	booksRepo "echo-mongo-project2/repository/mongodb"
)

func FindBooksAllData(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Book, error) {
	return booksRepo.FindBooks(ctx, collection)
}

func FindBooksOneData(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.Book, error) {
	return booksRepo.FindBook(ctx, id, collection)
}

func UpdateBooksData(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.Book, error) {
	return booksRepo.ModifyBook(ctx, id, reqBody, collection)
}

func CreateBooksData(ctx context.Context, books []models.Book, collection dbiface.CollectionAPI) ([]interface{}, error) {
	return booksRepo.InsertBook(ctx, books, collection)
}

func DeleteBooksData(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64, error) {
	return booksRepo.DeleteBook(ctx, id, collection)
}
