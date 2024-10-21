package gohal

import (
    "flag"
    "fmt"
    "os"
)

type CliArgs struct {
	FileName string
	DebugMode bool
	ShowVersion bool
}


func parseCli() (CliArgs, *HalError) {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: gohal [flags] [filename]\n\nArguments:\n  filename  The script file you want HAL to execute.\n\nFlags:")
		flag.PrintDefaults()
	}

	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, "Display the current version and exit.")
	flag.BoolVar(&versionFlag, "v", false, "Display the current version and exit.")

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
		return CliArgs{}, NewCriticalHalError("too many files to process. Only 1 is expected.", 0)
	}

	return CliArgs{FileName: filename, DebugMode: debugFlag, ShowVersion: versionFlag}, nil
}