// Package lang_2001/read_file - A way to appropriately read a file and present it in the way HAL expects it.
package lang_brainalpha

import (
	"fmt"
	"os"

	internal "github.com/MattLimb/GoHAL/internal"
)

// parseFile is a function which reads in a file line by line.
func parseFile(fileName string) (string, *internal.HalError) {
	output, err := os.ReadFile(fileName)

	if err != nil {
		return "", internal.NewCriticalHalError(fmt.Sprintf("cannot read from program file %s", fileName), 0)
	}

	return string(output), nil
}
