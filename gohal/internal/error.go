// Package gohal_internal/error - Error handles for HAL.
package gohal_internal

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

// NewHalError creates a HalError which HAL can ignore.
func NewHalError(errString string, lineNumber int) *HalError {
	return &HalError{mustEnd: false, lineNum: lineNumber, err: errors.New(errString)}
}

// NewCriticalHalError creates a HalError which HAL cannot ignore.
func NewCriticalHalError(errString string, lineNumber int) *HalError {
	return &HalError{mustEnd: true, lineNum: lineNumber, err: errors.New(errString)}
}
