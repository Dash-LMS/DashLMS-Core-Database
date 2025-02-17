package postgres

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresQueryDriver struct {
	DB                *gorm.DB
	ConnMaxLifetime   time.Duration
	MaxOpenConns      int
	MaxIdleConns      int
	ConnectionTimeout time.Duration
}

func (p *PostgresQueryDriver) Connect(connectionString string) error {
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

func (p *PostgresQueryDriver) Close() error {
	if p.DB == nil {
		return errors.New("database connection is not initialized")
	}

	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (p *PostgresQueryDriver) Read(collection string, filter interface{}) (interface{}, error) {
	if collection == "" || filter == nil {
		return nil, errors.New("collection name and filter cannot be empty")
	}

	var result interface{}
	err := p.DB.Table(collection).Where(filter).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return result, err
}
