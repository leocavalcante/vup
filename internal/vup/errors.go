package vup

import "errors"

var (
	ErrInvalidSemanticVersion = errors.New("invalid semantic version string")
	ErrMajorLessThanZero      = errors.New("major version must be greater than or equal to zero")
	ErrMinorLessThanOne       = errors.New("minor version must be greater than or equal to one")
	ErrPatchLessThanZero      = errors.New("patch version must be greater than or equal to zero")
)
