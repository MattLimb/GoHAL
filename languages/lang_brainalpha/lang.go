package lang_brainalpha

import (
	"fmt"
	"os"

	"github.com/MattLimb/GoHAL/internal"
)

type LangBrainalpha struct {
	langOpts internal.LanguageOptions
}

func New(langOpts internal.LanguageOptions) LangBrainalpha {
	return LangBrainalpha{langOpts: langOpts}
}

func (l LangBrainalpha) ParseFile(fileName string) (internal.Ast, *internal.HalError) {
	rawInput, err := parseFile(fileName)
	if err != nil {
		return internal.Ast{}, err
	}

	return parseBrainfuckCode(rawInput)
}

func (l LangBrainalpha) Display() internal.Displayer {
	return internal.BasicDisplay{DebugMode: l.langOpts.DebugMode}
}

func (l LangBrainalpha) CompileToFile(ast internal.Ast, outputFileName string) *internal.HalError {
	compiled, err := compileAst(ast)
	if err != nil {
		return err
	}

	f, osErr := os.Create(outputFileName)
	if osErr != nil {
		return internal.NewCriticalHalError(fmt.Sprintf("cannot create output file: %s", osErr.Error()), 0)
	}

	defer f.Close()

	_, osErr = f.WriteString(compiled)
	if osErr != nil {
		return internal.NewCriticalHalError(fmt.Sprintf("cannot write to output file: %s", osErr.Error()), 0)
	}

	return nil
}

func (l LangBrainalpha) Tape() internal.Taper {
	return NewBrainalphaTape()
}
