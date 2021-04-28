package mongodb

import (
	"context"
	"echo-mongo-project2/config/dbiface"
	models "echo-mongo-project2/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindLibraryBooks(ctx context.Context, collection dbiface.CollectionAPI) ([]models.LibraryBook, error) {
	var librarys []models.LibraryBook
	//cursor, err := collection.Find(ctx, bson.M{})
	//libname := "DISBOOK"
	libautor := "Fiódor Dostoievski"
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "books"}, {"localField", "_id"}, {"foreignField", "libraryId"}, {"as", "library"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$library"}}}}
	matchStage := bson.D{{"$match", bson.D{{"library.autor", libautor}}}}
	projectStage := bson.D{{"$project", bson.D{{"_id", 0}, {"'library._id'", 0}}}}
	sortStage := bson.D{{"$sort", bson.D{{"name", 1}, {"'library.year'", -1}}}}
	limitStage := bson.D{{"$limit", 5}}
	fmt.Println("Escuchando")
	showInfoCursor, err := collection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage, matchStage, projectStage, sortStage, limitStage})
	if err != nil {
		panic(err)
	}
	if err = showInfoCursor.All(ctx, &librarys); err != nil {
		panic(err)
	}
	return librarys, nil
}
