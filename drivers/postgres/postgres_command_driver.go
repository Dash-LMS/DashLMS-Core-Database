package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresCommandDriver struct {
	db *gorm.DB
}

func (p *PostgresCommandDriver) Connect(connectionString string) error {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *PostgresCommandDriver) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (p *PostgresCommandDriver) Create(collection string, data interface{}) error {
	return p.db.Table(collection).Create(data).Error
}

func (p *PostgresCommandDriver) Update(collection string, filter interface{}, update interface{}) error {
	return p.db.Table(collection).Where(filter).Updates(update).Error
}

func (p *PostgresCommandDriver) Delete(collection string, filter interface{}) error {
	return p.db.Table(collection).Where(filter).Delete(nil).Error
}
