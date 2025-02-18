package mysql_test

import (
	"reflect"
	"testing"
	"time"

	test_mysql "github.com/Dash-LMS/DashLMS-Core-Database/drivers/mysql"
	"github.com/Dash-LMS/DashLMS-Core-Database/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMysqlQueryDriver_Read(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to MySQL")

	err = db.AutoMigrate(&TestTable{})
	assert.NoError(t, err, "failed to migrate test_table")

	db.Create(&TestTable{Key: "value"})
	queryDriver := &test_mysql.MysqlQueryDriver{DB: db}

	time.Sleep(2 * time.Second)

	query := map[string]interface{}{"key": "value"}
	err = utils.ValidateQuery(query)
	assert.NoError(t, err, "query validation failed")

	result, err := queryDriver.Read("test_table", query)
	assert.NoError(t, err, "failed to read record")
	assert.NotNil(t, result, "expected non-nil result")

	resultValue := reflect.ValueOf(result)
	if resultValue.Kind() == reflect.Map {
		resultMap := result.(map[string]interface{})
		assert.Equal(t, "value", resultMap["key"], "expected key to match")
	} else {
		t.Errorf("unexpected type: %T", result)
	}
}

func TestMysqlQueryDriver_ReadAll(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to MySQL")

	err = db.AutoMigrate(&TestTable{})
	assert.NoError(t, err, "failed to migrate test_table")

	db.Create(&TestTable{Key: "value1"})
	db.Create(&TestTable{Key: "value2"})
	queryDriver := &test_mysql.MysqlQueryDriver{DB: db}

	time.Sleep(2 * time.Second)

	query := map[string]interface{}{}
	err = utils.ValidateQuery(query)
	assert.NoError(t, err, "query validation failed")

	results, err := queryDriver.ReadAll("test_table", query)
	assert.NoError(t, err, "failed to read all records")

	resultsValue := reflect.ValueOf(results)
	if resultsValue.Kind() == reflect.Slice {
		assert.GreaterOrEqual(t, resultsValue.Len(), 2, "expected at least 2 records")
	} else {
		t.Errorf("unexpected type: %T", results)
	}
}

func TestMysqlQueryDriver_Count(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to MySQL")

	err = db.AutoMigrate(&TestTable{})
	assert.NoError(t, err, "failed to migrate test_table")

	db.Create(&TestTable{Key: "value"})
	queryDriver := &test_mysql.MysqlQueryDriver{DB: db}

	time.Sleep(2 * time.Second)

	query := map[string]interface{}{}
	err = utils.ValidateQuery(query)
	assert.NoError(t, err, "query validation failed")

	count, err := queryDriver.Count("test_table", query)
	assert.NoError(t, err, "failed to count records")
	assert.GreaterOrEqual(t, count, int64(1), "expected at least 1 record")
}

func TestMysqlQueryDriver_ExecuteRawSQL(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to MySQL")

	queryDriver := &test_mysql.MysqlQueryDriver{DB: db}

	sql := "CREATE TABLE IF NOT EXISTS raw_test (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))"
	err = queryDriver.ExecuteRawSQL(sql)
	assert.NoError(t, err, "failed to execute raw SQL")
}

func TestMysqlQueryDriver_QueryRawSQL(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "failed to connect to MySQL")

	queryDriver := &test_mysql.MysqlQueryDriver{DB: db}

	sql := "INSERT INTO raw_test (name) VALUES (?)"
	err = queryDriver.ExecuteRawSQL(sql, "test_name")
	assert.NoError(t, err, "failed to insert record using raw SQL")

	query := "SELECT * FROM raw_test WHERE name = ?"
	results, err := queryDriver.QueryRawSQL(query, "test_name")
	assert.NoError(t, err, "failed to query records using raw SQL")
	assert.NotEmpty(t, results, "expected non-empty result set")
}

func TestMysqlQueryDriver_Read_Failure(t *testing.T) {
	queryDriver := &test_mysql.MysqlQueryDriver{DB: nil}

	query := map[string]interface{}{"key": "value"}
	err := utils.ValidateQuery(query)
	assert.NoError(t, err, "query validation failed")

	_, err = queryDriver.Read("nonexistent_table", query)
	assert.Error(t, err, "expected read error for nonexistent table")
}
