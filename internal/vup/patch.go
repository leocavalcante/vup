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

func (p *Patch) Inc(i int) error {
	next, err := increment(p.v, i)
	if err != nil {
		return err
	}
	p.v = next
	return nil
}

func (p *Patch) Dec(i int) error {
	next, err := decrement(p.v, i, ErrPatchLessThanZero)
	if err != nil {
		return err
	}
	p.v = next
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
