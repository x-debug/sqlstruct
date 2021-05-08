package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
)

type Bool struct {
	sql.NullBool
}

func NewBool(b bool, valid bool) Bool {
	return Bool{NullBool: sql.NullBool{Valid: valid, Bool: b}}
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		b.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &b.Bool); err != nil {
		return fmt.Errorf("null: couldn't unmarshal JSON: %w", err)
	}

	b.Valid = true
	return nil
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	if !b.Bool {
		return []byte("false"), nil
	}
	return []byte("true"), nil
}
