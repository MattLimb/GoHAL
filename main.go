// Package main - Entrypoint for the GoHAL program.
package main

import (
	"fmt"
	"os"

	"github.com/MattLimb/GoHAL/internal"
	"github.com/MattLimb/GoHAL/languages/lang_2001"
	"github.com/MattLimb/GoHAL/languages/lang_brainalpha"
	"github.com/MattLimb/GoHAL/languages/lang_brainfuck"
	"github.com/MattLimb/GoHAL/languages/lang_morsefuck"
)

const Version string = "v1.5.0"

var display internal.Displayer = internal.BasicDisplay{DebugMode: true}

// language identifies and sets up a known language. Errors if one cannot be found.
func language(langStr string, langOpts internal.LanguageOptions) (internal.Languager, *internal.HalError) {
	switch langStr {
	case "2001":
		return lang_2001.New(langOpts), nil
	case "brainfuck":
		return lang_brainfuck.New(langOpts), nil
	case "morsefuck":
		return lang_morsefuck.New(langOpts), nil
	case "brainalpha":
		return lang_brainalpha.New(langOpts), nil
	default:
		return nil, internal.NewCriticalHalError(fmt.Sprintf("unrecognised language: %q", langStr), 0)
	}
}

// runCommand takes an input file and tries to execute it as the given language.
func runCommand(runOpts internal.ProgOptions) *internal.HalError {
	lang, err := language(runOpts.Language.Input, runOpts.LangOptions)
	if err != nil {
		return err
	}

	ast, err := lang.ParseFile(runOpts.Files.Input)
	if err != nil {
		return err
	}

	// err = internal.InterpretAs t(ast, map[int]int32{}, lang.Display())
	err = internal.InterpretAst(ast, lang.Tape(), lang.Display())

	if err != nil {
		display.DisplayError(err)
	}

	return nil
}

// transpileCommand takes a file in one language and tries to render it in another.
func transpileCommand(runOpts internal.ProgOptions) *internal.HalError {
	inputLang, err := language(runOpts.Language.Input, runOpts.LangOptions)
	if err != nil {
		return err
	}

	outputLang, err := language(runOpts.Language.Output, runOpts.LangOptions)
	if err != nil {
		return err
	}

	progAst, err := inputLang.ParseFile(runOpts.Files.Input)
	if err != nil {
		return err
	}

	err = outputLang.CompileToFile(progAst, runOpts.Files.Output)
	if err != nil {
		return err
	}

	return nil
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

	switch runOpts.Command {
	case "run":
		err = runCommand(runOpts)
	case "transpile":
		err = transpileCommand(runOpts)
	}

	if err != nil {
		display.DisplayError(err)
	}
}
