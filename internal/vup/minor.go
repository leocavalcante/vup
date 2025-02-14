package vup

import (
	"strconv"
)

type Minor struct {
	Part
	v int
}

func NewMinor(s string) (*Minor, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &Minor{
		v: v,
	}, nil
}

func (m *Minor) Inc(i int) {
	m.v += i
}

func (m *Minor) Dec(i int) error {
	if m.v <= 1 {
		return ErrMinorLessThanOne
	}

	m.v -= i
	return nil
}

func (m *Minor) String() string {
	return strconv.Itoa(m.v)
}
