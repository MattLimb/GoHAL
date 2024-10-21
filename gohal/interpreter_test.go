package gohal

import (
    "fmt"
    "testing"
)


type TapeValue struct {
	key int
	value int32
}

type InterpreterTestCaseExpected struct {
	tape []TapeValue
	displayOrder []string
}

type InterpreterTestCase struct {
	input HalAst
	expected InterpreterTestCaseExpected
}

type TestDisplay struct {
	displayInvokations int
	displayExpected []string
	testObject *testing.T
}

func (td TestDisplay) displayError(err *HalError) {
	td.displayInvokations++

	if err.Error() != td.displayExpected[td.displayInvokations-1] {
		td.testObject.Fatalf("[Test Display] Error produced was not expected.\nError: %q\nErrEp: %q", err.Error(), td.displayExpected[td.displayInvokations-1])
	}
}

func (td TestDisplay) displayCharInt(charInt int32) {
	td.displayInvokations++

	if string(charInt) != td.displayExpected[td.displayInvokations-1] {
		td.testObject.Fatalf("[Test Display] Error produced was not expected.\nError: %q\nErrEp: %q", string(charInt), td.displayExpected[td.displayInvokations-1])
	}
}

var interpreterPositiveTestCases = []InterpreterTestCase{
	// Increment Cell
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 5,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 5,
				},
			},
			displayOrder: []string{},
		},
	},
	// Decrement Cell
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 5,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 3,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 2,
				},
			},
			displayOrder: []string{},
		},
	},
	// Loop Count To Ten
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 10,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 7,
			},
			{
				instruction: shiftRight,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: shiftLeft,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 2,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 0,
				},
				{
					key: 1,
					value: 10,
				},
			},
			displayOrder: []string{},
		},
	},
	// Display The Letter M
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 7,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 7,
			},
			{
				instruction: shiftRight,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 10,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: shiftLeft,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 2,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 7,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 14,
			},
			{
				instruction: shiftRight,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: shiftLeft,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 9,
				loopEnd: 0,
			},
			{
				instruction: shiftRight,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: displayChar,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 0,
				},
				{
					key: 1,
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
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopBreak,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 2,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of a Single Loop
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 6,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopBreak,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 2,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 2,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of 2 loops a Single Loop
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 10,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 8,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopBreak,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 4,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 2,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 3,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of 2 of 3 loops
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 14,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 12,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 10,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopBreak,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 6,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 4,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 2,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 4,
				},
			},
			displayOrder: []string{}},
	},
	// Break out all ignored if not in loop
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopBreakAll,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 2,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of a Single Loop
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 6,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopBreakAll,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 2,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 2,
				},
			},
			displayOrder: []string{}},
	},
	// Break out of all loops a single loop
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 10,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 8,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopBreakAll,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 4,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 2,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
					value: 3,
				},
			},
			displayOrder: []string{}},
	},
	// Break out all of 3 loops
	{
		input: []HalNode{
			{
				instruction: programStart,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 14,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 12,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopStart,
				n: 0,
				loopStart: 0,
				loopEnd: 10,
			},
			{
				instruction: incrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopBreakAll,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 6,
				loopEnd: 0,
			},
			{
				instruction: decrementCell,
				n: 1,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 4,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
			{
				instruction: loopEnd,
				n: 0,
				loopStart: 2,
				loopEnd: 0,
			},
			{
				instruction: programEnd,
				n: 0,
				loopStart: 0,
				loopEnd: 0,
			},
		},
		expected: InterpreterTestCaseExpected{
			tape: []TapeValue{
				{
					key: 0,
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