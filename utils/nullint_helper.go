package utils

import "database/sql"

func NullInt64ToPtr(value sql.NullInt64) *int {
	if !value.Valid {
		return nil
	}

	id := int(value.Int64)
	return &id
}
