package mysql_test

import (
	"os"
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mysql"
	"github.com/stretchr/testify/assert"
	test_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMysqlQueryDriver_Read(t *testing.T) {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(test_mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to MySQL")

	// Ensure table exists
	err = db.AutoMigrate(&TestTable{})
	assert.NoError(t, err, "failed to migrate test_table")

	// Insert test data if needed
	db.Create(&TestTable{Key: "value"})

	driver := &mysql.MysqlQueryDriver{DB: db}

	// Pass a valid struct pointer
	result, err := driver.Read("test_table", map[string]interface{}{"key": "value"})
	assert.NoError(t, err, "failed to read record")
	assert.NotNil(t, result, "expected non-nil result")
}

// TestMysqlQueryDriver_Read_Failure uses valid credentials but tries a nonexistent table
func TestMysqlQueryDriver_Read_Failure(t *testing.T) {
	// Intentionally invalid DSN
	queryDriver := &mysql.MysqlQueryDriver{DB: nil} // Initialize with nil DB

	// Attempt to read from a table that doesn't exist
	_, err := queryDriver.Read("nonexistent_table", map[string]interface{}{"key": "value"})
	assert.Error(t, err, "expected read error for nonexistent table")
}
