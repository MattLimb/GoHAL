// Package internal/display - A basic way to output Errors and Values to the Terminal
// This functionality is a way to ensure that the appropriate error text is displayed, and that debugMode is handled appropriately.
package internal

import (
	"fmt"
	"os"
)

// Displayer is an interface for Displays HAL supports. Mostly used for Testing
type Displayer interface {
	DisplayError(err *HalError)
	DisplayCharInt(charInt int32)
}

// Basic Display
type BasicDisplay struct {
	DebugMode bool
}

// DisplayError is a function will allows HalDisplay to appropriately emit errors.
// when HalDisplay.debugMode = true it outputs additional information about the error. Primarilly where and what the error is.
func (bd BasicDisplay) DisplayError(err *HalError) {
	fmt.Printf("ERROR:\n  -> Line: %d\n  -> HAL %s\n", err.LineNum, err.Error())

	if err.MustEnd {
		os.Exit(1)
	}
}

// DisplayCharInt is a function which converts integer (int32) into a Unicode character - then prints it to the terminal.
// This function does not emit a newline character after printing its input to the terminal.
func (bd BasicDisplay) DisplayCharInt(charInt int32) {
	fmt.Printf("%s", string(charInt))
}
