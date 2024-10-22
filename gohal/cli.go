// Package gohal/cli - Creates and handles the CLI input.
package gohal

import (
	"flag"
	"fmt"

	internal "github.com/MattLimb/GoHAL/gohal/internal"
)

// RunOptions is a struct which exposes all options availiable in the CLI.
type RunOptions struct {
	fileName    string
	debugMode   bool
	showVersion bool
}

// parseCli is the function to setup and parse through the CLI.
func parseCli() (RunOptions, *internal.HalError) {
	flag.Usage = func() {
		fmt.Println("Usage: gohal [flags] [filename]\n\nArguments:\n  filename  The script file you want HAL to execute.\n\nFlags:")
		flag.PrintDefaults()
	}

	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, "display the current version and exit.")
	flag.BoolVar(&versionFlag, "v", false, "display the current version and exit.")

	var debugFlag bool
	flag.BoolVar(&debugFlag, "debug", false, "Enable debug mode.")
	flag.BoolVar(&debugFlag, "d", false, "Enable debug mode.")

	flag.Parse()

	var filename string
	args := flag.Args()

	switch len(args) {
	case 0:
		filename = ""
	case 1:
		filename = args[0]
	default:
		return RunOptions{}, internal.NewCriticalHalError("too many files to process. Only 1 is expected.", 0)
	}

	return RunOptions{fileName: filename, debugMode: debugFlag, showVersion: versionFlag}, nil
}
