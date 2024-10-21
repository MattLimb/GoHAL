// Package gohal/ast - A Basic Abstract Syntax Tree for GoHAL.
// This file takes the user input file (as a []string) and parses out all variability into a standard format.
package gohal

import (
	"fmt"
    "regexp"
	"strings"
)

// HalInstruction is a descriptive type to tag its usage as a command type.
// They are used to tell HAL which instruction to run.
type HalInstruction string

const (
	incrementCell HalInstruction = "increment_cell"
	decrementCell HalInstruction = "decrement_cell"
	loopStart HalInstruction = "loop_start"
	loopEnd HalInstruction = "loop_end"
	loopBreak HalInstruction = "loop_break"
	loopBreakAll HalInstruction = "loop_break_all"
	shiftLeft HalInstruction = "shift_left"
	shiftRight HalInstruction = "shift_right"
	userInput HalInstruction = "user_input"
	displayChar HalInstruction = "display_char"
	programStart HalInstruction = "program_start"
	programEnd HalInstruction = "program_end"
)

// HalNode is a struct to encode functionality in an easier way.
// Currently supports commands with a variable number of instructions and more efficient loop processing.
type HalNode struct {
	instruction HalInstruction
	n           int32
	loopStart   int
	loopEnd		int
}

// HalAst is a descriptive type to refer to a collection of HalNodes.
type HalAst []HalNode

// buildAst is the function which parses each line of a file, and processes it into an HalNode.
// This function will return MustExit erros when it reaches a part of the syntax it cannot continue with.
func buildAst(fileLines []string) (HalAst, *HalError) {
	var lineNumber int
	programLength := len(fileLines)

	if programLength < 2 {
		return []HalNode{}, newCriticalHalError("program is too short. It must be at least 2 lines long.", lineNumber)
	}

	ast := make([]HalNode, programLength)
	loopStartIndexes := map[int]int{}
	currentLoopId := 0

	startprogramRegex := regexp.MustCompile(`Good afternoon, gentlemen. I am a (.*) computer\. I became operational at (.*) on (.*).`)

	for idx, instruction := range fileLines {
		lineNumber = idx + 1

		switch {
		// The Program MUST have a specified start phrase.
		case idx == 0:
			if startprogramRegex.MatchString(instruction) {
				ast[idx] = HalNode{instruction: programStart, n: 0, loopStart: 0, loopEnd: 0}
			} else {
				return []HalNode{}, newCriticalHalError("program must start with 'Good afternoon, gentlemen.' command", lineNumber)
			}
		// The Program MUST end with speficied Phrase. Error if it doesn't.
		case idx == (programLength-1) && instruction != "Stop, Dave.":
			return []HalNode{}, newCriticalHalError("program must end with 'Stop, Dave.' command", lineNumber)
		// End the Program
		case instruction == "Stop, Dave.":
			ast[idx] = HalNode{instruction: programEnd, n: 0, loopStart: 0, loopEnd: 0}

		// Increment & Decrement Operators
		case strings.Contains(instruction, "Hal?"):
			ast[idx] =  HalNode{instruction: incrementCell, n: int32(strings.Count(instruction, "Hal!")), loopStart: 0, loopEnd: 0}
		case strings.Contains(instruction, "I'm afraid. I'm afraid, Dave. Dave, my mind is going."):
			ast[idx] = HalNode{instruction: decrementCell, n: int32(strings.Count(instruction, "I can feel it.")), loopStart: 0, loopEnd: 0}

		// Looping Section

		case instruction ==  "What are you doing, Dave?":
			currentLoopId++
			loopStartIndexes[currentLoopId] = idx

			ast[idx] = HalNode{instruction: loopStart, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "Dave, this conversation can serve no purpose anymore. Goodbye.":
			if currentLoopId == 0 {
				return []HalNode{}, newCriticalHalError("program cannot end a loop without starting one", lineNumber)
			}

			loopStartIdx := loopStartIndexes[currentLoopId]

			ast[idx] = HalNode{instruction: loopEnd, n: 0, loopStart: loopStartIdx, loopEnd: 0}
			ast[loopStartIdx].loopEnd = idx

			currentLoopId--

		// Loop Break. This will break out of 2, 1 or 0 loops depending on loop depth.
		case instruction == "This mission is too important for me to allow you to jeopardize it.":
			ast[idx] = HalNode{instruction: loopBreak, n: 0, loopStart: 0, loopEnd: 0}
		// Break out of ALL loops. Doesn't matter the depth.
		case instruction == "I know I've made some very poor decisions recently, but I can give you my complete assurance that my work will be back to normal.":
			ast[idx] = HalNode{instruction: loopBreakAll, n: 0, loopStart: 0, loopEnd: 0}

		// Pointer Shift Operations
		case instruction == "I've picked up a fault in the AE-35 unit.":
			ast[idx] = HalNode{instruction: shiftLeft, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "Well, he acts like he has genuine emotions.":
			ast[idx] = HalNode{instruction: shiftRight, n: 0, loopStart: 0, loopEnd: 0}

		// User and Print Output
		case instruction == "Open the pod bay doors, HAL.":
			ast[idx] = HalNode{instruction: userInput, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "Close the pod bay doors, HAL.":
			ast[idx] = HalNode{instruction: displayChar, n: 0, loopStart: 0, loopEnd: 0}

		// If nothing matches - HAL does not understand the instruction. Bail out and tell the user.
		default:
			return []HalNode{}, newCriticalHalError(fmt.Sprintf("unrecognised instruction: %q", instruction), lineNumber)
		}
	}

	// If we get to the end of a program - and we haven't closed a loop - bail out and tell the user.
	if currentLoopId != 0 {
		return []HalNode{}, newCriticalHalError("loop was not closed", loopStartIndexes[currentLoopId]+1)
	}

	return ast, nil
}