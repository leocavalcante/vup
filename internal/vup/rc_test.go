package vup

import (
	"testing"
)

func TestNewRC(t *testing.T) {
	testCases := []struct {
		name    string
		version string
		want    int
		err     error
	}{
		{
			name:    "WithValue",
			version: "rc1",
			want:    1,
			err:     nil,
		},
		{
			name:    "Empty",
			version: "",
			want:    0,
			err:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v, err := NewRC(tc.version)
			if err != tc.err {
				t.Errorf("NewRC() error = %v, wantErr %v", err, tc.err)
			}
			if v.Value() != tc.want {
				t.Errorf("NewRC() = %v, want %v", v.Value(), tc.want)
			}
		})
	}
}

func TestRC_Inc(t *testing.T) {
	r := &RC{value: 1}
	r.Inc(1)
	if r.Value() != 2 {
		t.Errorf("Inc() = %v, want %v", r.Value(), 2)
	}
}

func TestRC_Dec(t *testing.T) {
	r := &RC{value: 2}
	err := r.Dec(1)
	if err != nil {
		t.Errorf("Dec() error = %v, wantErr %v", err, nil)
	}
	if r.Value() != 1 {
		t.Errorf("Dec() = %v, want %v", r.Value(), 1)
	}
}

func TestRC_Dec_Err(t *testing.T) {
	r := &RC{value: 0}
	err := r.Dec(1)
	if err != ErrInvalidVersion {
		t.Errorf("Dec() error = %v, wantErr %v", err, ErrInvalidVersion)
	}
}

func TestRC_Set(t *testing.T) {
	r := &RC{value: 1}
	r.Set(10)
	if r.Value() != 10 {
		t.Errorf("Set() = %v, want %v", r.Value(), 10)
	}
}

func TestRC_Clear(t *testing.T) {
	r := &RC{value: 1}
	r.Clear()
	if r.Value() != 0 {
		t.Errorf("Clear() = %v, want %v", r.Value(), 0)
	}
}

func TestRC_String(t *testing.T) {
	r := &RC{value: 1}
	if r.String() != "rc1" {
		t.Errorf("String() = %v, want %v", r.String(), "rc1")
	}
}
