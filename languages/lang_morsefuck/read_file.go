// Package read_file - A way to appropriately read a file and present it in the way HAL expects it.
package lang_morsefuck

import (
	"fmt"
	"os"
	"strings"

	internal "github.com/MattLimb/GoHAL/internal"
)

// parseFile is a function which reads in a file and chunks it.
func parseFile(fileName string) ([]string, *internal.HalError) {
	outputBytes, err := os.ReadFile(fileName)

	if err != nil {
		return []string{}, internal.NewCriticalHalError(fmt.Sprintf("cannot read from program file %s", fileName), 0)
	}

	outputStr := string(outputBytes)
	// Remove Carriage Returns and New Lines
	outputStr = strings.ReplaceAll(outputStr, "\n", "")
	outputStr = strings.ReplaceAll(outputStr, "\r", "")

	if len(outputStr)%3 != 0 {
		return []string{}, internal.NewCriticalHalError("input is not long enough to ensure proper decoding.", 0)
	}

	var output = make([]string, len(outputStr)/3)
	idx := 0

	for _, char := range outputStr {
		output[idx] += string(char)

		if len(output[idx]) == 3 {
			idx++
		}
	}

	return output, nil
}
