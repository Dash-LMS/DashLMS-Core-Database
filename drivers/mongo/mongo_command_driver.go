package drivers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCommandDriver struct {
	client       *mongo.Client
	databaseName string
}

func (m *MongoCommandDriver) Connect(connectionString string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

func (m *MongoCommandDriver) Close() error {
	return m.client.Disconnect(context.TODO())
}

func (m *MongoCommandDriver) Create(collection string, data interface{}) error {
	coll := m.client.Database(m.databaseName).Collection(collection)
	_, err := coll.InsertOne(context.TODO(), data)
	return err
}

func (m *MongoCommandDriver) Update(collection string, filter interface{}, update interface{}) error {
	coll := m.client.Database(m.databaseName).Collection(collection)
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

func (m *MongoCommandDriver) Delete(collection string, filter interface{}) error {
	coll := m.client.Database(m.databaseName).Collection(collection)
	_, err := coll.DeleteOne(context.TODO(), filter)
	return err
}
