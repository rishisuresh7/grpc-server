package book

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"grpc-server/driver"
	"grpc-server/proto"
	"grpc-server/query"
)

type book struct {
	proto.UnimplementedBookServiceServer
	query  query.Book
	driver driver.MongoDriver
}

func NewBook(d driver.MongoDriver, q query.Book) proto.BookServiceServer {
	return &book{query: q, driver: d}
}

type Book struct {
	Id     string `bson:"_id"`
	Name   string `bson:"name"`
	Author string `bson:"author"`
}

func (bo *book) Get(ctx context.Context, b *proto.Book) (*proto.Book, error) {
	obId, err := primitive.ObjectIDFromHex(b.Id)
	if err != nil {
		return nil, fmt.Errorf("get: invalid 'id' format: %s", err.Error())
	}

	q, opts := bo.query.Get(obId)
	res, err := bo.driver.FindOne(ctx, q, opts)
	if err != nil {
		return nil, fmt.Errorf("get: unable to retrieve book: %s", err.Error())
	}

	var result Book
	err = res.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("get: unable to decode book: %s", err.Error())
	}

	return &proto.Book{
		Id:     result.Id,
		Name:   result.Name,
		Author: result.Author,
	}, nil
}

func (bo *book) Create(ctx context.Context, b *proto.Book) (*proto.Book, error) {
	q, opts := bo.query.Create(b)
	res, err := bo.driver.InsertOne(ctx, q, opts)
	if err != nil {
		return nil, fmt.Errorf("create: unbale to create book: %s", err.Error())
	}

	return &proto.Book{
		Id:     res.Hex(),
		Name:   b.Name,
		Author: b.Author,
	}, nil
}

func (bo *book) Update(ctx context.Context, b *proto.Book) (*proto.Book, error) {
	_, err := primitive.ObjectIDFromHex(b.Id)
	if err != nil {
		return nil, fmt.Errorf("update: invalid 'id' format: %s", err.Error())
	}

	filter, q, opts := bo.query.Update(b)
	res, err := bo.driver.UpdateOne(ctx, filter, q, opts)
	if err != nil {
		return nil, fmt.Errorf("update: unable to update book: %s", err.Error())
	}

	var result Book
	err = res.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("update: unable to decode book: %s", err.Error())
	}

	return &proto.Book{
		Id:     result.Id,
		Name:   result.Name,
		Author: result.Author,
	}, nil
}

func (bo *book) Delete(ctx context.Context, b *proto.Book) (*proto.Book, error) {
	obId, err := primitive.ObjectIDFromHex(b.Id)
	if err != nil {
		return nil, fmt.Errorf("delete: invalid 'id' format: %s", err.Error())
	}

	q, opts := bo.query.Delete(obId)
	res, err := bo.driver.DeleteOne(ctx, q, opts)
	if err != nil {
		return nil, fmt.Errorf("delete: unable to delete book: %s", err.Error())
	}

	var result Book
	err = res.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("delete: unable to decode book: %s", err.Error())
	}

	return &proto.Book{
		Id:     result.Id,
		Name:   result.Name,
		Author: result.Author,
	}, nil
}
