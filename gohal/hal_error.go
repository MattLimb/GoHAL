// Package gohal/hal_error - Error handles for HAL.
package gohal

import "errors"

// HalError is a struct which includes everything necessary for HAL to specify an error.
type HalError struct {
	mustEnd bool
	lineNum int
	err     error
}

// Error function to ensure compatibility with the Error interface
func (e HalError) Error() string {
	return e.err.Error()
}

// newHalError creates a HalError which HAL can ignore.
func newHalError(errString string, lineNumber int) *HalError {
	return &HalError{mustEnd: false, lineNum: lineNumber, err: errors.New(errString)}
}

// newCriticalHalErro creates a HalError which HAL cannot ignore.
func newCriticalHalError(errString string, lineNumber int) *HalError {
	return &HalError{mustEnd: true, lineNum: lineNumber, err: errors.New(errString)}
}
