package lang_brainalpha

import (
	"fmt"

	"github.com/MattLimb/GoHAL/internal"
)

func parseBrainfuckCode(rawInput string) (internal.Ast, *internal.HalError) {
	var charNumber int = 1

	ast := []internal.Node{
		{Instruction: internal.ProgramStart, N: 0, LoopStart: 0, LoopEnd: 0},
	}
	loopStartIndexes := map[int]int{}
	currentLoopId := 0

	for _, instruction := range rawInput {
		charNumber++

		switch {
		// Pointer Shift Operations
		case instruction == 'A':
			ast = append(ast, internal.Node{Instruction: internal.ShiftRight, N: 0, LoopStart: 0, LoopEnd: 0})
		case instruction == 'B':
			ast = append(ast, internal.Node{Instruction: internal.ShiftLeft, N: 0, LoopStart: 0, LoopEnd: 0})

		// Increment & Decrement Operators
		case instruction == 'C':
			if lastInstruction := ast[len(ast)-1]; lastInstruction.Instruction == internal.IncrementCell {
				ast[len(ast)-1].N += 1
			} else {
				ast = append(ast, internal.Node{Instruction: internal.IncrementCell, N: 1, LoopStart: 0, LoopEnd: 0})
			}
		case instruction == 'D':
			if lastInstruction := ast[len(ast)-1]; lastInstruction.Instruction == internal.DecrementCell {
				ast[len(ast)-1].N += 1
			} else {
				ast = append(ast, internal.Node{Instruction: internal.DecrementCell, N: 1, LoopStart: 0, LoopEnd: 0})
			}
		// User and Print Output
		case instruction == 'E':
			ast = append(ast, internal.Node{Instruction: internal.DisplayChar, N: 0, LoopStart: 0, LoopEnd: 0})
		case instruction == 'F':
			ast = append(ast, internal.Node{Instruction: internal.UserInput, N: 0, LoopStart: 0, LoopEnd: 0})

		// Looping Section
		case instruction == 'G':
			currentLoopId++
			ast = append(ast, internal.Node{Instruction: internal.LoopStart, N: 0, LoopStart: 0, LoopEnd: 0})

			loopStartIndexes[currentLoopId] = len(ast) - 1
		case instruction == 'H':
			if currentLoopId == 0 {
				return []internal.Node{}, internal.NewCriticalHalError("program cannot end a loop without starting one", charNumber)
			}

			loopStartIdx := loopStartIndexes[currentLoopId]

			ast = append(ast, internal.Node{Instruction: internal.LoopEnd, N: 0, LoopStart: loopStartIdx, LoopEnd: 0})
			ast[loopStartIdx].LoopEnd = len(ast) - 1

			currentLoopId--
		case instruction == '\n' || instruction == ' ' || instruction == '\r': // Do Nothing
		// If nothing matches - HAL does not understand the instruction. Bail out and tell the user.
		default:
			return []internal.Node{}, internal.NewCriticalHalError(fmt.Sprintf("unrecognised Instruction: %q", instruction), charNumber)
		}
	}

	// If we get to the end of a program - and we haven't closed a loop - bail out and tell the user.
	if currentLoopId != 0 {
		return []internal.Node{}, internal.NewCriticalHalError("loop was not closed", loopStartIndexes[currentLoopId]+1)
	}

	ast = append(ast, internal.Node{Instruction: internal.ProgramEnd, N: 0, LoopStart: 0, LoopEnd: 0})

	return ast, nil
}
