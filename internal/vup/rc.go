package vup

import (
	"fmt"
	"strconv"
	"strings"
)

func NewRC(v string) (Part, error) {
	if v == "" {
		return &RC{
			value: 0,
		}, nil
	}
	s := strings.Split(v, "rc")
	i, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, err
	}
	return &RC{
		value: i,
	}, nil
}

type RC struct {
	value int
}

func (r *RC) Set(v int) {
	r.value = v
}

func (r *RC) Inc(v int) {
	r.value += v
}

func (r *RC) Dec(v int) error {
	if r.value-v < 0 {
		return ErrInvalidVersion
	}
	r.value -= v
	return nil
}

func (r *RC) Value() int {
	return r.value
}

func (r *RC) Clear() {
	r.value = 0
}

func (r *RC) String() string {
	return fmt.Sprintf("rc%d", r.value)
}
