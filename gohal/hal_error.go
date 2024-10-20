package gohal

import "errors"

type HalError struct {
	MustEnd bool
	Err     error
}

func (e HalError) Error() string {
	return e.Err.Error()
}

func NewHalError(errString string) HalError {
	return HalError{MustEnd: false, Err: errors.New(errString)}
}

func NewCriticalHalError(errString string) HalError {
	return HalError{MustEnd: true, Err: errors.New(errString)}
}