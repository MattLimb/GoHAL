package lang_brainfuck

import (
	"testing"

	internal "github.com/MattLimb/GoHAL/internal"
)

type CompileTestCase struct {
	input         internal.Ast
	expected      string
	expectedError string
}

var compilePositiveTestCases = []CompileTestCase{
	// Test All Static Commands Are Detected Properly
	{
		expected: "++--[]<>,.",
		input: internal.Ast{
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
		expected: "[[[+-]]]",
		input: internal.Ast{
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

func TestPositiveCompileAst(t *testing.T) {
	for idx, test := range compilePositiveTestCases {
		compile, err := compileAst(test.input)
		if err != nil {
			t.Fatalf("[TestPositiveCompileAst] internal.Ast failed to ccompile input for test case %d.\n%s", idx, err.Error())
		}

		if compile != test.expected {
			t.Fatalf("[TestPositiveCompileAst] internal.Ast failed to compile input for test case %d.\nCompiled: %s\nExpected: %s", idx, compile, test.expected)
		}
	}
}

var compileNegativeTestCases = []CompileTestCase{
	{
		expected: "",
		input: internal.Ast{
			{
				Instruction: internal.LoopBreak,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expectedError: "brainfuck does not support Breaking Loops",
	},
	{
		expected: "",
		input: internal.Ast{
			{
				Instruction: internal.LoopBreakAll,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expectedError: "brainfuck does not support Breaking Loops",
	},
}

func TestNegativeCompileAst(t *testing.T) {
	for idx, test := range compileNegativeTestCases {
		_, err := compileAst(test.input)

		if err == nil {
			t.Fatalf("[TestNegativeCompileAst] gohal_internal.Ast successfully parsed on Case %d", idx)
		}

		if err.Error() != test.expectedError {
			t.Fatalf("[TestNegativeCompileAst] gohal_internal.Ast produced the wrong error.\nError: %q\nErrEx: %q", err.Error(), test.expectedError)
		}
	}
}
