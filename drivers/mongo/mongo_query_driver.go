package drivers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoQueryDriver struct {
	client       *mongo.Client
	databaseName string
}

func (m *MongoQueryDriver) Connect(connectionString string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

func (m *MongoQueryDriver) Close() error {
	return m.client.Disconnect(context.TODO())
}

func (m *MongoQueryDriver) Read(collection string, filter interface{}) (interface{}, error) {
	coll := m.client.Database(m.databaseName).Collection(collection)
	result := coll.FindOne(context.TODO(), filter)
	var doc interface{}
	err := result.Decode(&doc)
	return doc, err
}
