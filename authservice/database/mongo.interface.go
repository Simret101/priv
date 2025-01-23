package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CollectionInterface defines the methods that a MongoDB collection should implement
type CollectionInterface interface {
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) SingleResultInterface
	Find(context.Context, interface{}, ...*options.FindOptions) (CursorInterface, error)
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (DeleteResultInterface, error)
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) SingleResultInterface
	Indexes() IndexView
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
}

// IndexView defines the methods for managing indexes in a MongoDB collection
type IndexView interface {
	CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
}

// CursorInterface defines the methods for iterating over a MongoDB query result set
type CursorInterface interface {
	Next(context.Context) bool
	Decode(interface{}) error
	Close(context.Context) error
}

// SingleResultInterface defines the methods for handling a single MongoDB query result
type SingleResultInterface interface {
	Decode(v interface{}) error
}

// DeleteResultInterface defines the methods for handling the result of a delete operation in MongoDB
type DeleteResultInterface interface {
	DeletedCount() int64
}
