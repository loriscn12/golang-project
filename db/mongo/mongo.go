package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

// Cliente satisfies the mongo.Client interface.
type Client interface{
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)
	ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
}

func New(ctx context.Context, mongoAddress string) (Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddress))
	if err != nil {
		return nil, err
	}
	return client, nil
}
