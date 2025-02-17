package postgres_test

import (
	"os"
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/postgres"
	"github.com/stretchr/testify/assert"
	test_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TestTable represents the table structure
type TestTable struct {
	ID  uint `gorm:"primaryKey"`
	Key string
}

// TableName overrides the default GORM table name.
func (TestTable) TableName() string {
	return "test_table"
}

func TestPostgresCommandDriver_Create(t *testing.T) {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "host=localhost port=5432 user=test password=test dbname=test sslmode=disable"
	}

	db, err := gorm.Open(test_postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to PostgreSQL")

	err = db.AutoMigrate(&TestTable{})
	assert.NoError(t, err, "failed to migrate test_table")

	driver := &postgres.PostgresCommandDriver{DB: db}
	err = driver.Create("test_table", map[string]interface{}{"key": "value"})
	assert.NoError(t, err, "failed to create record")
}

func TestPostgresCommandDriver_Create_Failure(t *testing.T) {
	// Deliberately invalid DSN/port so we expect a connection error
	dsn := "postgres://test:test@localhost:5433/testdb?sslmode=disable"
	_, err := gorm.Open(test_postgres.Open(dsn), &gorm.Config{})
	assert.Error(t, err, "expected connection error with invalid DSN/port")
}
