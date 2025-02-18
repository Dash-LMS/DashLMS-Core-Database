package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	test_mongo "github.com/Dash-LMS/DashLMS-Core-Database/drivers/mongo"
)

const testConnectionString = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=10000"
const testDatabase = "testdb"
const testCollection = "testCollection"

func setupTestDatabase(t *testing.T) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testConnectionString))
	assert.NoError(t, err, "failed to connect to MongoDB")
	return client
}

func cleanupTestDatabase(t *testing.T, client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := client.Database(testDatabase).Drop(ctx)
	assert.NoError(t, err, "failed to clean up test database")
	client.Disconnect(ctx)
}

func TestMongoCommandDriver(t *testing.T) {
	driver := &test_mongo.MongoCommandDriver{
		ConnectionTimeout: 10 * time.Second,
	}
	err := driver.Connect(testConnectionString)
	assert.NoError(t, err, "failed to connect via MongoCommandDriver")
	defer driver.Close()
	driver.SetCommandDatabase(testDatabase)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := setupTestDatabase(t)
	defer cleanupTestDatabase(t, client)

	err = client.Database(testDatabase).CreateCollection(ctx, testCollection)
	assert.NoError(t, err, "failed to create test collection")

	// Test Create
	t.Run("Create", func(t *testing.T) {
		document := bson.M{"name": "test", "value": int32(123)}
		err := driver.Create(testCollection, document)
		assert.NoError(t, err, "failed to create document")

		// Verify document insertion
		collection := client.Database(testDatabase).Collection(testCollection)
		var result bson.M
		err = collection.FindOne(ctx, bson.M{"name": "test"}).Decode(&result)
		assert.NoError(t, err, "failed to find created document")
		assert.Equal(t, int32(123), result["value"], "document value mismatch")
	})

	// Test Update
	t.Run("Update", func(t *testing.T) {
		filter := bson.M{"name": "test"}
		update := bson.M{"value": int32(456)}
		err := driver.Update(testCollection, filter, update)
		assert.NoError(t, err, "failed to update document")

		// Verify update
		collection := client.Database(testDatabase).Collection(testCollection)
		var result bson.M
		err = collection.FindOne(ctx, filter).Decode(&result)
		assert.NoError(t, err, "failed to find updated document")
		assert.Equal(t, int32(456), result["value"], "document value mismatch after update")
	})

	// Test Delete
	t.Run("Delete", func(t *testing.T) {
		filter := bson.M{"name": "test"}
		err := driver.Delete(testCollection, filter)
		assert.NoError(t, err, "failed to delete document")

		// Verify deletion
		collection := client.Database(testDatabase).Collection(testCollection)
		var result bson.M
		err = collection.FindOne(ctx, filter).Decode(&result)
		assert.Error(t, err, "expected error when finding deleted document")
	})
}
