package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type Int struct {
	sql.NullInt64
}

func NewInt(value int64, valid bool) Int {
	return Int{NullInt64: sql.NullInt64{Valid: valid, Int64: value}}
}

func (v *Int) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		v.Valid = false
		return nil
	}

	if err := json.Unmarshal(b, &v.Int64); err != nil {
		return fmt.Errorf("null: couldn't unmarshal JSON: %w", err)
	}

	v.Valid = true
	return nil
}

func (v *Int) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return []byte(strconv.FormatInt(int64(v.Int64), 10)), nil
	} else {
		return []byte("null"), nil
	}
}
