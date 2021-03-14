package factory

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"grpc-server/book"
	"grpc-server/config"
	"grpc-server/driver"
	"grpc-server/proto"
	"grpc-server/query"
)

// Factory interface for object creation
type Factory interface {
	NewBook() proto.BookServiceServer
	NewMongoDriver() driver.MongoDriver
}

// factory implements Factory
type factory struct {
	config *config.AppConfig
	client *mongo.Client
	logger *logrus.Logger
}

var once sync.Once

// NewFactory constructor to initialize factory
func NewFactory(c *config.AppConfig, l *logrus.Logger) Factory {
	return &factory{logger: l, config: c}
}

// NewBook object of proto.BookServiceServer interface
func (f *factory) NewBook() proto.BookServiceServer {
	return book.NewBook(f.NewMongoDriver(), query.NewBookQuery())
}

// newMongoClient internal method to create mongodb connection
func (f *factory) newMongoClient() (*mongo.Client, error) {
	var err error
	once.Do(func() {
		ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(30*time.Second))
		defer cancel()

		opts := options.Client().ApplyURI(f.config.MongoUri)
		f.client, err = mongo.Connect(ctx, opts)
	})

	return f.client, nil
}

// NewMongoDriver creates an object of mongodb driver
func (f *factory) NewMongoDriver() driver.MongoDriver {
	client, err := f.newMongoClient()
	if err != nil {
		f.logger.Fatalf("NewMongoClient: unable to connect to mongo: %s", err.Error())
	}

	return driver.NewMongoDriver(client)
}
