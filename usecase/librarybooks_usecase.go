package usecase

import (
	"context"
	"echo-mongo-project2/config/dbiface"
	models "echo-mongo-project2/models"
	librarybooksRepo "echo-mongo-project2/repository/mongodb"
)

func FindLbraryBooksAllData(ctx context.Context, collection dbiface.CollectionAPI) ([]models.LibraryBook, error) {
	return librarybooksRepo.FindLibraryBooks(ctx, collection)
}
