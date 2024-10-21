package gohal

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(fileName string) ([]string, *HalError) {
    var output []string

	file, err := os.Open(fileName)
    if err != nil {
		return output, NewCriticalHalError(fmt.Sprintf("cannot open specified program file %q", fileName), 0)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        output = append(output, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
		return output, NewCriticalHalError(fmt.Sprintf("cannot read from program file %s", fileName), 0)
    }

	return output, nil
}