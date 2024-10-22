// Package internal/interpreter - Core interpreter loop.
package internal

import "fmt"

// InterpretAst is the main HAL loop. It runs every instruction.
func InterpretAst(ast Ast, tape map[int]int32, display Displayer) {
	var instruction Node

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
		case IncrementCell:
			cellValue += instruction.N
		case DecrementCell:
			cellValue -= instruction.N
		// Tape Movement Operations
		case ShiftLeft:
			tape[cellPointer] = cellValue
			cellValue = 0
			cellPointer--

			// Get the value of the current point
			// Maps return the int default value if the key doesn't exist.
			cellValue = tape[cellPointer]
		case ShiftRight:
			tape[cellPointer] = cellValue
			cellValue = 0
			cellPointer++

			// Get the value of the current point
			// Maps return the int default value if the key doesn't exist.
			cellValue = tape[cellPointer]
		// display
		case DisplayChar:
			display.DisplayCharInt(cellValue)
		// User Input
		case UserInput:
			var inputString string

			_, err := fmt.Scan(&inputString)

			if err != nil {
				display.DisplayError(NewHalError("no character inputted from user", currentIndex))
			}

			inputRune := []rune(inputString)[0]

			cellValue = inputRune

		// Looping
		case LoopStart:
			if cellValue == 0 {
				currentIndex = instruction.LoopEnd
			} else {
				loopDepth++
			}
		case LoopEnd:
			if cellValue != 0 {
				currentIndex = instruction.LoopStart
			} else {
				loopDepth--
			}
		// Loop Breaks
		case LoopBreakAll:
			for {
				if loopDepth == 0 {
					break
				}

				currentIndex++
				instruction = ast[currentIndex]

				if instruction.Instruction == LoopEnd {
					loopDepth--
				}
			}
		case LoopBreak:
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

				if instruction.Instruction == LoopEnd {
					loopDepth--
				}
			}
		case ProgramEnd:
			endProgram = true
		}

		// Write the current value to the tape
		tape[cellPointer] = cellValue

		// Step through the instructions
		currentIndex++
		tasksCompleted++
	}
}
