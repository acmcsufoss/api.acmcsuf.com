package mapper

import (
	"database/sql"
	"time"
)

// ---- Functions to check validity ----
func intToNullInt64(i *int) sql.NullInt64 {
	var val int64
	var valid bool
	if i != nil {
		deref := *i
		val = int64(deref)
	}

	return sql.NullInt64{Int64: val, Valid: valid}
}

func stringToNullString(s *string) sql.NullString {
	var val string
	var valid bool
	if s != nil {
		val = *s
	}

	return sql.NullString{String: val, Valid: valid}
}

func boolToNullBool(b *bool) sql.NullBool {
	var val bool
	var valid bool
	if b != nil {
		val = *b
	}

	return sql.NullBool{Bool: val, Valid: valid}
}

func timeToNullInt64(t *time.Time) sql.NullInt64 {
	var val int64
	var valid bool
	if t != nil {
		deref := *t
		val = deref.Unix()
	}

	return sql.NullInt64{Int64: val, Valid: valid}
}
