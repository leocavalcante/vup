package vup

import (
	"strconv"
)

type Patch struct {
	Part
	v int
}

func NewPatch(s string) (*Patch, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &Patch{
		v: v,
	}, nil
}

func (p *Patch) Inc(i int) {
	p.v += i
}

func (p *Patch) Dec(i int) error {
	if p.v-i < 0 {
		return ErrPatchLessThanZero
	}
	p.v -= i
	return nil
}

func (p *Patch) Value() int {
	return p.v
}

func (p *Patch) Clear() {
	p.v = 0
}

func (p *Patch) Set(i int) {
	p.v = i
}

func (p *Patch) String() string {
	return strconv.Itoa(p.v)
}
