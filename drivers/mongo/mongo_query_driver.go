package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoQueryDriver struct {
	client            *mongo.Client
	databaseName      string
	ConnectionTimeout time.Duration
}

func (m *MongoQueryDriver) Connect(connectionString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.ConnectionTimeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	m.client = client
	return nil
}

func (m *MongoQueryDriver) SetQueryDatabase(databaseName string) {
	m.databaseName = databaseName
}

func (m *MongoQueryDriver) Close() error {
	const HALF = 2
	ctx, cancel := context.WithTimeout(context.Background(), m.ConnectionTimeout/HALF)
	defer cancel()

	return m.client.Disconnect(ctx)
}

func (m *MongoQueryDriver) Read(collection string, filter interface{}) (interface{}, error) {
	const HALF = 2
	if filter == nil {
		return nil, errors.New("filter cannot be nil")
	}

	coll := m.client.Database(m.databaseName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), m.ConnectionTimeout/HALF)
	defer cancel()

	result := coll.FindOne(ctx, filter)
	var doc interface{}
	err := result.Decode(&doc)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return doc, err
}
