package lang_2001

import (
	"fmt"
	"os"

	internal "github.com/MattLimb/GoHAL/internal"
)

// HalDisplay is a simple struct which governs how HAL will output data and errors.
type Lang2001Display struct {
	debugMode bool
}

// DisplayError is a function will allows HalDisplay to appropriately emit errors.
// when HalDisplay.debugMode = true it outputs additional information about the error. Primarilly where and what the error is.
func (ld Lang2001Display) DisplayError(err *internal.HalError) {
	if ld.debugMode {
		fmt.Printf("I'm sorry Dave, I'm afraid I can't do that.\n  -> Line: %d\n  -> HAL %s\n", err.LineNum, err.Error())
	} else {
		fmt.Println("I'm sorry Dave, I'm afraid I can't do that.")
	}

	if err.MustEnd {
		os.Exit(1)
	}
}

// DisplayCharInt is a function which converts integer (int32) into a Unicode character - then prints it to the terminal.
// This function does not emit a newline character after printing its input to the terminal.
func (ld Lang2001Display) DisplayCharInt(charInt int32) {
	fmt.Printf("%s", string(charInt))
}
