package vup

import (
	"fmt"
	"strings"
)

type Version struct {
	Major Part
	Minor Part
	Patch Part
}

func NewVersion(v string) (*Version, error) {
	ps := strings.Split(v, ".")

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

	return &Version{
		Major: ma,
		Minor: mi,
		Patch: pa,
	}, nil
}

func (v *Version) String() string {
	return fmt.Sprintf("%s.%s.%s", v.Major, v.Minor, v.Patch)
}
