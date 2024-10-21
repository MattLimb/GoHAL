// Package gohal/display - A basic way to output Errors and Values to the Terminal
// This functionality is a way to ensure that the appropriate error text is displayed, and that debugMode is handled appropriately.
package gohal

import (
    "fmt"
    "os"
)

// HalDisplay is a simple struct which governs how HAL will output data and errors.
type HalDisplay struct {
	debugMode bool
}

// displayError is a function will allows HalDisplay to appropriately emit errors.
// when HalDisplay.debugMode = true it outputs additional information about the error. Primarilly where and what the error is.
func  (hd HalDisplay) displayError(err *HalError) {
    if hd.debugMode {
        fmt.Printf("I'm sorry Dave, I'm afraid I can't do that.\n  -> Line: %d\n  -> HAL %s\n", err.lineNum, err.Error())
    } else {
        fmt.Println("I'm sorry Dave, I'm afraid I can't do that.")
    }

    if err.mustEnd {
        os.Exit(1)
    }
}

// displayCharInt is a function which converts integer (int32) into a Unicode character - then prints it to the terminal.
// This function does not emit a newline character after printing its input to the terminal.
func (hd HalDisplay) displayCharInt(charInt int32) {
	fmt.Printf("%s", string(charInt))
}
