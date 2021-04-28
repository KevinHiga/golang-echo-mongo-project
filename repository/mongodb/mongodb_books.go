package mongodb

import (
	"context"
	"echo-mongo-project2/config/dbiface"
	models "echo-mongo-project2/models"
	
	"encoding/json"
	"io"
	"log"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
)

func FindBooks(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Book, error) {
	var books []models.Book
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Incapaz de encontrar el libro %+v", err)
	}
	err = cursor.All(ctx, &books)
	if err != nil {
		log.Printf("Incapaz de leer el cursor %+v", err)
	}
	return books, nil
}

func FindBook(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.Book, error) {
	var book models.Book
	res := collection.FindOne(ctx, bson.M{"_id": id})
	err := res.Decode(&book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func ModifyBook(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.Book, error) {
	var book models.Book
	//find if the product exits, if err return 404
	filter := bson.M{"_id": id}
	res := collection.FindOne(ctx, filter)
	log.Println(res)
	log.Println(filter)
	if err := res.Decode(&book); err != nil {
		log.Printf("unable to decode to book :%v", err)
		return book, err
	}

	if err := json.NewDecoder(reqBody).Decode(&book); err != nil {
		log.Printf("unable to decode using reqbody : %v", err)
		return book, err
	}

	//update the product, if err return 500
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": book})
	if err != nil {
		log.Printf("Incapaz de actualizar el libro : %v", err)
		return book, err
	}
	return book, nil
}

func InsertBook(ctx context.Context, books []models.Book, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, book := range books {
		book.ID = ksuid.New().String()
		insertID, err := collection.InsertOne(ctx, book)
		if err != nil {
			log.Printf("Incapaz de insertar en la base de datos:%v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

func DeleteBook(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64, error) {
	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("Incapaz de eliminar un libro : %v", err)
		return 0, err
	}
	return res.DeletedCount, nil
}
