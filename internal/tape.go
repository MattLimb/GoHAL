package internal

type Taper interface {
	ShiftLeft()
	ShiftRight()
	IncrementCell(n int32)
	DecrementCell(n int32)
	SetCell(n int32)
	ReturnCell() int32
}

type DefaultTape struct {
	pointer int
	tape    map[int]int32
}

func NewDefaultTape() *DefaultTape {
	return &DefaultTape{pointer: 0, tape: map[int]int32{0: 0}}
}

func (t *DefaultTape) ShiftRight() {
	t.pointer += 1
}

func (t *DefaultTape) ShiftLeft() {
	t.pointer -= 1
}

func (t DefaultTape) IncrementCell(n int32) {
	t.tape[t.pointer] += n
}

func (t DefaultTape) DecrementCell(n int32) {
	t.tape[t.pointer] -= n
}

func (t DefaultTape) SetCell(n int32) {
	t.tape[t.pointer] = n
}

func (t DefaultTape) ReturnCell() int32 {
	return t.tape[t.pointer]
}
