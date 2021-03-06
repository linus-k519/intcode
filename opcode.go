package main

import (
	"strconv"
)

// opcode is a two-digit operation code, like 99(END) or 01(ADD). See const
// declaration for concrete values.
type opcode uint8

// newOpcode extracts an opcode from an instruction value (Such as 01 from 12201).
func newOpcode(val int64) opcode {
	return opcode(val % 1e2)
}

type opcodeInfo struct {
	Name   string
	ArgNum int
	Fn     func(p *Program, argIndexes []int)
}

var opcodes = [...]opcodeInfo{
	1: {
		Name:   "Add",
		ArgNum: 3,
		Fn:     Add,
	},
	2: {
		Name:   "Multiply",
		ArgNum: 3,
		Fn:     Multiply,
	},
	3: {
		Name:   "Input",
		ArgNum: 1,
		Fn:     Input,
	},
	4: {
		Name:   "Output",
		ArgNum: 1,
		Fn:     Output,
	},
	5: {
		Name:   "Jump non-zero",
		ArgNum: 2,
		Fn:     JumpNonZero,
	},
	6: {
		Name:   "Jump zero",
		ArgNum: 2,
		Fn:     JumpZero,
	},
	7: {
		Name:   "Less than",
		ArgNum: 3,
		Fn:     LessThan,
	},
	8: {
		Name:   "Equal",
		ArgNum: 3,
		Fn:     Equal,
	},
	9: {
		Name:   "Add relative base",
		ArgNum: 1,
		Fn:     AddRelativeBase,
	},
	10: {
		Name:   "Bitwise And",
		ArgNum: 3,
		Fn:     BitAnd,
	},
	11: {
		Name:   "Bitwise Or",
		ArgNum: 3,
		Fn:     BitOr,
	},
	12: {
		Name:   "Bitwise Xor",
		ArgNum: 3,
		Fn:     BitXor,
	},
	13: {
		Name:   "Division",
		ArgNum: 3,
		Fn:     Division,
	},
	14: {
		Name:   "Modulo",
		ArgNum: 3,
		Fn:     Modulo,
	},
	15: {
		Name:   "Left Shift",
		ArgNum: 3,
		Fn:     LeftShift,
	},
	16: {
		Name:   "Right shift",
		ArgNum: 3,
		Fn:     RightShift,
	},
	17: {
		Name:   "Negate",
		ArgNum: 2,
		Fn:     Negate,
	},
	18: {
		Name:   "Timestamp",
		ArgNum: 1,
		Fn:     Timestamp,
	},
	19: {
		Name:   "Random",
		ArgNum: 1,
		Fn:     Random,
	},
	20: {
		Name:   "Absolute",
		ArgNum: 2,
		Fn:     Absolute,
	},
	80: {
		Name:   "Syscall",
		ArgNum: 3,
		Fn:     Syscall,
	},
	99: {
		Name:   "End",
		ArgNum: 0,
		Fn:     End,
	},
}

func (o opcode) String() string {
	name := "(" + strconv.Itoa(int(o)) + ")"
	if int(o) < len(opcodes) {
		name = opcodes[o].Name + name
	}
	return name
}
