package lang_brainalpha

import (
	"testing"

	internal "github.com/MattLimb/GoHAL/internal"
)

type AstTestCase struct {
	input         string
	expected      internal.Ast
	expectedError string
}

var astPositiveTestCases = []AstTestCase{
	// Test All Static Commands Are Detected Properly
	{
		input: "++--[]<>,.",
		expected: internal.Ast{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           2,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           2,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     4,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   3,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ShiftLeft,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ShiftRight,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.UserInput,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DisplayChar,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expectedError: "",
	},
	// Test Multiple Loops are Parsed Correctly
	{
		input: "[[[+-]]]",
		expected: internal.Ast{
			internal.Node{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     8,
			},
			internal.Node{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     7,
			},
			internal.Node{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     6,
			},
			internal.Node{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   3,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   1,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expectedError: "",
	},
}

func TestPositiveBuildAst(t *testing.T) {
	for idx, test := range astPositiveTestCases {
		computedAST, err := parseBrainfuckCode(test.input)

		if err != nil {
			t.Fatalf("[TestBuildAst] gohal_internal.Ast failed to parse input for test case %d.\n%s", idx, err.Error())
		}

		if len(computedAST) < len(test.expected) {
			t.Fatalf("[TestBuildAst] BuildAst missed instructions.\nComputed: %+v\nExpected: %+v", computedAST, test.expected)
		} else if len(computedAST) > len(test.expected) {
			t.Fatalf("[TestBuildAst] BuildAst added instructions.\nComputed: %+v\nExpected: %+v", computedAST, test.expected)
		}

		for nodeIdx, node := range computedAST {
			expectedNode := test.expected[nodeIdx]

			if node.Instruction != expectedNode.Instruction || node.N != expectedNode.N || node.LoopStart != expectedNode.LoopStart || node.LoopEnd != expectedNode.LoopEnd {
				t.Fatalf("[TestBuildAst] BuildAST generated inconsistent command for program line %d.\nComputed: %+v\nExpected: %+v", nodeIdx+1, node, expectedNode)
			}
		}
	}
}

var astNegativeTestCases = []AstTestCase{
	// No Start Loop
	{
		input:         "]",
		expected:      internal.Ast{},
		expectedError: "program cannot end a loop without starting one",
	},
	// No End Loop
	{
		input:         "[",
		expected:      internal.Ast{},
		expectedError: "loop was not closed",
	},
	// Unknown Instruction
	{
		input:         "?",
		expected:      internal.Ast{},
		expectedError: "unrecognised Instruction: '?'",
	},
}

func TestNegativeBuildAst(t *testing.T) {
	for idx, test := range astNegativeTestCases {
		_, err := parseBrainfuckCode(test.input)

		if err == nil {
			t.Fatalf("[TestNegBuildAst] gohal_internal.Ast successfully parsed on Case %d", idx)
		}

		if err.Error() != test.expectedError {
			t.Fatalf("[TestNegBuildAst] gohal_internal.Ast produced the wrong error.\nError: %q\nErrEx: %q", err.Error(), test.expectedError)
		}
	}
}
