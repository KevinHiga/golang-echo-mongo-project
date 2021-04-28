package usecase

import (
	"context"
	"io"
	"echo-mongo-project2/config/dbiface"
	models "echo-mongo-project2/models"
	libraryRepo "echo-mongo-project2/repository/mongodb"
)

func FindLibraryAllData(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Library, error) {
	return libraryRepo.FindLibrarys(ctx, collection)
}

func FindLibraryOneData(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.Library, error) {
	return libraryRepo.FindLybrary(ctx, id, collection)
}

func UpdateLibraryData(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.Library, error) {
	return libraryRepo.ModifyLybrary(ctx, id, reqBody, collection)
}

func CreateLibraryData(ctx context.Context, librarys []models.Library, collection dbiface.CollectionAPI) ([]interface{}, error) {
	return libraryRepo.InsertLybrary(ctx, librarys, collection)
}

func DeleteLibraryData(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64, error) {
	return libraryRepo.DeleteLybrary(ctx, id, collection)
}