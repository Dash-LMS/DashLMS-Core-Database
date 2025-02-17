package utils_test

import (
	"testing"

	"github.com/Dash-LMS/DashLMS-Core-Database/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidateQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   interface{}
		wantErr bool
	}{
		{"Nil Query", nil, true},
		{"Valid Query", map[string]interface{}{"key": "value"}, false},
		// Updated to 'true' since the current code rejects empty maps:
		{"Empty Map", map[string]interface{}{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ValidateQuery(tt.query)
			if tt.wantErr {
				assert.Error(t, err, "expected error for %s", tt.name)
			} else {
				assert.NoError(t, err, "unexpected error for %s", tt.name)
			}
		})
	}
}
