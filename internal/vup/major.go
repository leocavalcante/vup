package vup

import (
	"strconv"
	"strings"
)

type Major struct {
	Part
	p string
	v int
}

func NewMajor(s string) (*Major, error) {
	p := ""

	if strings.HasPrefix(s, "v") {
		p = s[:1]
		s = s[1:]
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &Major{
		p: p,
		v: v,
	}, nil
}

func (m *Major) Inc(i int) {
	m.v += i
}

func (m *Major) Dec(i int) error {
	if m.v <= 0 {
		return ErrMajorLessThanZero
	}

	m.v -= i
	return nil
}

func (m *Major) String() string {
	return m.p + strconv.Itoa(m.v)
}
