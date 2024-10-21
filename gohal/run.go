// Package gohal/run - Entrypoint for the GoHAL program.
package gohal

import (
	"fmt"
	"os"
)

const Version string = "v1.1.0"

var display HalDisplay = HalDisplay{debugMode: true}

var display = HalDisplay{debugMode: true}

// RunHal is the public interface to run HAL. This requires RunOptions to be passed in.
func RunHal(runOpts RunOptions) {
	if runOpts.showVersion {
		fmt.Printf("GoHAL %s\n", Version)
		os.Exit(0)
	}

	display.debugMode = runOpts.debugMode
	if runOpts.fileName == "" {
		os.Exit(0)
	}

	inputFile, err := parseFile(runOpts.fileName)
	if err != nil {
		display.displayError(err)
		os.Exit(1)
	}

	ast, err := buildAst(inputFile)
	if err != nil {
		display.displayError(err)
		os.Exit(1)
	}

	interpretAst(ast, map[int]int32{}, display)
}

// RunHalBinary is the public interface to run HAL. This parses the CLI and then runs hal.
func RunHalBinary() {
	runOpts, err := parseCli()
	if err != nil {
		display.displayError(err)
	}

	RunHal(runOpts)
}
