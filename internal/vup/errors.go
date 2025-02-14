package vup

import "errors"

var (
	ErrMajorLessThanZero = errors.New("major version must be greater than or equal to zero")
	ErrMinorLessThanOne  = errors.New("minor version must be greater than or equal to one")
	ErrPatchLessThanZero = errors.New("patch version must be greater than or equal to zero")
)
