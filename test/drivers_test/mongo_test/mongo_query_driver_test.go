package mongo_test

import (
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mongo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestMongoQueryDriver_Read(t *testing.T) {
	mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock)).Run("TestMongoQueryDriver_Read", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "testdb.testCollection", mtest.FirstBatch, bson.D{{Key: "key", Value: "value"}}))

		queryDriver := &mongo.MongoQueryDriver{ConnectionTimeout: 5}
		queryDriver.SetClient(mt.Client)
		queryDriver.SetDatabaseName("testdb")

		result, err := queryDriver.Read("testCollection", bson.M{"key": "value"})
		assert.NoError(t, err, "failed to read document")
		assert.NotNil(t, result, "expected a document to be found")
	})

	mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock)).Run("TestMongoQueryDriver_Read_Failure", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 11000, Message: "document not found"}))

		queryDriver := &mongo.MongoQueryDriver{ConnectionTimeout: 5}
		queryDriver.SetClient(mt.Client)
		queryDriver.SetDatabaseName("testdb")

		_, err := queryDriver.Read("testCollection", bson.M{"key": "nonexistent"})
		assert.Error(t, err, "expected a document not found error")
	})
}
