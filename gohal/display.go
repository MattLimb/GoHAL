package gohal

import "fmt"

type HalDisplay struct {
	LogLevel int
}

func  (hd HalDisplay) DisplayError(err error) {
	if hd.LogLevel == 0 {
		fmt.Println("I'm sorry Dave, I'm afraid I can't do that.")
	} else {
		fmt.Printf("I'm sorry Dave, I'm afraid I can't do that. (%s)\n", err.Error())
	}
}

func (hd HalDisplay) DisplayCharInt(charInt int32) {
	fmt.Printf("%s", string(charInt))
}
