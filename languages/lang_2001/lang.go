package lang_2001

import (
	"fmt"
	"os"

	"github.com/MattLimb/GoHAL/internal"
)

type Lang2001 struct {
	langOpts internal.LanguageOptions
}

func New(langOpts internal.LanguageOptions) Lang2001 {
	return Lang2001{langOpts: langOpts}
}

func (l Lang2001) ParseFile(fileName string) (internal.Ast, *internal.HalError) {
	file_output, err := parseFile(fileName)

	if err != nil {
		return internal.Ast{}, err
	}

	return parse2001Code(file_output)
}

func (l Lang2001) Display() internal.Displayer {
	return Lang2001Display{debugMode: l.langOpts.DebugMode}
}

func (l Lang2001) CompileToFile(ast internal.Ast, outputFileName string) *internal.HalError {
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

func (l Lang2001) Tape() internal.Taper {
	return internal.NewDefaultTape()
}
