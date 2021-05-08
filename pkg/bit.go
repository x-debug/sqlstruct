package pkg

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Bit bool

type NullBit struct {
	Bit
	Valid bool
}

func NewBit(b byte, valid bool) NullBit {
	return NullBit{Valid: valid, Bit: Bit('1' == b)}
}

func (b Bit) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

func (b *Bit) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}

func (b *NullBit) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		b.Valid = false
		return nil
	}

	b.Valid = true
	b.Bit = v[0] == 1
	return nil
}

func (b *NullBit) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		b.Valid = false
		return nil
	}

	var byt byte
	if err := json.Unmarshal(data, &byt); err != nil {
		return fmt.Errorf("null: couldn't unmarshal JSON: %w", err)
	}

	b.Bit = byt == 1
	b.Valid = true
	return nil
}

func (b NullBit) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	if !b.Bit {
		return []byte{'0'}, nil
	}
	return []byte{'1'}, nil
}
