package pkg

import (
	"encoding/json"
	"testing"
)

func assertUnmarshalFloat(t *testing.T, unvalid Float, ok Float, msg string) {
	if unvalid.Valid != ok.Valid || unvalid.Float64 != ok.Float64 {
		t.Errorf("assertUnmarshalFloat: %s", msg)
	}
}

func TestMarshalJSONFloat(t *testing.T) {
	//test true value
	typ1 := NewFloat(8.8, true)
	d1, err := json.Marshal(&typ1)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d1, "8.8", "float value(8.8) error")

	//test false value
	typ2 := NewFloat(0.0, true)
	d2, err := json.Marshal(&typ2)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d2, "0", "float value(0.0) error")

	//test nil value
	typ3 := NewFloat(0.0, false)
	d3, err := json.Marshal(&typ3)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d3, "null", "float value(nil) error")
}

func TestUnmarshalJSONFloat(t *testing.T) {
	var typ1 Float
	err := json.Unmarshal([]byte("8.8"), &typ1)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalFloat(t, typ1, NewFloat(8.8, true), "unmarshal float(8.8) error")

	var typ2 Float
	err = json.Unmarshal([]byte("0.0"), &typ2)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalFloat(t, typ2, NewFloat(0.0, true), "unmarshal float(0.0) error")

	var typ3 Float
	err = json.Unmarshal([]byte("null"), &typ3)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalFloat(t, typ3, NewFloat(0.0, false), "unmarshal float(nil) error")
}
