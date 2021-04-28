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

func FindLibrarys(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Library, error) {
	var librarys []models.Library
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Incapaz de encontrar el libro %+v", err)
	}
	err = cursor.All(ctx, &librarys)
	if err != nil {
		log.Printf("Incapaz de leer el cursor %+v", err)
	}
	return librarys, nil
}

func FindLybrary(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.Library, error) {
	var library models.Library
	res := collection.FindOne(ctx, bson.M{"_id": id})
	err := res.Decode(&library)
	if err != nil {
		return library, err
	}
	return library, nil
}

func ModifyLybrary(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.Library, error) {
	var library models.Library
	//find if the product exits, if err return 404
	filter := bson.M{"_id": id}
	res := collection.FindOne(ctx, filter)
	log.Println(res)
	log.Println(filter)
	if err := res.Decode(&library); err != nil {
		log.Printf("unable to decode to library :%v", err)
		return library, err
	}

	if err := json.NewDecoder(reqBody).Decode(&library); err != nil {
		log.Printf("unable to decode using reqbody : %v", err)
		return library, err
	}

	//update the product, if err return 500
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": library})
	if err != nil {
		log.Printf("Incapaz de actualizar el libro : %v", err)
		return library, err
	}
	return library, nil
}

func InsertLybrary(ctx context.Context, librarys []models.Library, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, library := range librarys {
		library.ID = ksuid.New().String()
		insertID, err := collection.InsertOne(ctx, library)
		if err != nil {
			log.Printf("Incapaz de insertar en la base de datos:%v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

func DeleteLybrary(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64, error) {
	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("Incapaz de eliminar un libro : %v", err)
		return 0, err
	}
	return res.DeletedCount, nil
}
