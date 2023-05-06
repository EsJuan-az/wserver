package wserver

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct{
	client *mongo.Client
	collection *mongo.Collection
}
func NewCollection(mongo_cnn, db, name string)(*mongo.Collection, error){
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(mongo_cnn)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil{
		return &mongo.Collection{}, err
	}
	return client.Database(db).Collection(name), err
}
