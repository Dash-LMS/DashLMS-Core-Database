package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCommandDriver struct {
	client            *mongo.Client
	databaseName      string
	ConnectionTimeout time.Duration
}

func (m *MongoCommandDriver) Connect(connectionString string) error {
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

func (m *MongoCommandDriver) SetCommandDatabase(databaseName string) {
	m.databaseName = databaseName
}

func (m *MongoCommandDriver) Close() error {
	const HALF = 2
	ctx, cancel := context.WithTimeout(context.Background(), m.ConnectionTimeout/HALF)
	defer cancel()

	return m.client.Disconnect(ctx)
}

func (m *MongoCommandDriver) Create(collection string, data interface{}) error {
	const HALF = 2
	if data == nil {
		return errors.New("data cannot be nil")
	}

	coll := m.client.Database(m.databaseName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), m.ConnectionTimeout/HALF)
	defer cancel()

	_, err := coll.InsertOne(ctx, data)
	return err
}

func (m *MongoCommandDriver) Update(collection string, filter interface{}, update interface{}) error {
	const HALF = 2
	if filter == nil || update == nil {
		return errors.New("filter and update cannot be nil")
	}

	coll := m.client.Database(m.databaseName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), m.ConnectionTimeout/HALF)
	defer cancel()

	_, err := coll.UpdateOne(ctx, filter, bson.M{"$set": update})
	return err
}

func (m *MongoCommandDriver) Delete(collection string, filter interface{}) error {
	const HALF = 2
	if filter == nil {
		return errors.New("filter cannot be nil")
	}

	coll := m.client.Database(m.databaseName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), m.ConnectionTimeout/HALF)
	defer cancel()

	_, err := coll.DeleteOne(ctx, filter)
	return err
}
