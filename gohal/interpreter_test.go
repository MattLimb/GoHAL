package gohal

import (
	"fmt"
	"testing"

	internal "github.com/MattLimb/GoHAL/gohal/internal"
)

type TapeValue struct {
	key   int
	value int32
}

type InterpreterTestCaseExpected struct {
	tape         []TapeValue
	displayOrder []string
}

type InterpreterTestCase struct {
	input    internal.Ast
	expected InterpreterTestCaseExpected
}

type TestDisplay struct {
	displayInvokations int
	displayExpected    []string
	testObject         *testing.T
}

func (td TestDisplay) DisplayError(err *internal.HalError) {
	td.displayInvokations++

	if err.Error() != td.displayExpected[td.displayInvokations-1] {
		td.testObject.Fatalf("[Test Display] Error produced was not expected.\nError: %q\nErrEp: %q", err.Error(), td.displayExpected[td.displayInvokations-1])
	}
}

func (td TestDisplay) DisplayCharInt(charInt int32) {
	td.displayInvokations++

	if string(charInt) != td.displayExpected[td.displayInvokations-1] {
		td.testObject.Fatalf("[Test Display] Error produced was not expected.\nError: %q\nErrEp: %q", string(charInt), td.displayExpected[td.displayInvokations-1])
	}
}

var interpreterPositiveTestCases = []InterpreterTestCase{
	// Increment Cell
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           5,
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
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 5,
				},
			},
			displayOrder: []string{},
		},
	},
	// Decrement Cell
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           5,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           3,
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
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 2,
				},
			},
			displayOrder: []string{},
		},
	},
	// Loop Count To Ten
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           10,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     7,
			},
			{
				Instruction: internal.ShiftRight,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ShiftLeft,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 0,
				},
				{
					key:   1,
					value: 10,
				},
			},
			displayOrder: []string{},
		},
	},
	// Display The Letter M
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           7,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     7,
			},
			{
				Instruction: internal.ShiftRight,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           10,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ShiftLeft,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           7,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     14,
			},
			{
				Instruction: internal.ShiftRight,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ShiftLeft,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   9,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ShiftRight,
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
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 0,
				},
				{
					key:   1,
					value: 77,
				},
			},
			displayOrder: []string{
				"M",
			},
		},
	},
	// Break out ignored If not In Loop
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopBreak,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
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
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 2,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of a Single Loop
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     6,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopBreak,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 2,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of 2 loops a Single Loop
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     10,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     8,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopBreak,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   4,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 3,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of 2 of 3 loops
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     14,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     12,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     10,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopBreak,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   6,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   4,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 4,
				},
			},
			displayOrder: []string{}},
	},
	// Break out all ignored if not in loop
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopBreakAll,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
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
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 2,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of a Single Loop
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     6,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopBreakAll,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 2,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of all loops a single loop
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     10,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     8,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopBreakAll,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   4,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 3,
				},
			},
			displayOrder: []string{}},
	},
	// Break out all of 3 loops
	{
		input: []internal.Node{
			{
				Instruction: internal.ProgramStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     14,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     12,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopStart,
				N:           0,
				LoopStart:   0,
				LoopEnd:     10,
			},
			{
				Instruction: internal.IncrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopBreakAll,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   6,
				LoopEnd:     0,
			},
			{
				Instruction: internal.DecrementCell,
				N:           1,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   4,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
			{
				Instruction: internal.LoopEnd,
				N:           0,
				LoopStart:   2,
				LoopEnd:     0,
			},
			{
				Instruction: internal.ProgramEnd,
				N:           0,
				LoopStart:   0,
				LoopEnd:     0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key:   0,
					value: 4,
				},
			},
			displayOrder: []string{}},
	},
}

func TestPositiveInterpretAst(t *testing.T) {
	for idx, test := range interpreterPositiveTestCases {
		computedTape := map[int]int32{}
		testDisplay := TestDisplay{displayInvokations: 0, displayExpected: test.expected.displayOrder, testObject: t}

		interpretAst(test.input, computedTape, testDisplay)

		fmt.Printf("[%d] %+v\n", idx, computedTape)

		for _, kv := range test.expected.tape {
			value, ok := computedTape[kv.key]

			if !ok {
				t.Fatalf("[TestPositiveInterpretAst] Tape Value %d not found within Tape.", kv.key)
			}

			if value != kv.value {
				t.Fatalf("[TestPositiveInterpretAst] Tape Key %d is the wrong value.\nVal: %d\nExp: %d", kv.key, value, kv.value)
			}
		}
	}
}
