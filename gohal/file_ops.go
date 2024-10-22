// Package gohal/file_ops - A way to appropriately read a file and present it in the way HAL expects it.
package gohal

import (
	"bufio"
	"fmt"
	"os"
)

// parseFile is a function which reads in a file line by line.
func parseFile(fileName string) ([]string, *HalError) {
	var output []string

	file, err := os.Open(fileName)
	if err != nil {
		return output, newCriticalHalError(fmt.Sprintf("cannot open specified program file %q", fileName), 0)
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
		return output, newCriticalHalError(fmt.Sprintf("cannot read from program file %s", fileName), 0)
	}

	return output, nil
}
