// Package lang_2001/read_file - A way to appropriately read a file and present it in the way HAL expects it.
package lang_2001

import (
	"bufio"
	"fmt"
	"os"

	internal "github.com/MattLimb/GoHAL/internal"
)

// parseFile is a function which reads in a file line by line.
func parseFile(fileName string) ([]string, *internal.HalError) {
	var output []string

	file, err := os.Open(fileName)
	if err != nil {
		return output, internal.NewCriticalHalError(fmt.Sprintf("cannot open specified program file %q", fileName), 0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		text := scanner.Text()

		if text != "" {
			output = append(output, text)
		}
	}

	if err := scanner.Err(); err != nil {
		return output, internal.NewCriticalHalError(fmt.Sprintf("cannot read from program file %s", fileName), 0)
	}

	return output, nil
}
