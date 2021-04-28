package main

import (
	"context"
	config "echo-mongo-project2/config/db"
	_HttpDeliveryMiddleware "echo-mongo-project2/config/middleware"
	_HttpDelivery "echo-mongo-project2/delivery/http"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c       *mongo.Client
	db      *mongo.Database
	bocol   *mongo.Collection
	licol   *mongo.Collection
	libocol *mongo.Collection
	cfg     config.Properties
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("La configuracion no puede ser leida: %v", err)
	}
	connectURI := fmt.Sprintf("%s", cfg.DBMongo)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Incapaz de conectarse a la base de datos: %v", err)
	}
	db = c.Database(cfg.DBName)
	bocol = db.Collection(cfg.BooksCollection)
	licol = db.Collection(cfg.LibraryCollection)
	libocol = db.Collection(cfg.LibBooCollection)
}
func main() {
	e := echo.New()
	e.Use(_HttpDeliveryMiddleware.CORS)
	_HttpDelivery.NewBooksHandler(e, bocol)
	_HttpDelivery.NewLibraryHandler(e, licol)
	_HttpDelivery.NewLybraryBooksHandler(e, libocol)
	e.Logger.Infof("Escuchando en %s:%s", cfg.Host, cfg.Port)
	fmt.Println("Escuchando")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
