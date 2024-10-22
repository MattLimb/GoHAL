package lang_2001

import (
	"fmt"
	"regexp"
	"strings"

	internal "github.com/MattLimb/GoHAL/internal"
)

// parse2001Code is the function which parses each line of a file, and processes it into an HalNode.
// This function will return MustExit erros when it reaches a part of the syntax it cannot continue with.
func parse2001Code(fileLines []string) (internal.Ast, *internal.HalError) {
	var lineNumber int
	programLength := len(fileLines)

	if programLength < 2 {
		return []internal.Node{}, internal.NewCriticalHalError("program is too short. It must be at least 2 lines long.", lineNumber)
	}

	ast := make([]internal.Node, programLength)
	loopStartIndexes := map[int]int{}
	currentLoopId := 0

	startprogramRegex := regexp.MustCompile(`Good afternoon, gentlemen. I am a (.*) computer\. I became operational at (.*) on (.*).`)

	for idx, instruction := range fileLines {
		lineNumber = idx + 1

		switch {
		// The Program MUST have a specified start phrase.
		case idx == 0:
			if startprogramRegex.MatchString(instruction) {
				ast[idx] = internal.Node{Instruction: internal.ProgramStart, N: 0, LoopStart: 0, LoopEnd: 0}
			} else {
				return []internal.Node{}, internal.NewCriticalHalError("program must start with 'Good afternoon, gentlemen.' command", lineNumber)
			}
		// The Program MUST end with speficied Phrase. Error if it doesn't.
		case idx == (programLength-1) && instruction != "Stop, Dave.":
			return []internal.Node{}, internal.NewCriticalHalError("program must end with 'Stop, Dave.' command", lineNumber)
		// End the Program
		case instruction == "Stop, Dave.":
			ast[idx] = internal.Node{Instruction: internal.ProgramEnd, N: 0, LoopStart: 0, LoopEnd: 0}

		// Increment & Decrement Operators
		case strings.Contains(instruction, "Hal?"):
			ast[idx] = internal.Node{Instruction: internal.IncrementCell, N: int32(strings.Count(instruction, "Hal!")), LoopStart: 0, LoopEnd: 0}
		case strings.Contains(instruction, "I'm afraid. I'm afraid, Dave. Dave, my mind is going."):
			ast[idx] = internal.Node{Instruction: internal.DecrementCell, N: int32(strings.Count(instruction, "I can feel it.")), LoopStart: 0, LoopEnd: 0}

		// Looping Section
		case instruction == "What are you doing, Dave?":
			currentLoopId++
			loopStartIndexes[currentLoopId] = idx

			ast[idx] = internal.Node{Instruction: internal.LoopStart, N: 0, LoopStart: 0, LoopEnd: 0}
		case instruction == "Dave, this conversation can serve no purpose anymore. Goodbye.":
			if currentLoopId == 0 {
				return []internal.Node{}, internal.NewCriticalHalError("program cannot end a loop without starting one", lineNumber)
			}

			loopStartIdx := loopStartIndexes[currentLoopId]

			ast[idx] = internal.Node{Instruction: internal.LoopEnd, N: 0, LoopStart: loopStartIdx, LoopEnd: 0}
			ast[loopStartIdx].LoopEnd = idx

			currentLoopId--

		// Loop Break. This will break out of 2, 1 or 0 loops depending on loop depth.
		case instruction == "This mission is too important for me to allow you to jeopardize it.":
			ast[idx] = internal.Node{Instruction: internal.LoopBreak, N: 0, LoopStart: 0, LoopEnd: 0}
		// Break out of ALL loops. Doesn't matter the depth.
		case instruction == "I know I've made some very poor decisions recently, but I can give you my complete assurance that my work will be back to normal.":
			ast[idx] = internal.Node{Instruction: internal.LoopBreakAll, N: 0, LoopStart: 0, LoopEnd: 0}

		// Pointer Shift Operations
		case instruction == "I've picked up a fault in the AE-35 unit.":
			ast[idx] = internal.Node{Instruction: internal.ShiftLeft, N: 0, LoopStart: 0, LoopEnd: 0}
		case instruction == "Well, he acts like he has genuine emotions.":
			ast[idx] = internal.Node{Instruction: internal.ShiftRight, N: 0, LoopStart: 0, LoopEnd: 0}

		// User and Print Output
		case instruction == "Open the pod bay doors, HAL.":
			ast[idx] = internal.Node{Instruction: internal.UserInput, N: 0, LoopStart: 0, LoopEnd: 0}
		case instruction == "Close the pod bay doors, HAL.":
			ast[idx] = internal.Node{Instruction: internal.DisplayChar, N: 0, LoopStart: 0, LoopEnd: 0}

		// If nothing matches - HAL does not understand the instruction. Bail out and tell the user.
		default:
			return []internal.Node{}, internal.NewCriticalHalError(fmt.Sprintf("unrecognised Instruction: %q", instruction), lineNumber)
		}
	}

	// If we get to the end of a program - and we haven't closed a loop - bail out and tell the user.
	if currentLoopId != 0 {
		return []internal.Node{}, internal.NewCriticalHalError("loop was not closed", loopStartIndexes[currentLoopId]+1)
	}

	return ast, nil
}
