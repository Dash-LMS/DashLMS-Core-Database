package factory

import (
	"fmt"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mongo"
	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mysql"
	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/postgres"
	"github.com/Dash-LMS/DashLMS-Core-Database/interfaces"
)

type DatabaseFactory struct{}

// NewCommandDatabase creates a command database driver based on the provided type.
func (f *DatabaseFactory) NewCommandDatabase(driverType string) (interfaces.CommandDatabase, error) {
	switch driverType {
	case "mongo":
		return &mongo.MongoCommandDriver{}, nil
	case "postgres":
		return &postgres.PostgresCommandDriver{}, nil
	case "mysql":
		return &mysql.MysqlCommandDriver{}, nil
	default:
		return nil, fmt.Errorf("unsupported command database driver: %s", driverType)
	}
}

// NewQueryDatabase creates a query database driver based on the provided type.
func (f *DatabaseFactory) NewQueryDatabase(driverType string) (interfaces.QueryDatabase, error) {
	switch driverType {
	case "mongo":
		return &mongo.MongoQueryDriver{}, nil
	case "postgres":
		return &postgres.PostgresQueryDriver{}, nil
	case "mysql":
		return &mysql.MysqlQueryDriver{}, nil
	default:
		return nil, fmt.Errorf("unsupported query database driver: %s", driverType)
	}
}
