package lang_morsefuck

import (
	"fmt"
	"os"

	"github.com/MattLimb/GoHAL/internal"
)

type LangMorsefuck struct {
	langOpts internal.LanguageOptions
}

func New(langOpts internal.LanguageOptions) LangMorsefuck {
	return LangMorsefuck{langOpts: langOpts}
}

func (l LangMorsefuck) ParseFile(fileName string) (internal.Ast, *internal.HalError) {
	rawInput, err := parseFile(fileName)
	if err != nil {
		return internal.Ast{}, err
	}

	return parseMorsefuckCode(rawInput)
}

func (l LangMorsefuck) Display() internal.Displayer {
	return internal.BasicDisplay{DebugMode: l.langOpts.DebugMode}
}

func (l LangMorsefuck) CompileToFile(ast internal.Ast, outputFileName string) *internal.HalError {
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
