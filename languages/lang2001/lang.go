package lang_2001

import (
	"github.com/MattLimb/GoHAL/internal"
)

type Lang2001 struct {
	runOpts internal.RunOptions
}

func New(runOpts internal.RunOptions) Lang2001 {
	return Lang2001{runOpts: runOpts}
}

// type Languager interface {
// 	ParseFile(fileName string) (Ast, HalError)
// 	ParseString(raw string) (Ast, HalError)
// 	Display() Displayer
// }

func (l Lang2001) ParseFile() (internal.Ast, *internal.HalError) {
	file_output, err := parseFile(l.runOpts.FileName)

	if err != nil {
		return internal.Ast{}, err
	}

	return parse2001Code(file_output)
}

func (l Lang2001) Display() internal.Displayer {
	return Lang2001Display{debugMode: l.runOpts.DebugMode}
}
