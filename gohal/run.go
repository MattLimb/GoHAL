package gohal

import (
	"fmt"
	"os"
)

const Version string = "v1.1.0"

func RunHal() {
	cliArgs, err := parseCli()
	if err != nil {
		HalDisplay{debugMode: true}.displayError(err)
	}

	if cliArgs.showVersion {
		fmt.Printf("GoHAL %s\n", Version)
		os.Exit(0)
	}

	display := HalDisplay{debugMode: cliArgs.debugMode}
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