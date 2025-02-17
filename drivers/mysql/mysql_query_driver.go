package drivers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlQueryDriver struct {
	db *gorm.DB
}

func (m *MysqlQueryDriver) Connect(connectionString string) error {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	m.db = db
	return nil
}

func (m *MysqlQueryDriver) Close() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (m *MysqlQueryDriver) Read(collection string, filter interface{}) (interface{}, error) {
	var result interface{}
	err := m.db.Table(collection).Where(filter).First(&result).Error
	return result, err
}
