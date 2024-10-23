package lang_2001

import (
	"strings"

	"github.com/MattLimb/GoHAL/internal"
)

func compileAst(ast internal.Ast) (string, *internal.HalError) {
	outputAst := make([]string, len(ast))

	for idx, instruction := range ast {
		switch instruction.Instruction {
		case internal.ProgramStart:
			outputAst[idx] = "Good afternoon, gentlemen. I am a <COMPUTER TYPE> computer. I became operational at <LOCATION> on <DATE>."
		case internal.IncrementCell:
			str := "Hal?"

			for range instruction.N {
				str += " Hal!"
			}

			outputAst[idx] = str
		case internal.DecrementCell:
			str := "I'm afraid. I'm afraid, Dave. Dave, my mind is going."

			for range instruction.N {
				str += " I can feel it."
			}

			outputAst[idx] = str
		case internal.LoopStart:
			outputAst[idx] = "What are you doing, Dave?"
		case internal.LoopEnd:
			outputAst[idx] = "Dave, this conversation can serve no purpose anymore. Goodbye."
		case internal.LoopBreak:
			outputAst[idx] = "This mission is too important for me to allow you to jeopardize it."
		case internal.LoopBreakAll:
			outputAst[idx] = "I know I've made some very poor decisions recently, but I can give you my complete assurance that my work will be back to normal."
		case internal.ShiftLeft:
			outputAst[idx] = "I've picked up a fault in the AE-35 unit."
		case internal.ShiftRight:
			outputAst[idx] = "Well, he acts like he has genuine emotions."
		case internal.UserInput:
			outputAst[idx] = "Open the pod bay doors, HAL."
		case internal.DisplayChar:
			outputAst[idx] = "Close the pod bay doors, HAL."
		case internal.ProgramEnd:
			outputAst[idx] = "Stop, Dave."
		}
	}

	return strings.Join(outputAst, "\n"), nil
}
