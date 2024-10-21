package gohal

import (
    "fmt"
    "os"
)

type HalDisplay struct {
	debugMode bool
}

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

func (hd HalDisplay) displayCharInt(charInt int32) {
	fmt.Printf("%s", string(charInt))
}
