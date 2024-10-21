package gohal

import (
	"fmt"
	"os"
)

const Version string = "v0.1.0"

func RunHal() {
	cliArgs, err := parseCli()

	if err != nil {
		HalDisplay{debugMode: true}.displayError(err)
	}

	if cliArgs.ShowVersion {
		fmt.Printf("GoHAL %s\n", Version)
		os.Exit(0)
	}

	display := HalDisplay{debugMode: cliArgs.DebugMode}

	if cliArgs.FileName == "" {
		os.Exit(0)
	}

	inputFile, err := parseFile(cliArgs.FileName)

	if err != nil {
		display.displayError(err)
		os.Exit(1)
	}

	ast, err := BuildAst(inputFile)

	if err != nil {
		display.displayError(err)
		os.Exit(1)
	}

	interpretAst(ast, display)
}