package mongo_test

import (
	"context"
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mongo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	test_mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoCommandDriver_Create(t *testing.T) {
	connectionString := "mongodb://localhost:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.3.9"
	// Connect directly to MongoDB to ensure collection is created
	client, err := test_mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	assert.NoError(t, err, "failed to connect to MongoDB")
	defer client.Disconnect(context.TODO())

	// Create database and collection if not exists
	db := client.Database("testdb")
	err = db.CreateCollection(context.TODO(), "testCollection")
	assert.NoError(t, err, "failed to create test collection")

	// Ensure MongoCommandDriver can connect and insert data
	driver := &mongo.MongoCommandDriver{}
	err = driver.Connect(connectionString)
	assert.NoError(t, err, "failed to connect via driver")

	// Insert a document
	err = driver.Create("testCollection", bson.M{"key": "value"})
	assert.NoError(t, err, "failed to insert document")
}
