package factory_test

import (
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/factory"
	"github.com/stretchr/testify/assert"
)

func TestNewCommandDatabase(t *testing.T) {
	factoryInstance := &factory.DatabaseFactory{}
	drivers := []string{"mongo", "postgres", "mysql"}
	for _, driver := range drivers {
		db, err := factoryInstance.NewCommandDatabase(driver)
		assert.NoError(t, err, "failed to create %s command database", driver)
		assert.NotNil(t, db, "expected non-nil database for %s", driver)
	}

	_, err := factoryInstance.NewCommandDatabase("invalid")
	assert.Error(t, err, "expected error for unsupported driver")
}

func TestNewQueryDatabase(t *testing.T) {
	factoryInstance := &factory.DatabaseFactory{}
	drivers := []string{"mongo", "postgres", "mysql"}
	for _, driver := range drivers {
		db, err := factoryInstance.NewQueryDatabase(driver)
		assert.NoError(t, err, "failed to create %s query database", driver)
		assert.NotNil(t, db, "expected non-nil database for %s", driver)
	}

	_, err := factoryInstance.NewQueryDatabase("invalid")
	assert.Error(t, err, "expected error for unsupported driver")
}
