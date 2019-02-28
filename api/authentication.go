package api

import ("context"
		"fmt"
		"log"
		"time"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
		)

				//"github.com/mongodb/mongo-go-driver/bson"

func ConnectToDb() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://rob:test123@ds155665.mlab.com:55665/deployment-app"))

		// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
