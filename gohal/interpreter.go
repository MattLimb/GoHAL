// Package gohal/interpreter - Core interpreter loop.
package gohal

import (
	"fmt"

	internal "github.com/MattLimb/GoHAL/gohal/internal"
)

// interpretAst is the main HAL loop. It runs every instruction.
func interpretAst(ast internal.Ast, tape map[int]int32, display internal.HalDisplayer) {
	var instruction internal.Node

	instructionLen := len(ast)

	cellPointer := 0
	var cellValue int32 = 0

	loopDepth := 0
	currentIndex := 0

	tasksCompleted := 0
	endProgram := false

	for {
		if currentIndex == instructionLen || endProgram {
			break
		}
		instruction = ast[currentIndex]

		switch instruction.Instruction {
		// Tape Value Operations
		case internal.IncrementCell:
			cellValue += instruction.N
		case internal.DecrementCell:
			cellValue -= instruction.N
		// Tape Movement Operations
		case internal.ShiftLeft:
			tape[cellPointer] = cellValue
			cellValue = 0
			cellPointer--

			// Get the value of the current point
			// Maps return the int default value if the key doesn't exist.
			cellValue = tape[cellPointer]
		case internal.ShiftRight:
			tape[cellPointer] = cellValue
			cellValue = 0
			cellPointer++

			// Get the value of the current point
			// Maps return the int default value if the key doesn't exist.
			cellValue = tape[cellPointer]
		// display
		case internal.DisplayChar:
			display.DisplayCharInt(cellValue)
		// User Input
		case internal.UserInput:
			var inputString string

			_, err := fmt.Scan(&inputString)

			if err != nil {
				display.DisplayError(internal.NewHalError("no character inputted from user", currentIndex))
			}

			inputRune := []rune(inputString)[0]

			cellValue = inputRune

		// Looping
		case internal.LoopStart:
			if cellValue == 0 {
				currentIndex = instruction.LoopEnd
			} else {
				loopDepth++
			}
		case internal.LoopEnd:
			if cellValue != 0 {
				currentIndex = instruction.LoopStart
			} else {
				loopDepth--
			}
		// Loop Breaks
		case internal.LoopBreakAll:
			for {
				if loopDepth == 0 {
					break
				}

				currentIndex++
				instruction = ast[currentIndex]

				if instruction.Instruction == internal.LoopEnd {
					loopDepth--
				}
			}
		case internal.LoopBreak:
			expectedLoopEqual := loopDepth

			if loopDepth >= 2 {
				expectedLoopEqual -= 2
			} else if loopDepth == 1 {
				expectedLoopEqual -= 1
			}

			for {
				if loopDepth == expectedLoopEqual {
					break
				}

				currentIndex++
				instruction = ast[currentIndex]

				if instruction.Instruction == internal.LoopEnd {
					loopDepth--
				}
			}
		case internal.ProgramEnd:
			endProgram = true
		}

		// Write the current value to the tape
		tape[cellPointer] = cellValue

		// Step through the instructions
		currentIndex++
		tasksCompleted++
	}
}
