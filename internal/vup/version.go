package vup

import (
	"fmt"
	"strings"
)

type Version struct {
	Major Part
	Minor Part
	Patch Part
	RC    Part
}

func NewVersion(v string) (*Version, error) {
	var rc string
	p := strings.Split(v, "-")
	if len(p) > 1 {
		rc = p[1]
	}

	ps := strings.Split(p[0], ".")

	ma, err := NewMajor(ps[0])
	if err != nil {
		return nil, err
	}

	mi, err := NewMinor(ps[1])
	if err != nil {
		return nil, err
	}

	pa, err := NewPatch(ps[2])
	if err != nil {
		return nil, err
	}

	r, err := NewRC(rc)
	if err != nil {
		return nil, err
	}

	return &Version{
		Major: ma,
		Minor: mi,
		Patch: pa,
		RC:    r,
	}, nil
}

func (v *Version) String() string {
	s := fmt.Sprintf("%s.%s.%s", v.Major, v.Minor, v.Patch)
	if v.RC.Value() > 0 {
		s = fmt.Sprintf("%s-%s", s, v.RC)
	}
	return s
}
