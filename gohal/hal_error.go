package gohal

import "errors"

type HalError struct {
	mustEnd bool
	lineNum int
	err     error
}

func (e HalError) Error() string {
	return e.err.Error()
}

func newHalError(errString string, lineNumber int) *HalError {
	return &HalError{mustEnd: false, lineNum: lineNumber, err: errors.New(errString)}
}

func newCriticalHalError(errString string, lineNumber int) *HalError {
	return &HalError{mustEnd: true, lineNum: lineNumber, err: errors.New(errString)}
}