package pkg

import (
	"encoding/json"
	"testing"
)

func assertUnmarshalInt(t *testing.T, unvalid Int, ok Int, msg string) {
	if unvalid.Valid != ok.Valid || unvalid.Int64 != ok.Int64 {
		t.Errorf("assertUnmarshalInt: %s", msg)
	}
}

func TestMarshalJSONInt(t *testing.T) {
	//test true value
	typ1 := NewInt(10, true)
	d1, err := json.Marshal(&typ1)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d1, "10", "bool int(10) error")

	//test false value
	typ2 := NewInt(0, true)
	d2, err := json.Marshal(&typ2)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d2, "0", "bool int(0) error")

	//test nil value
	typ3 := NewInt(0, false)
	d3, err := json.Marshal(&typ3)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d3, "null", "bool int(nil) error")
}

func TestUnmarshalJSONInt(t *testing.T) {
	var typ1 Int
	err := json.Unmarshal([]byte("10"), &typ1)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalInt(t, typ1, NewInt(10, true), "unmarshal int(10) error")

	var typ2 Int
	err = json.Unmarshal([]byte("0"), &typ2)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalInt(t, typ2, NewInt(0, true), "unmarshal int(0) error")

	var typ3 Int
	err = json.Unmarshal([]byte("null"), &typ3)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalInt(t, typ3, NewInt(0, false), "unmarshal int(nil) error")
}
