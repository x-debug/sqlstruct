package pkg

import (
	"encoding/json"
	"testing"
	"time"
)

func assertUnmarshalTime(t *testing.T, unvalid Time, ok Time, msg string) {
	if unvalid.Valid != ok.Valid || unvalid.Time != ok.Time {
		t.Errorf("assertUnmarshalBool: %s", msg)
	}
}

func TestMarshalJSONTime(t *testing.T) {
	//test true value
	t1, _ := time.Parse("2006-01-02 15:04:05", "1989-10-18 10:10:10")
	typ1 := NewTime(t1, true)
	d1, err := json.Marshal(&typ1)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	//timezone
	assertJSON(t, d1, "\"1989-10-18T10:10:10Z\"", "time value(1989-10-18 10:10:10) error")

	//test false value
	t2, _ := time.Parse("2006-01-02 15:04:05", "1989-Feb-18 10:10:10")
	typ2 := NewTime(t2, true)
	d2, err := json.Marshal(&typ2)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d2, "\"0001-01-01T00:00:00Z\"", "time value(0001-01-01 00:00:00 +0000 UTC) error")

	//test nil value
	typ3 := NewTime(t2, false)
	d3, err := json.Marshal(&typ3)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
		return
	}

	assertJSON(t, d3, "null", "time value(nil) error")
}

func TestUnmarshalJSONTime(t *testing.T) {
	var typ1 Time
	err := json.Unmarshal([]byte("\"1989-10-18T10:10:10Z\""), &typ1)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	t1, _ := time.Parse("2006-01-02 15:04:05", "1989-10-18 10:10:10")
	assertUnmarshalTime(t, typ1, NewTime(t1, true), "unmarshal (true) error")

	var typ2 Time
	err = json.Unmarshal([]byte("\"0001-01-01T00:00:00Z\""), &typ2)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	t2, _ := time.Parse("2006-01-02 15:04:05", "1989-Feb-18 10:10:10")
	assertUnmarshalTime(t, typ2, NewTime(t2, true), "unmarshal (false) error")

	var typ3 Time
	err = json.Unmarshal([]byte("null"), &typ3)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
		return
	}

	assertUnmarshalTime(t, typ3, NewTime(t2, false), "unmarshal (nil) error")
}
