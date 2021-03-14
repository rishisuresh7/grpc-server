package driver

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database   = "grpc"
	collection = "books"
)

type MongoDriver interface {
	FindOne(c context.Context, query bson.M, opts *options.FindOneOptions) (*mongo.SingleResult, error)
	DeleteOne(c context.Context, query bson.M, opts *options.FindOneAndDeleteOptions) (*mongo.SingleResult, error)
	InsertOne(c context.Context, query bson.M, opts *options.InsertOneOptions) (primitive.ObjectID, error)
	UpdateOne(ctx context.Context, filter bson.M, query bson.D, opts *options.FindOneAndUpdateOptions) (*mongo.SingleResult, error)
}

type mongoDriver struct {
	client *mongo.Client
}

func NewMongoDriver(c *mongo.Client) MongoDriver {
	return &mongoDriver{client: c}
}

func (m *mongoDriver) FindOne(ctx context.Context, query bson.M, opts *options.FindOneOptions) (*mongo.SingleResult, error) {
	coll := m.client.Database(database).Collection(collection)
	res := coll.FindOne(ctx, query, opts)

	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("FindOne: unable to find book: %s", err)
	}

	return res, nil
}

func (m *mongoDriver) DeleteOne(ctx context.Context, query bson.M, opts *options.FindOneAndDeleteOptions) (*mongo.SingleResult, error) {
	coll := m.client.Database(database).Collection(collection)
	res := coll.FindOneAndDelete(ctx, query, opts)
	if err := res.Err(); err != nil {
		fmt.Errorf("DeleteOne: unable to delete book: %s", err)
	}

	return res, nil
}

func (m *mongoDriver) InsertOne(ctx context.Context, query bson.M, opts *options.InsertOneOptions) (primitive.ObjectID, error) {
	coll := m.client.Database(database).Collection(collection)
	res, err := coll.InsertOne(ctx, query, opts)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("InsertOne: unable to insert book: %s", err.Error())
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (m *mongoDriver) UpdateOne(ctx context.Context, filter bson.M, query bson.D, opts *options.FindOneAndUpdateOptions) (*mongo.SingleResult, error) {
	coll := m.client.Database(database).Collection(collection)
	res := coll.FindOneAndUpdate(ctx, filter, query, opts)
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("UpdateOne: unable to update book: %s", err)
	}

	return res, nil
}
