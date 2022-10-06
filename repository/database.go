package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	mongoCtx context.Context
	//client   *mongo.Client
	database *mongo.Database
}

func (db *DB) FindAllByLimit(collectionName string, document any) error {
	collection := db.database.Collection(collectionName)
	opts := options.Find().SetLimit(100)
	cur, err := collection.Find(db.mongoCtx, bson.D{}, opts)
	if err != nil {
		return err
	}
	if cur != nil {
		err = cur.All(db.mongoCtx, document)
	}
	return err
}
