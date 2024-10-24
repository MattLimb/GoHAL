package lang_brainalpha

const CharOffset int32 = 65

type BrainalphaTape struct {
	pointer int
	tape    map[int]int32
}

func NewBrainalphaTape() *BrainalphaTape {
	return &BrainalphaTape{pointer: 0, tape: map[int]int32{0: 0}}
}

func (t *BrainalphaTape) ShiftRight() {
	t.pointer += 1
}

func (t *BrainalphaTape) ShiftLeft() {
	t.pointer -= 1
}

func (t BrainalphaTape) IncrementCell(n int32) {
	val := t.tape[t.pointer]

	t.tape[t.pointer] = (val + n) % CharOffset
}

func (t BrainalphaTape) DecrementCell(n int32) {
	val := t.tape[t.pointer]

	t.tape[t.pointer] = (val - n) % CharOffset
}

func (t BrainalphaTape) SetCell(n int32) {
	if n >= 65 && n <= 90 {
		t.tape[t.pointer] = n % CharOffset
	}
}

func (t BrainalphaTape) ReturnCell() int32 {
	return t.tape[t.pointer] + CharOffset
}
