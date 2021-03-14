package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"grpc-server/proto"
)

// Book interface for all mongodb queries
type Book interface {
	Get(id primitive.ObjectID) (bson.M, *options.FindOneOptions)
	Create(b *proto.Book) (bson.M, *options.InsertOneOptions)
	Delete(id primitive.ObjectID) (bson.M, *options.FindOneAndDeleteOptions)
	Update(b *proto.Book) (bson.M, bson.D, *options.FindOneAndUpdateOptions)
}

// book implements Book
type book struct{}

// NewBookQuery constructor to initialize Book
func NewBookQuery() Book {
	return &book{}
}

// Get returns query to get a document with a given id
func (bo *book) Get(id primitive.ObjectID) (bson.M, *options.FindOneOptions) {
	opts := options.FindOneOptions{}
	return bson.M{"_id": id}, &opts
}

// Create returns query to insert a document
func (bo *book) Create(b *proto.Book) (bson.M, *options.InsertOneOptions) {
	opts := options.InsertOneOptions{}
	return bson.M{"name": b.Name, "author": b.Author}, &opts
}

// Delete returns query to delete a document with a given id
func (bo *book) Delete(id primitive.ObjectID) (bson.M, *options.FindOneAndDeleteOptions) {
	opts := options.FindOneAndDeleteOptions{}
	return bson.M{"_id": id}, &opts
}

// Update returns query to update a document with a given id
func (bo *book) Update(b *proto.Book) (bson.M, bson.D, *options.FindOneAndUpdateOptions) {
	opts := options.FindOneAndUpdateOptions{}
	opts.SetReturnDocument(options.After)
	id, _ := primitive.ObjectIDFromHex(b.Id)
	return bson.M{"_id": id}, bson.D{{Key: "$set", Value: bson.M{"name": b.Name, "author": b.Author}}}, &opts
}
