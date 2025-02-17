package mysql

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlCommandDriver struct {
	DB                *gorm.DB
	ConnMaxLifetime   time.Duration
	MaxOpenConns      int
	MaxIdleConns      int
	ConnectionTimeout time.Duration
}

func (m *MysqlCommandDriver) Connect(connectionString string) error {
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

func (m *MysqlCommandDriver) Close() error {
	if m.DB == nil {
		return errors.New("database connection is not initialized")
	}

	sqlDB, err := m.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (m *MysqlCommandDriver) Create(collection string, data interface{}) error {
	if collection == "" || data == nil {
		return errors.New("collection name and data cannot be empty")
	}

	return m.DB.Table(collection).Create(data).Error
}

func (m *MysqlCommandDriver) Update(collection string, filter interface{}, update interface{}) error {
	if collection == "" || filter == nil || update == nil {
		return errors.New("collection name, filter, and update data cannot be empty")
	}

	return m.DB.Table(collection).Where(filter).Updates(update).Error
}

func (m *MysqlCommandDriver) Delete(collection string, filter interface{}) error {
	if collection == "" || filter == nil {
		return errors.New("collection name and filter cannot be empty")
	}

	return m.DB.Table(collection).Where(filter).Delete(nil).Error
}
