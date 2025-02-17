package mongo_test

import (
	"os"
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mongo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestMongoQueryDriver_Read(t *testing.T) {
	connectionString := os.Getenv("MONGO_URI")
	mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock)).Run("TestMongoQueryDriver_Read", func(mt *mtest.T) {
		// Mock a successful read operation
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "testDatabase.testCollection", mtest.FirstBatch, bson.D{{Key: "key", Value: "value"}}))

		driver := &mongo.MongoQueryDriver{}

		err := driver.Connect(connectionString)
		assert.NoError(t, err, "failed to connect to MongoDB")

		result, err := driver.Read("testCollection", bson.M{"key": "value"})
		assert.NoError(t, err, "failed to read document")
		assert.NotNil(t, result, "expected non-nil result")
	})

	mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock)).Run("TestMongoQueryDriver_Read_Failure", func(mt *mtest.T) {
		// Mock a failed read operation
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 11000, Message: "document not found"}))

		driver := &mongo.MongoQueryDriver{}
		err := driver.Connect(connectionString)
		assert.NoError(t, err, "failed to connect to MongoDB")

		_, err = driver.Read("testCollection", bson.M{"key": "nonexistent"})
		assert.Error(t, err, "expected document not found error")
	})
}
