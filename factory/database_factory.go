package factory

import (
	"errors"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers"
	"github.com/Dash-LMS/DashLMS-Core-Database/interfaces"
)

func NewCommandDatabase(driverType string) (interfaces.CommandDatabase, error) {
	switch driverType {
	case "mongo":
		return &drivers.MongoCommandDriver{}, nil
	case "postgres":
		return &drivers.PostgresCommandDriver{}, nil
	case "mysql":
		return &drivers.MysqlCommandDriver{}, nil
	default:
		return nil, errors.New("unsupported driver type")
	}
}

func NewQueryDatabase(driverType string) (interfaces.QueryDatabase, error) {
	switch driverType {
	case "mongo":
		return &drivers.MongoQueryDriver{}, nil
	case "postgres":
		return &drivers.PostgresQueryDriver{}, nil
	case "mysql":
		return &drivers.MysqlQueryDriver{}, nil
	default:
		return nil, errors.New("unsupported driver type")
	}
}
