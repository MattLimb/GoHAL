package lang_brainfuck

import (
	"github.com/MattLimb/GoHAL/internal"
)

type LangBrainfuck struct {
	runOpts internal.RunOptions
}

func New(runOpts internal.RunOptions) LangBrainfuck {
	return LangBrainfuck{runOpts: runOpts}
}

func (l LangBrainfuck) ParseFile() (internal.Ast, *internal.HalError) {
	rawInput, err := parseFile(l.runOpts.FileName)
	if err != nil {
		return internal.Ast{}, err
	}

	return parseBrainfuckCode(rawInput)
}

func (l LangBrainfuck) Display() internal.Displayer {
	return internal.BasicDisplay{DebugMode: l.runOpts.DebugMode}
}
