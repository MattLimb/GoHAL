package gohal

import "fmt"

func InterpretAst(ast HalAst, display HalDisplay) {
    var instruction HalNode

    instructionLen := len(ast)
    tape := map[int]int32{}

    cellPointer := 0
    var cellValue int32 = 0

    loopDepth := 0
    currentIndex := 0

    tasksCompleted := 0

    for {
        if currentIndex == instructionLen {
            break
        }
        instruction = ast[currentIndex]

        switch instruction.Instruction {
        // Tape Value Operations
        case incrementCell:
            cellValue += instruction.n
        case decrementCell:
            cellValue -= instruction.n
        // Tape Movement Operations
        case shiftLeft:
            tape[cellPointer] = cellValue
            cellValue = 0
            cellPointer--

            // Get the value of the current point
            // Maps return the int default value if the key doesn't exist.
            cellValue = tape[cellPointer]
        case shiftRight:
            tape[cellPointer] = cellValue
            cellValue = 0
            cellPointer++

            // Get the value of the current point
            // Maps return the int default value if the key doesn't exist.
            cellValue = tape[cellPointer]
        // Display
        case displayChar:
            display.DisplayCharInt(cellValue)
        // User Input
        case userInput:
            var inputString string

            _, err := fmt.Scan(&inputString)

            if err != nil {
                display.DisplayError(NewHalError("no character inputted from user"))
            }

            inputRune := []rune(inputString)[0]

            cellValue = int32(inputRune)

        // Looping
        case loopStart:
            if cellValue == 0 {
                currentIndex = instruction.loopEnd
            }

            loopDepth++
        case loopEnd:
			if cellValue != 0 {
				currentIndex = instruction.loopStart
			}

			loopDepth--
        // Loop Breaks
        case loopBreakAll:
            for {
                if loopDepth == 0 {
                    break
                }

                currentIndex++
                instruction = ast[currentIndex]

                if instruction.Instruction == loopEnd {
                    loopDepth--
                }
            }
        case loopBreak:
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

                if instruction.Instruction == loopEnd {
                    loopDepth--
                }
            }
        }

        // Write the current value to the tape
        tape[cellPointer] = cellValue

        // Step through the instructions
        currentIndex++
        tasksCompleted++
    }
}
