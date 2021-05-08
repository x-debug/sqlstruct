package pkg

import (
	"encoding/json"
	"testing"
)

func assertUnmarshalString(t *testing.T, unvalid String, ok String, msg string) {
	if unvalid.Valid != ok.Valid || unvalid.String != ok.String {
		t.Errorf("assertUnmarshalBool: %s", msg)
	}
}

func TestMarshalJSONString(t *testing.T) {
	//test true value
	typ1 := NewString("a", true)
	d1, err := json.Marshal(&typ1)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d1, "\"a\"", "string value(a) error")

	//test false value
	typ2 := NewString("", true)
	d2, err := json.Marshal(&typ2)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d2, "\"\"", "string value('') error")

	//test nil value
	typ3 := NewString("", false)
	d3, err := json.Marshal(&typ3)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d3, "null", "string value(nil) error")
}

func TestUnmarshalJSONString(t *testing.T) {
	var typ1 String
	err := json.Unmarshal([]byte("\"a\""), &typ1)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalString(t, typ1, NewString("a", true), "unmarshal string(a) error")

	var typ2 String
	err = json.Unmarshal([]byte("\"\""), &typ2)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalString(t, typ2, NewString("", true), "unmarshal string('') error")

	var typ3 String
	err = json.Unmarshal([]byte("null"), &typ3)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalString(t, typ3, NewString("", false), "unmarshal string(nil) error")
}
