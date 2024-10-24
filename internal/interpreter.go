// Package internal/interpreter - Core interpreter loop.
package internal

import "fmt"

const LoopMaxIteration int = 10_000

// InterpretAst is the main HAL loop. It runs every instruction.
func InterpretAst(ast Ast, tape Taper, display Displayer) *HalError {
	var ranInstructionsInLoop int
	var instruction Node

	instructionLen := len(ast)

	loopDepth := 0
	currentIndex := 0

	tasksCompleted := 0
	endProgram := false

	for {
		if currentIndex == instructionLen || endProgram {
			break
		}

		if ranInstructionsInLoop >= LoopMaxIteration {
			return NewCriticalHalError("loop max iteration exceeded", currentIndex)
		}
		instruction = ast[currentIndex]

		switch instruction.Instruction {
		// Tape Value Operations
		case IncrementCell:
			tape.IncrementCell(instruction.N)
		case DecrementCell:
			tape.DecrementCell(instruction.N)
		// Tape Movement Operations
		case ShiftLeft:
			tape.ShiftLeft()
		case ShiftRight:
			tape.ShiftRight()
		// display
		case DisplayChar:
			display.DisplayCharInt(tape.ReturnCell())
		// User Input
		case UserInput:
			var inputString string

			_, err := fmt.Scan(&inputString)

			if err != nil {
				display.DisplayError(NewHalError("no character inputted from user", currentIndex))
			}

			inputRune := []rune(inputString)[0]

			tape.SetCell(int32(inputRune))

		// Looping
		case LoopStart:
			cellValue := tape.ReturnCell()

			if cellValue == 0 {
				currentIndex = instruction.LoopEnd
			} else {
				loopDepth++
				ranInstructionsInLoop = 0
			}
		case LoopEnd:
			cellValue := tape.ReturnCell()

			if cellValue != 0 {
				currentIndex = instruction.LoopStart
			} else {
				loopDepth--
				ranInstructionsInLoop = 0
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

			ranInstructionsInLoop = 0
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
			ranInstructionsInLoop = 0
		case ProgramEnd:
			endProgram = true
		}
		// Step through the instructions
		currentIndex++
		tasksCompleted++

		// Up the number of instructions run
		if loopDepth > 0 {
			ranInstructionsInLoop++
		}
	}

	return nil
}
