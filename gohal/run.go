// Package gohal/run - Entrypoint for the GoHAL program.
package gohal

import (
	"fmt"
	"os"
)

const Version string = "v1.1.0"

// RunHal is the public interface to run HAL. This currently assumes the CLI is wanted.
func RunHal() {
	display := HalDisplay{true}

	cliArgs, err := parseCli()
	if err != nil {
		display.displayError(err)
	}

	if cliArgs.showVersion {
		fmt.Printf("GoHAL %s\n", Version)
		os.Exit(0)
	}

	display.debugMode = cliArgs.debugMode
	if cliArgs.fileName == "" {
		os.Exit(0)
	}

	inputFile, err := parseFile(cliArgs.fileName)
	if err != nil {
		display.displayError(err)
		os.Exit(1)
	}

	ast, err := buildAst(inputFile)
	if err != nil {
		display.displayError(err)
		os.Exit(1)
	}

	interpretAst(ast, display)
}