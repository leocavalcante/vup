package vup

import (
	"fmt"
	"regexp"
	"strconv"
)

// rcPattern matches a canonical rc label: "rc" plus a nonnegative integer
// without leading zeros. rc0 is valid.
var rcPattern = regexp.MustCompile(`^rc(0|[1-9][0-9]*)$`)

func NewRC(v string) (Part, error) {
	if v == "" {
		return &RC{}, nil
	}

	match := rcPattern.FindStringSubmatch(v)
	if match == nil {
		return nil, fmt.Errorf("invalid rc string: %q", v)
	}

	i, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, err
	}

	return &RC{
		value:   i,
		present: true,
	}, nil
}

type RC struct {
	value   int
	present bool
}

func (r *RC) Set(v int) {
	r.value = v
	r.present = true
}

func (r *RC) Inc(v int) {
	r.value += v
	r.present = true
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

func (r *RC) Present() bool {
	return r.present
}

func (r *RC) Clear() {
	r.value = 0
	r.present = false
}

func (r *RC) String() string {
	return fmt.Sprintf("rc%d", r.value)
}
