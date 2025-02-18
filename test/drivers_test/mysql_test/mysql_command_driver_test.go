package mysql_test

import (
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mysql"
	"github.com/stretchr/testify/assert"
	test_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestTable struct {
	ID  uint   `gorm:"primaryKey"`
	Key string `gorm:"index"`
}

func (TestTable) TableName() string {
	return "test_table"
}

// TestMysqlCommandDriver_Create ensures we can create a record with valid credentials.
func TestMysqlCommandDriver_Create(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(test_mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to MySQL")

	err = db.AutoMigrate(&TestTable{})
	assert.NoError(t, err, "failed to migrate test_table")

	commandDriver := &mysql.MysqlCommandDriver{DB: db}
	err = commandDriver.Create("test_table", map[string]interface{}{"key": "value"})
	assert.NoError(t, err, "failed to create record")
}

// TestMysqlCommandDriver_Create_Failure tests invalid credentials, expecting a connection error.
func TestMysqlCommandDriver_Create_Failure(t *testing.T) {
	// Intentionally invalid DSN
	dsn := "invalid_user:invalid_password@tcp(localhost:3306)/invalid_db"

	_, err := gorm.Open(test_mysql.Open(dsn), &gorm.Config{})
	// We expect an error (e.g., "Access denied")
	assert.Error(t, err, "expected connection error with invalid credentials")
}
