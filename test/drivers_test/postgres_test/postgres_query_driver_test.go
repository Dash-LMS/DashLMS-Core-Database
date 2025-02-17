package postgres_test

import (
	"os"
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/postgres"
	"github.com/stretchr/testify/assert"
	test_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgresQueryDriver_Read(t *testing.T) {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(test_postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to PostgreSQL")

	// Ensure table exists
	err = db.AutoMigrate(&TestTable{})
	assert.NoError(t, err, "failed to migrate test_table")

	// Insert test data if needed
	db.Create(&TestTable{Key: "value"})

	driver := &postgres.PostgresQueryDriver{DB: db}

	// Pass a valid struct pointer
	result, err := driver.Read("test_table", map[string]interface{}{"key": "value"})
	assert.NoError(t, err, "failed to read record")
	assert.NotNil(t, result, "expected non-nil result")
}

func TestPostgresQueryDriver_Read_Failure(t *testing.T) {
	// Use a valid DSN so we can connect, but read from a nonexistent table
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(test_postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to PostgreSQL")

	queryDriver := &postgres.PostgresQueryDriver{DB: db}

	// Attempt to read from a table that doesn't exist
	_, err = queryDriver.Read("nonexistent_table", map[string]interface{}{"key": "value"})
	assert.Error(t, err, "expected read error for nonexistent table")
}
