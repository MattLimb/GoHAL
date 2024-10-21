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

func NewHalError(errString string, lineNumber int) *HalError {
	return &HalError{mustEnd: false, lineNum: lineNumber, err: errors.New(errString)}
}

func NewCriticalHalError(errString string, lineNumber int) *HalError {
	return &HalError{mustEnd: true, lineNum: lineNumber, err: errors.New(errString)}
}