package postgres

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresCommandDriver struct {
	DB                *gorm.DB
	ConnMaxLifetime   time.Duration
	MaxOpenConns      int
	MaxIdleConns      int
	ConnectionTimeout time.Duration
}

func (p *PostgresCommandDriver) Connect(connectionString string) error {
	if connectionString == "" {
		return errors.New("connection string cannot be empty")
	}

	_, cancel := context.WithTimeout(context.Background(), p.ConnectionTimeout)
	defer cancel()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Apply configurable connection settings
	sqlDB.SetConnMaxLifetime(p.ConnMaxLifetime)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)

	p.DB = db
	return nil
}

func (p *PostgresCommandDriver) Close() error {
	if p.DB == nil {
		return errors.New("database connection is not initialized")
	}

	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (p *PostgresCommandDriver) Create(collection string, data interface{}) error {
	if collection == "" || data == nil {
		return errors.New("collection name and data cannot be empty")
	}

	return p.DB.Table(collection).Create(data).Error
}

func (p *PostgresCommandDriver) Update(collection string, filter interface{}, update interface{}) error {
	if collection == "" || filter == nil || update == nil {
		return errors.New("collection name, filter, and update data cannot be empty")
	}

	return p.DB.Table(collection).Where(filter).Updates(update).Error
}

func (p *PostgresCommandDriver) Delete(collection string, filter interface{}) error {
	if collection == "" || filter == nil {
		return errors.New("collection name and filter cannot be empty")
	}

	return p.DB.Table(collection).Where(filter).Delete(nil).Error
}
