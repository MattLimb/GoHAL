// Package internal/error - Error handles for HAL.
package internal

import "errors"

// HalError is a struct which includes everything necessary for HAL to specify an error.
type HalError struct {
	MustEnd bool
	LineNum int
	Err     error
}

// Error function to ensure compatibility with the Error interface
func (e HalError) Error() string {
	return e.Err.Error()
}

// NewHalError creates a HalError which HAL can ignore.
func NewHalError(errString string, lineNumber int) *HalError {
	return &HalError{MustEnd: false, LineNum: lineNumber, Err: errors.New(errString)}
}

// NewCriticalHalError creates a HalError which HAL cannot ignore.
func NewCriticalHalError(errString string, lineNumber int) *HalError {
	return &HalError{MustEnd: true, LineNum: lineNumber, Err: errors.New(errString)}
}
