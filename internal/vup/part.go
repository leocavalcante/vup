package vup

import (
	"fmt"
	"math"
)

type Inc interface {
	Inc(int) error
}

type Dec interface {
	Dec(int) error
}

type Clear interface {
	Clear()
}

type Part interface {
	Inc
	Dec
	Clear
	Value() int
	Set(int) error
}

func increment(value, step int) (int, error) {
	if step < 0 {
		return value, fmt.Errorf("%w: increment must not be negative", ErrInvalidVersion)
	}
	if value > math.MaxInt-step {
		return value, fmt.Errorf("%w: increment overflows", ErrInvalidVersion)
	}
	return value + step, nil
}

func decrement(value, step int, belowZero error) (int, error) {
	if step < 0 {
		return value, fmt.Errorf("%w: decrement must not be negative", ErrInvalidVersion)
	}
	if step > value {
		return value, belowZero
	}
	return value - step, nil
}
