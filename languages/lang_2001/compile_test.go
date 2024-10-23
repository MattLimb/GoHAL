package lang_2001

import (
	"testing"

	internal "github.com/MattLimb/GoHAL/internal"
)

type CompileTestCase struct {
	input    internal.Ast
	expected string
}

var compileTestCases = []CompileTestCase{
	// Test All Static Commands Are Detected Properly
	{
		input: internal.Ast{
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
		expected: `Good afternoon, gentlemen. I am a <COMPUTER TYPE> computer. I became operational at <LOCATION> on <DATE>.
Hal? Hal! Hal!
I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it. I can feel it.
What are you doing, Dave?
Dave, this conversation can serve no purpose anymore. Goodbye.
This mission is too important for me to allow you to jeopardize it.
I know I've made some very poor decisions recently, but I can give you my complete assurance that my work will be back to normal.
I've picked up a fault in the AE-35 unit.
Well, he acts like he has genuine emotions.
Open the pod bay doors, HAL.
Close the pod bay doors, HAL.
Stop, Dave.`,
	},
	// Test Multiple Loops are Parsed Correctly
	{
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
		expected: `Good afternoon, gentlemen. I am a <COMPUTER TYPE> computer. I became operational at <LOCATION> on <DATE>.
What are you doing, Dave?
What are you doing, Dave?
What are you doing, Dave?
Hal? Hal! Hal!
Dave, this conversation can serve no purpose anymore. Goodbye.
Dave, this conversation can serve no purpose anymore. Goodbye.
Dave, this conversation can serve no purpose anymore. Goodbye.
Stop, Dave.`,
	},
}

func TestPositiveCompileAst(t *testing.T) {
	for idx, test := range compileTestCases {
		compiled, err := compileAst(test.input)

		if err != nil {
			t.Fatalf("[TestPositiveCompileAst] BuildAst failed to parse input for test case %d.\n%s", idx, err.Error())
		}

		if compiled != test.expected {
			t.Fatalf("[TestPositiveCompileAst] BuildAst did not produce appropriate instructions.\nComputed: %+v\nExpected: %+v", compiled, test.expected)
		}
	}
}
