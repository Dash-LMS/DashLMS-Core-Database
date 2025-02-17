package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresQueryDriver struct {
	db *gorm.DB
}

func (p *PostgresQueryDriver) Connect(connectionString string) error {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *PostgresQueryDriver) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (p *PostgresQueryDriver) Read(collection string, filter interface{}) (interface{}, error) {
	var result interface{}
	err := p.db.Table(collection).Where(filter).First(&result).Error
	return result, err
}
