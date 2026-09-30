package vup

import "errors"

var (
	ErrInvalidSemanticVersion = errors.New("invalid semantic version string")
	ErrMajorLessThanZero      = errors.New("major version must be greater than or equal to zero")
	ErrMinorLessThanOne       = errors.New("minor version must be greater than or equal to one")
	ErrMinorLessThanZero      = errors.New("minor version must be greater than or equal to zero")
	ErrPatchLessThanZero      = errors.New("patch version must be greater than or equal to zero")
	ErrInvalidVersion         = errors.New("invalid version")
	ErrRCLessThanZero         = errors.New("rc version must be greater than or equal to zero")
)
