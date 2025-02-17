package drivers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlCommandDriver struct {
	db *gorm.DB
}

func (m *MysqlCommandDriver) Connect(connectionString string) error {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	m.db = db
	return nil
}

func (m *MysqlCommandDriver) Close() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (m *MysqlCommandDriver) Create(collection string, data interface{}) error {
	return m.db.Table(collection).Create(data).Error
}

func (m *MysqlCommandDriver) Update(collection string, filter interface{}, update interface{}) error {
	return m.db.Table(collection).Where(filter).Updates(update).Error
}

func (m *MysqlCommandDriver) Delete(collection string, filter interface{}) error {
	return m.db.Table(collection).Where(filter).Delete(nil).Error
}
