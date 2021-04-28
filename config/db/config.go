package config

type Properties struct{
	Port string `env:"MY_APP_PORT" env-default:"3030"`
	Host string `env:"HOST" env-default:"localhost"`
	DBMongo string `env:"HOST" env-default:"mongodb+srv://platzi-admin:eRTM4Ly38IzPMmCi@curso-platzi.js3sy.mongodb.net/echo-mongo-project?retryWrites=true&w=majority"`
	DBHost string `env:"DB_HOST" env-default:"localhost"`
	DBPort string `env:"DB_PORT" env-default:"27017"`
	DBName string `env:"DB_NAME" env-default:"echo-mongo-project"`
	BooksCollection string `env:"COLLECTION_NAME" env-default:"books"`
	LibraryCollection string `env:"COLLECTION_NAME" env-default:"library"`
	LibBooCollection string `env:"COLLECTION_NAME" env-default:"library.books"`
}