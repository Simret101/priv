package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCursor wraps the mongo.Cursor to implement the CursorInterface
type MongoCursor struct {
	*mongo.Cursor
}

// Next advances the cursor to the next document in the result set
func (c *MongoCursor) Next(ctx context.Context) bool {
	return c.Cursor.Next(ctx)
}

// Decode decodes the current document into the provided value
func (c *MongoCursor) Decode(v interface{}) error {
	return c.Cursor.Decode(v)
}

// Close closes the cursor
func (c *MongoCursor) Close(ctx context.Context) error {
	return c.Cursor.Close(ctx)
}

// MongoIndexView wraps the mongo.IndexView to implement the IndexView interface
type MongoIndexView struct {
	mongo.IndexView
}

// CreateOne creates a single index in the collection
func (MI *MongoIndexView) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return MI.IndexView.CreateOne(ctx, model, opts...)
}

// MongoSingleResult wraps the mongo.SingleResult to implement the SingleResultInterface
type MongoSingleResult struct {
	*mongo.SingleResult
}

// Decode decodes the single result into the provided value
func (r *MongoSingleResult) Decode(v interface{}) error {
	return r.SingleResult.Decode(v)
}

// MongoDeleteResult wraps the mongo.DeleteResult to implement the DeleteResultInterface
type MongoDeleteResult struct {
	*mongo.DeleteResult
}

// DeletedCount returns the number of documents deleted
func (r *MongoDeleteResult) DeletedCount() int64 {
	return r.DeleteResult.DeletedCount
}

// MongoCollection wraps the mongo.Collection to implement the CollectionInterface
type MongoCollection struct {
	*mongo.Collection
}

// FindOne finds a single document in the collection
func (c *MongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) SingleResultInterface {
	return &MongoSingleResult{SingleResult: c.Collection.FindOne(ctx, filter, opts...)}
}

// Find finds multiple documents in the collection
func (c *MongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (CursorInterface, error) {
	cursor, err := c.Collection.Find(ctx, filter, opts...)
	return &MongoCursor{Cursor: cursor}, err
}

// InsertOne inserts a single document into the collection
func (c *MongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.Collection.InsertOne(ctx, document, opts...)
}

// UpdateOne updates a single document in the collection
func (c *MongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.Collection.UpdateOne(ctx, filter, update, opts...)
}

// FindOneAndUpdate finds a single document and updates it
func (c *MongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) SingleResultInterface {
	return &MongoSingleResult{SingleResult: c.Collection.FindOneAndUpdate(ctx, filter, update, opts...)}
}

// DeleteOne deletes a single document from the collection
func (c *MongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (DeleteResultInterface, error) {
	result, err := c.Collection.DeleteOne(ctx, filter, opts...)
	return &MongoDeleteResult{DeleteResult: result}, err
}

// Indexes returns the index view of the collection
func (c *MongoCollection) Indexes() IndexView {
	return &MongoIndexView{IndexView: c.Collection.Indexes()}
}

// CountDocuments counts the number of documents in the collection
func (c *MongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return c.Collection.CountDocuments(ctx, filter, opts...)
}
