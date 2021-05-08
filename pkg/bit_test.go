package pkg

import (
	"encoding/json"
	"testing"
)

func assertUnmarshalBit(t *testing.T, unvalid NullBit, ok NullBit, msg string) {
	if unvalid.Valid != ok.Valid || unvalid.Bit != ok.Bit {
		t.Errorf("assertUnmarshalBit: %s, expected: %v, actual: %v", msg, ok.Bit, unvalid.Bit)
	}
}

func TestMarshalJSONBit(t *testing.T) {
	//test true value
	typ1 := NewBit('1', true)
	d1, err := json.Marshal(&typ1)
	if err != nil {
		t.Errorf("json.Marshal1 error: %v", err)
		return
	}

	assertJSON(t, d1, "1", "bit value(0x01) error")

	//test false value
	typ2 := NewBit('0', true)
	d2, err := json.Marshal(&typ2)
	if err != nil {
		t.Errorf("json.Marshal2 error: %v", err)
		return
	}

	assertJSON(t, d2, "0", "bit value(0x00) error")

	//test nil value
	typ3 := NewBit('0', false)
	d3, err := json.Marshal(&typ3)
	if err != nil {
		t.Errorf("json.Marshal3 error: %v", err)
		return
	}

	assertJSON(t, d3, "null", "bit value(nil) error")
}

func TestUnmarshalJSONBit(t *testing.T) {
	var typ1 NullBit
	err := json.Unmarshal([]byte{'1'}, &typ1)
	if err != nil {
		t.Errorf("json.Unmarshal1 error: %v", err)
		return
	}

	assertUnmarshalBit(t, typ1, NewBit('1', true), "unmarshal bit(0x01) error")
	return

	var typ2 NullBit
	err = json.Unmarshal([]byte{'0'}, &typ2)
	if err != nil {
		t.Errorf("json.Unmarshal2 error: %v", err)
		return
	}

	assertUnmarshalBit(t, typ2, NewBit('0', true), "unmarshal bit(0x00) error")

	var typ3 NullBit
	err = json.Unmarshal([]byte("null"), &typ3)
	if err != nil {
		t.Errorf("json.Unmarshal3 error: %v", err)
		return
	}

	assertUnmarshalBit(t, typ3, NewBit('0', false), "unmarshal bit(nil) error")
}
