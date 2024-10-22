// Package gohal/run - Entrypoint for the GoHAL program.
package gohal

import (
	"fmt"
	"os"

	gohal_internal "github.com/MattLimb/GoHAL/gohal/internal"
	lang_2001 "github.com/MattLimb/GoHAL/gohal/languages/lang2001"
)

const Version string = "v1.1.0"

var display = gohal_internal.HalDisplay{DebugMode: true}

// RunHal is the public interface to run HAL. This requires RunOptions to be passed in.
func RunHal(runOpts RunOptions) {
	if runOpts.showVersion {
		fmt.Printf("GoHAL %s\n", Version)
		os.Exit(0)
	}

	display.DebugMode = runOpts.debugMode
	if runOpts.fileName == "" {
		os.Exit(0)
	}

	inputFile, err := lang_2001.ParseFile(runOpts.fileName)
	if err != nil {
		display.DisplayError(err)
		os.Exit(1)
	}

	ast, err := lang_2001.Parse2001Code(inputFile)
	if err != nil {
		display.DisplayError(err)
		os.Exit(1)
	}

	interpretAst(ast, map[int]int32{}, display)
}

// RunHalBinary is the public interface to run HAL. This parses the CLI and then runs hal.
func RunHalBinary() {
	runOpts, err := parseCli()
	if err != nil {
		display.DisplayError(err)
	}

	RunHal(runOpts)
}
