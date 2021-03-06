package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	sql.NullTime
}

func NewTime(value time.Time, valid bool) Time {
	return Time{NullTime: sql.NullTime{Time: value, Valid: valid}}
}

func (t Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		t.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &t.Time); err != nil {
		return fmt.Errorf("null: couldn't unmarshal JSON: %w", err)
	}

	t.Valid = true
	return nil
}
