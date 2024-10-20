package gohal

import (
    "regexp"
	"strings"
)

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

type HalNode struct {
	Instruction HalInstruction
	n           int32
	loopStart   int
	loopEnd		int
}

type HalAst []HalNode

func BuildAst(fileLines []string) (HalAst, error) {
	programLength := len(fileLines)

	ast := make([]HalNode, programLength)
	loopStartIndexes := map[int]int{}
	currentLoopId := 0

	startprogramRegex := regexp.MustCompile(`Good afternoon, gentlemen. I am a (.*) computer\. I became operational at (.*) on (.*).`)

	for idx, instruction := range fileLines {
		switch {
		case idx == 0:
			if startprogramRegex.MatchString(instruction) {
				ast[idx] = HalNode{Instruction: programStart, n: 0, loopStart: 0, loopEnd: 0}
			} else {
				return []HalNode{}, NewCriticalHalError("program must start with 'Good afternoon, gentlemen.' command")
			}
		case idx == (programLength-1) && instruction != "Stop, Dave.":
			return []HalNode{}, NewCriticalHalError("program must end with 'Stop, Dave.' command")
		case instruction == "Stop, Dave.":
			ast[idx] = HalNode{Instruction: programEnd, n: 0, loopStart: 0, loopEnd: 0}
		case strings.Contains(instruction, "Hal?"):
			ast[idx] =  HalNode{Instruction: incrementCell, n: int32(strings.Count(instruction, "Hal!")), loopStart: 0, loopEnd: 0}
		case strings.Contains(instruction, "I'm afraid. I'm afraid, Dave. Dave, my mind is going."):
			ast[idx] = HalNode{Instruction: decrementCell, n: int32(strings.Count(instruction, "I can feel it.")), loopStart: 0, loopEnd: 0}
		case instruction ==  "What are you doing, Dave?":
			currentLoopId++
			loopStartIndexes[currentLoopId] = idx

			ast[idx] = HalNode{Instruction: loopStart, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "Dave, this conversation can serve no purpose anymore. Goodbye.":
			if currentLoopId == 0 {
				return []HalNode{}, NewCriticalHalError("program cannot end a loop without starting one")
			}

			loopStartIdx := loopStartIndexes[currentLoopId]

			ast[idx] = HalNode{Instruction: loopEnd, n: 0, loopStart: loopStartIdx, loopEnd: 0}
			ast[loopStartIdx].loopEnd = idx

			currentLoopId--
		case instruction == "This mission is too important for me to allow you to jeopardize it.":
			ast[idx] = HalNode{Instruction: loopBreak, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "I know I've made some very poor decisions recently, but I can give you my complete assurance that my work will be back to normal.":
			ast[idx] = HalNode{Instruction: loopBreakAll, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "I've picked up a fault in the AE-35 unit.":
			ast[idx] = HalNode{Instruction: shiftLeft, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "Well, he acts like he has genuine emotions.":
			ast[idx] = HalNode{Instruction: shiftRight, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "Open the pod bay doors, HAL.":
			ast[idx] = HalNode{Instruction: userInput, n: 0, loopStart: 0, loopEnd: 0}
		case instruction == "Close the pod bay doors, HAL.":
			ast[idx] = HalNode{Instruction: displayChar, n: 0, loopStart: 0, loopEnd: 0}
		}
	}

	return ast, nil
}