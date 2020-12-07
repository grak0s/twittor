package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*variable is used to connect the MongoDB database*/
var MongoCN = ConectarDB()

var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

/* ConnectDatabase is used to connect the MongoDB database*/
func ConectarDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa en la BD")
	return client
}

/* chequeo is used to connect the MongoDB database*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {

		return 0
	}
	return 1

}
