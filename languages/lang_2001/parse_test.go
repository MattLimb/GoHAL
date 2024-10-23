package lang_2001

import (
	"testing"

	internal "github.com/MattLimb/GoHAL/internal"
)

type AstTestCase struct {
	input         []string
	expected      internal.Ast
	expectedError string
}

var astPositiveTestCases = []AstTestCase{
	// Test All Static Commands Are Detected Properly
	{
		input: []string{
			"Good afternoon, gentlemen. I am a Hello World computer. I became operational at Testing Lane on October 21st, 2024.",
			"Hal? Hal! Hal!",
			"I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it. I can feel it.",
			"What are you doing, Dave?",
			"Dave, this conversation can serve no purpose anymore. Goodbye.",
			"This mission is too important for me to allow you to jeopardize it.",
			"I know I've made some very poor decisions recently, but I can give you my complete assurance that my work will be back to normal.",
			"I've picked up a fault in the AE-35 unit.",
			"Well, he acts like he has genuine emotions.",
			"Open the pod bay doors, HAL.",
			"Close the pod bay doors, HAL.",
			"Stop, Dave.",
		},
		expected: internal.Ast{
			internal.Node{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.IncrementCell,
				N:           2,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.DecrementCell,
				N:           2,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     4,
			},
			internal.Node{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   3,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.LoopBreak,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.LoopBreakAll,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.ShiftLeft,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.ShiftRight,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.UserInput,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			internal.Node{
				Instruction: internal.DisplayChar,
				N:           0,
				LoopStart:   0,
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
	// Test Multiple Loops are Parsed Correctly
	{
		input: []string{
			"Good afternoon, gentlemen. I am a Hello World computer. I became operational at Testing Lane on October 21st, 2024.",
			"What are you doing, Dave?", // [1] Loop 1 Start [End Loop: 7]
			"What are you doing, Dave?", // [2] Loop 2 Start [End Loop: 6]
			"What are you doing, Dave?", // [3] Loop 1 Start [End Loop: 5]
			"Hal? Hal! Hal!",
			"Dave, this conversation can serve no purpose anymore. Goodbye.", // [5] Loop 3 End [Start Loop: 3]
			"Dave, this conversation can serve no purpose anymore. Goodbye.", // [6] Loop 3 End [Start Loop: 2]
			"Dave, this conversation can serve no purpose anymore. Goodbye.", // [7] Loop 3 End [Start Loop: 1]
			"Stop, Dave.",
		},
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
				LoopEnd:     7,
			},
			internal.Node{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     6,
			},
			internal.Node{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     5,
			},
			internal.Node{
				Instruction: internal.IncrementCell,
				N:           2,
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
		computedAST, err := parse2001Code(test.input)

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
	// Missing Ending
	{
		input: []string{
			"Good afternoon, gentlemen. I am a Hello World computer. I became operational at Testing Lane on October 21st, 2024.",
		},
		expected:      internal.Ast{},
		expectedError: "program is too short. It must be at least 2 lines long.",
	},
	{
		input: []string{
			"Good afternoon, gentlemen. I am a Hello World computer. I became operational at Testing Lane on October 21st, 2024.",
			"Hal? Hal!",
		},
		expected:      internal.Ast{},
		expectedError: "program must end with 'Stop, Dave.' command",
	},
	// No Start String
	{
		input: []string{
			"Hal? Hal! Hal!",
			"Stop, Dave.",
		},
		expected:      internal.Ast{},
		expectedError: "program must start with 'Good afternoon, gentlemen.' command",
	},
	// No Start Loop
	{
		input: []string{
			"Good afternoon, gentlemen. I am a Hello World computer. I became operational at Testing Lane on October 21st, 2024.",
			"Dave, this conversation can serve no purpose anymore. Goodbye.",
			"Stop, Dave.",
		},
		expected:      internal.Ast{},
		expectedError: "program cannot end a loop without starting one",
	},
	// No End Loop
	{
		input: []string{
			"Good afternoon, gentlemen. I am a Hello World computer. I became operational at Testing Lane on October 21st, 2024.",
			"What are you doing, Dave?",
			"Stop, Dave.",
		},
		expected:      internal.Ast{},
		expectedError: "loop was not closed",
	},
	// Unknown Instruction
	{
		input: []string{
			"Good afternoon, gentlemen. I am a Hello World computer. I became operational at Testing Lane on October 21st, 2024.",
			"I'm sorry, Dave. I'm afraid I can't do that.",
			"Stop, Dave.",
		},
		expected:      internal.Ast{},
		expectedError: "unrecognised Instruction: \"I'm sorry, Dave. I'm afraid I can't do that.\"",
	},
}

func TestNegativeBuildAst(t *testing.T) {
	for idx, test := range astNegativeTestCases {
		_, err := parse2001Code(test.input)
		if err == nil {
			t.Fatalf("[TestNegBuildAst] gohal_internal.Ast successfully parsed on Case %d", idx)
		}

		if err.Error() != test.expectedError {
			t.Fatalf("[TestNegBuildAst] gohal_internal.Ast produced the wrong error.\nError: %q\nErrEx: %q", err.Error(), test.expectedError)
		}
	}
}
