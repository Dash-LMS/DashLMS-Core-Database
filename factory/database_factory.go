package factory

import (
	"errors"

	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mongo"
	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/mysql"
	"github.com/Dash-LMS/DashLMS-Core-Database/drivers/postgres"
	"github.com/Dash-LMS/DashLMS-Core-Database/interfaces"
)

func NewCommandDatabase(driverType string) (interfaces.CommandDatabase, error) {
	switch driverType {
	case "mongo":
		return &mongo.MongoCommandDriver{}, nil
	case "postgres":
		return &postgres.PostgresCommandDriver{}, nil
	case "mysql":
		return &mysql.MysqlCommandDriver{}, nil
	default:
		return nil, errors.New("unsupported driver type")
	}
}

func NewQueryDatabase(driverType string) (interfaces.QueryDatabase, error) {
	switch driverType {
	case "mongo":
		return &mongo.MongoQueryDriver{}, nil
	case "postgres":
		return &postgres.PostgresQueryDriver{}, nil
	case "mysql":
		return &mysql.MysqlQueryDriver{}, nil
	default:
		return nil, errors.New("unsupported driver type")
	}
}
