// Package internal/cli - Creates and handles the CLI input.
package internal

import (
	"flag"
	"fmt"
)

// RunOptions is a struct which exposes all options availiable in the CLI.
type RunOptions struct {
	FileName    string
	DebugMode   bool
	ShowVersion bool
	Language    string
}

// ParseCli is the function to setup and parse through the CLI.
func ParseCli() (RunOptions, *HalError) {
	flag.Usage = func() {
		fmt.Println("Usage: gohal [flags] [filename]\n\nArguments:\n  filename  The script file you want HAL to execute.\n\nFlags:")
		flag.PrintDefaults()
	}

	var language string
	flag.StringVar(&language, "language", "2001", "specify which language to try to parse and run.")
	flag.StringVar(&language, "l", "2001", "specify which language to try to parse and run.")

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
		return RunOptions{}, NewCriticalHalError("too many files to process. Only 1 is expected.", 0)
	}

	return RunOptions{FileName: filename, DebugMode: debugFlag, ShowVersion: versionFlag, Language: language}, nil
}
