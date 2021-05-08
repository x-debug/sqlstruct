package pkg

import (
	"encoding/json"
	"testing"
)

func assertJSON(t *testing.T, data []byte, ok string, msg string) {
	if string(data) != ok {
		t.Errorf("assertJSON: %s", msg)
	}
}

func assertUnmarshalBool(t *testing.T, unvalid Bool, ok Bool, msg string) {
	if unvalid.Valid != ok.Valid || unvalid.Bool != ok.Bool {
		t.Errorf("assertUnmarshalBool: %s", msg)
	}
}

func TestMarshalJSON(t *testing.T) {
	//test true value
	typ1 := NewBool(true, true)
	d1, err := json.Marshal(&typ1)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d1, "true", "bool value(true) error")

	//test false value
	typ2 := NewBool(false, true)
	d2, err := json.Marshal(&typ2)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d2, "false", "bool value(false) error")

	//test nil value
	typ3 := NewBool(false, false)
	d3, err := json.Marshal(&typ3)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d3, "null", "bool value(nil) error")
}

func TestUnmarshalJSON(t *testing.T) {
	var typ1 Bool
	err := json.Unmarshal([]byte("true"), &typ1)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalBool(t, typ1, NewBool(true, true), "unmarshal (true) error")

	var typ2 Bool
	err = json.Unmarshal([]byte("false"), &typ2)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalBool(t, typ2, NewBool(false, true), "unmarshal (false) error")

	var typ3 Bool
	err = json.Unmarshal([]byte("null"), &typ3)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalBool(t, typ3, NewBool(false, false), "unmarshal (nil) error")
}
