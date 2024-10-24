package lang_morsefuck

import (
	"github.com/MattLimb/GoHAL/internal"
)

func compileAst(ast internal.Ast) (string, *internal.HalError) {
	outputAst := ""

	for _, instruction := range ast {
		switch instruction.Instruction {
		case internal.IncrementCell:
			str := ""

			for range instruction.N {
				str += "..-"
			}

			outputAst += str
		case internal.DecrementCell:
			str := ""

			for range instruction.N {
				str += "-.."
			}

			outputAst += str
		case internal.LoopStart:
			outputAst += "---"
		case internal.LoopEnd:
			outputAst += "..."
		case internal.ShiftLeft:
			outputAst += "--."
		case internal.ShiftRight:
			outputAst += ".--"
		case internal.UserInput:
			outputAst += ".-."
		case internal.DisplayChar:
			outputAst += "-.-"
		// Unproducable
		case internal.LoopBreak:
			return "", internal.NewCriticalHalError("morsefuck does not support Breaking Loops", 0)
		case internal.LoopBreakAll:
			return "", internal.NewCriticalHalError("morsefuck does not support Breaking Loops", 0)
		}
	}

	return outputAst, nil
}
