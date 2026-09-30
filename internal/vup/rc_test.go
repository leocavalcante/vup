package vup

import (
	"math"
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
	r := &RC{value: 1, present: true}
	if err := r.Inc(1); err != nil {
		t.Fatalf("Inc() error = %v", err)
	}
	if r.Value() != 2 {
		t.Errorf("Inc() = %v, want %v", r.Value(), 2)
	}
}

func TestRC_Inc_RejectsInvalid(t *testing.T) {
	r := &RC{value: 1, present: true}
	if err := r.Inc(-2); err == nil {
		t.Fatal("Inc(-2) error = nil, want error")
	}
	if r.Value() != 1 {
		t.Errorf("Value() = %v, want 1", r.Value())
	}

	overflow := &RC{value: math.MaxInt, present: true}
	if err := overflow.Inc(1); err == nil {
		t.Fatal("Inc(1) at MaxInt error = nil, want error")
	}
	if overflow.Value() != math.MaxInt {
		t.Errorf("Value() = %v, want MaxInt", overflow.Value())
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
	r := &RC{value: 0, present: true}
	err := r.Dec(1)
	if err != ErrRCLessThanZero {
		t.Errorf("Dec() error = %v, wantErr %v", err, ErrRCLessThanZero)
	}

	negative := &RC{value: 1, present: true}
	if err := negative.Dec(-1); err == nil {
		t.Fatal("Dec(-1) error = nil, want error")
	}
	if negative.Value() != 1 {
		t.Errorf("Value() = %v, want 1", negative.Value())
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
	r := &RC{value: 1, present: true}
	r.Clear()
	if r.Value() != 0 {
		t.Errorf("Clear() = %v, want %v", r.Value(), 0)
	}
	if r.Present() {
		t.Errorf("Clear() present = true, want false")
	}
}

func TestNewRC_Canonical(t *testing.T) {
	valid := []struct {
		in   string
		want int
	}{
		{in: "rc0", want: 0},
		{in: "rc1", want: 1},
		{in: "rc10", want: 10},
	}
	for _, tc := range valid {
		t.Run(tc.in, func(t *testing.T) {
			got, err := NewRC(tc.in)
			if err != nil {
				t.Fatalf("NewRC() error = %v", err)
			}
			if got.Value() != tc.want {
				t.Errorf("Value() = %v, want %v", got.Value(), tc.want)
			}
			if !got.(*RC).Present() {
				t.Errorf("Present() = false, want true")
			}
		})
	}

	invalid := []string{"xrc1", "rc1rc2", "rc01", "rc", "RC1"}
	for _, in := range invalid {
		t.Run(in, func(t *testing.T) {
			if _, err := NewRC(in); err == nil {
				t.Fatalf("NewRC(%q) error = nil, want error", in)
			}
		})
	}
}

func TestRC_String(t *testing.T) {
	r := &RC{value: 1}
	if r.String() != "rc1" {
		t.Errorf("String() = %v, want %v", r.String(), "rc1")
	}
}
