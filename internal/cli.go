// Package internal/cli - Creates and handles the CLI input.
package internal

import (
	"flag"
	"fmt"
)

type LanguageOptions struct {
	DebugMode bool
}

type InputOutput struct {
	Input  string
	Output string
}

// ProgOptions is a struct which exposes all options availiable in the CLI.
type ProgOptions struct {
	Command     string
	Files       InputOutput
	Language    InputOutput
	ShowVersion bool
	LangOptions LanguageOptions
}

// ParseCli is the function to setup and parse through the CLI.
func ParseCli() (ProgOptions, *HalError) {
	flag.Usage = func() {
		fmt.Println("Usage:\n  gohal -v\n  gohal [--debug,--language] run [inputFilename]\n  gohal [--debug,--language,--outputLanguage] transpile [inputFilename] [outputFilename]\n\nArguments:\n  filename  The script file you want HAL to execute.\n\nFlags:")
		flag.PrintDefaults()
	}

	var language string
	flag.StringVar(&language, "language", "2001", "specify which language to parse the input with.")
	flag.StringVar(&language, "l", "2001", "specify which language to parse the input with.")

	var outputLanguage string
	flag.StringVar(&outputLanguage, "outputLanguage", "brainfuck", "specify which language to try to parse and run.")
	flag.StringVar(&outputLanguage, "o", "brainfuck", "specify which language to try to parse and run.")

	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, "display the current version and exit.")
	flag.BoolVar(&versionFlag, "v", false, "display the current version and exit.")

	var debugFlag bool
	flag.BoolVar(&debugFlag, "debug", false, "Global flag to enable debug mode.")
	flag.BoolVar(&debugFlag, "d", false, "Global flag to enable debug mode.")

	flag.Parse()

	var inputFilename string
	var outputFilename string

	args := flag.Args()
	numArgs := len(args)

	if versionFlag {
		args = append(args, "version")
	} else if numArgs < 2 {
		return ProgOptions{}, NewCriticalHalError("too few arguments. Expected at least 2 arguments.", 0)
	}

	switch args[0] {
	case "run":
		if numArgs > 2 {
			return ProgOptions{}, NewCriticalHalError("too many arguments for 'run', Run expects exactly 1 file.", 0)
		}

		inputFilename = args[1]
	case "transpile":
		if numArgs > 3 {
			return ProgOptions{}, NewCriticalHalError("too many arguments for 'transpile', Run expects exactly 2 files.", 0)
		}

		inputFilename = args[1]
		outputFilename = args[2]
	case "version":
	default:
		return ProgOptions{}, NewCriticalHalError(fmt.Sprintf("unreccognized command %q", args[0]), 0)
	}

	return ProgOptions{
		Command:     args[0],
		Files:       InputOutput{Input: inputFilename, Output: outputFilename},
		Language:    InputOutput{Input: language, Output: outputLanguage},
		ShowVersion: versionFlag,
		LangOptions: LanguageOptions{DebugMode: debugFlag}}, nil
}
