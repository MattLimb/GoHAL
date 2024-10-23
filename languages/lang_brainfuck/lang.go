package lang_brainfuck

import (
	"fmt"
	"os"

	"github.com/MattLimb/GoHAL/internal"
)

type LangBrainfuck struct {
	langOpts internal.LanguageOptions
}

func New(langOpts internal.LanguageOptions) LangBrainfuck {
	return LangBrainfuck{langOpts: langOpts}
}

func (l LangBrainfuck) ParseFile(fileName string) (internal.Ast, *internal.HalError) {
	rawInput, err := parseFile(fileName)
	if err != nil {
		return internal.Ast{}, err
	}

	return parseBrainfuckCode(rawInput)
}

func (l LangBrainfuck) Display() internal.Displayer {
	return internal.BasicDisplay{DebugMode: l.langOpts.DebugMode}
}

func (l LangBrainfuck) CompileToFile(ast internal.Ast, outputFileName string) *internal.HalError {
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
