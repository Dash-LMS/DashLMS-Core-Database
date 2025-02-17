package utils

import "errors"

func ValidateQuery(query interface{}) error {
	if query == nil {
		return errors.New("query cannot be nil")
	}
	return nil
}
