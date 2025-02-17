package mysql

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlQueryDriver struct {
	DB                *gorm.DB
	ConnMaxLifetime   time.Duration
	MaxOpenConns      int
	MaxIdleConns      int
	ConnectionTimeout time.Duration
}

func (m *MysqlQueryDriver) Connect(connectionString string) error {
	if connectionString == "" {
		return errors.New("connection string cannot be empty")
	}

	_, cancel := context.WithTimeout(context.Background(), m.ConnectionTimeout)
	defer cancel()

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Apply configurable connection settings
	sqlDB.SetConnMaxLifetime(m.ConnMaxLifetime)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)

	m.DB = db
	return nil
}

func (m *MysqlQueryDriver) Close() error {
	if m.DB == nil {
		return errors.New("database connection is not initialized")
	}

	sqlDB, err := m.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (m *MysqlQueryDriver) Read(collection string, filter interface{}) (interface{}, error) {
	if collection == "" || filter == nil {
		return nil, errors.New("collection name and filter cannot be empty")
	}

	if m.DB == nil { // Handle nil DB connection
		return nil, errors.New("database connection not established")
	}

	var result interface{}
	err := m.DB.Table(collection).Where(filter).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
