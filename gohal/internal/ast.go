// Package gohal_internal/ast - A Basic Abstract Syntax Tree for GoHAL.
// This file takes the user input file (as a []string) and parses out all variability into a standard format.
package gohal_internal

// Instruction is a descriptive type to tag its usage as a command type.
// They are used to tell HAL which instruction to run.
type Instruction string

const (
	IncrementCell Instruction = "increment_cell"
	DecrementCell Instruction = "decrement_cell"
	LoopStart     Instruction = "loop_start"
	LoopEnd       Instruction = "loop_end"
	LoopBreak     Instruction = "loop_break"
	LoopBreakAll  Instruction = "loop_break_all"
	ShiftLeft     Instruction = "shift_left"
	ShiftRight    Instruction = "shift_right"
	UserInput     Instruction = "user_input"
	DisplayChar   Instruction = "display_char"
	ProgramStart  Instruction = "program_start"
	ProgramEnd    Instruction = "program_end"
)

// Node is a struct to encode functionality in an easier way.
// Currently supports commands with a variable number of instructions and more efficient loop processing.
type Node struct {
	Instruction Instruction
	N           int32
	LoopStart   int
	LoopEnd     int
}

// Ast is a descriptive type to refer to a collection of HalNodes.
type Ast []Node
