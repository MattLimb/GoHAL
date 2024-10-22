// Package gohal/run - Entrypoint for the GoHAL program.
package main

import (
	"fmt"
	"os"

	"github.com/MattLimb/GoHAL/internal"
	lang_2001 "github.com/MattLimb/GoHAL/languages/lang2001"
	"github.com/MattLimb/GoHAL/languages/lang_brainfuck"
)

const Version string = "v1.1.0"

var display internal.Displayer = internal.BasicDisplay{DebugMode: true}

// RunHal is the public interface to run HAL. This requires RunOptions to be passed in.
func RunHal(runOpts internal.RunOptions) {
	if runOpts.ShowVersion {
		fmt.Printf("GoHAL %s\n", Version)
		os.Exit(0)
	}

	lang := lang_2001.New(runOpts)
	display = lang.Display()

	ast, err := lang.ParseFile()
	if err != nil {
		display.DisplayError(err)
	}

	internal.InterpretAst(ast, map[int]int32{}, display)
}

// main is the public interface to run HAL. This parses the CLI and then runs hal.
func main() {
	runOpts, err := internal.ParseCli()
	if err != nil {
		display.DisplayError(err)
	}

	if runOpts.ShowVersion {
		fmt.Printf("GoHAL %s\n", Version)
		os.Exit(0)
	}

	var lang internal.Languager

	switch runOpts.Language {
	case "lang2001":
		lang = lang_2001.New(runOpts)
	case "brainfuck":
		lang = lang_brainfuck.New(runOpts)
	default:
		display.DisplayError(internal.NewCriticalHalError(
			fmt.Sprintf("unknown language: %q", runOpts.Language),
			0,
		))
	}

	ast, err := lang.ParseFile()
	if err != nil {
		display.DisplayError(err)
	}

	internal.InterpretAst(ast, map[int]int32{}, display)
}
