package gongfeng

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"RFC3339 with offset", `"2024-01-15T10:30:00+08:00"`, false},
		{"RFC3339 UTC", `"2024-01-15T10:30:00Z"`, false},
		{"date only", `"2024-01-15"`, false},
		{"null", `null`, false},
		{"empty string", `""`, false},
		{"with milliseconds", `"2024-01-15T10:30:00.000+08:00"`, false},
		{"standard RFC3339", `"2024-01-15T10:30:00+00:00"`, false},
		{"invalid", `"not-a-date"`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tm Time
			err := tm.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON(%s) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestTimeMarshalJSON(t *testing.T) {
	var zero Time
	data, err := zero.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "null" {
		t.Fatalf("zero Time should marshal to null, got %s", data)
	}

	tm := Time{Time: time.Date(2024, 1, 15, 10, 30, 0, 0, time.FixedZone("CST", 8*3600))}
	data, err = tm.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `"2024-01-15T10:30:00+08:00"` {
		t.Fatalf("unexpected marshal result: %s", data)
	}
}

func TestTimeRoundTrip(t *testing.T) {
	type wrapper struct {
		T Time `json:"t"`
	}

	original := `{"t":"2024-01-15T10:30:00+08:00"}`
	var w wrapper
	if err := json.Unmarshal([]byte(original), &w); err != nil {
		t.Fatal(err)
	}

	data, err := json.Marshal(w)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != original {
		t.Fatalf("round trip failed: got %s, want %s", data, original)
	}
}

func TestTimeString(t *testing.T) {
	tm := Time{Time: time.Date(2024, 1, 15, 10, 30, 0, 0, time.FixedZone("CST", 8*3600))}
	if s := tm.String(); s != "2024-01-15T10:30:00+08:00" {
		t.Fatalf("unexpected String(): %s", s)
	}
}
