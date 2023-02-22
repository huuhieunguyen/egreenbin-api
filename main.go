package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GDSC-UIT/egreenbin-api/component"
	"github.com/GDSC-UIT/egreenbin-api/handlers"
	middleware "github.com/GDSC-UIT/egreenbin-api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// const uri = "mongodb://admin:123123123@localhost:27017/?maxPoolSize=20&w=majority"
// const uri = "mongodb+srv://admin:123123123@gdsc.uytfb9v.mongodb.net/?retryWrites=true&w=majority"

func main() {
	// Set up the MongoDB connection
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	uriDb := os.Getenv("CONNECTION_STRING")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uriDb))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	router := gin.Default()
	db := client.Database("egreenbin")

	appContext := component.NewAppContext(db)

	router.Use(middleware.Recover(appContext))

	apiR := router.Group("/api")
	handlers.NewStudentHandler(apiR, appContext, db)

	log.Fatal(router.Run(":3000"))
}
