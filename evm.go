package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

// Stack tượng trưng để lưu trữ biểu thức (dạng chuỗi) thay vì giá trị thật
type SymbolicStack struct {
	items []string
}

func (s *SymbolicStack) Push(val string) {
	s.items = append(s.items, val)
}

func (s *SymbolicStack) Pop() string {
	if len(s.items) == 0 {
		return "???" // Xử lý lỗi underflow nếu bytecode bị lỗi
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func main4() {
	// Bytecode mẫu: If (storage[0] > 1) { storage[1] = 20 } else { storage[1] = 10 }
	bytecodeHex := "600160005411600f57600a600155005b601460015500"
	bytecode, err := hex.DecodeString(bytecodeHex)
	if err != nil {
		log.Fatalf("Lỗi decode hex: %v", err)
	}

	stack := &SymbolicStack{}

	fmt.Println("=== 1. DISASSEMBLY (Dịch mã máy sang Opcode) ===")
	// Cấu trúc lưu lại để in mã giả
	type PseudoCode struct {
		Offset int
		Code   string
	}
	var pseudoCodes []PseudoCode

	for pc := 0; pc < len(bytecode); {
		op := bytecode[pc]
		offset := pc

		var opName, pseudo string

		switch {
		case op == 0x00: // STOP
			opName = "STOP"
			pseudo = "return;"
			pc++

		case op >= 0x60 && op <= 0x7f: // PUSH1 đến PUSH32
			n := int(op - 0x60 + 1)
			opName = fmt.Sprintf("PUSH%d", n)
			if pc+n < len(bytecode) {
				argHex := hex.EncodeToString(bytecode[pc+1 : pc+n+1])
				// Chuyển hex sang số nguyên dạng chuỗi cho dễ đọc
				argInt, _ := strconv.ParseInt(argHex, 16, 64)
				argStr := fmt.Sprintf("%d", argInt)

				stack.Push(argStr) // Đẩy giá trị vào stack
				opName = fmt.Sprintf("%s 0x%s", opName, argHex)
			}
			pc += n + 1

		case op == 0x54: // SLOAD
			opName = "SLOAD"
			slot := stack.Pop()
			expr := fmt.Sprintf("storage[%s]", slot)
			stack.Push(expr) // Đẩy biểu thức đọc storage vào stack
			pc++

		case op == 0x55: // SSTORE
			opName = "SSTORE"
			val := stack.Pop()
			slot := stack.Pop()
			pseudo = fmt.Sprintf("storage[%s] = %s;", slot, val)
			pc++

		case op == 0x11: // GT (Greater Than)
			opName = "GT"
			a := stack.Pop() // EVM stack pop: phần tử đầu là a
			b := stack.Pop() // phần tử tiếp là b
			expr := fmt.Sprintf("(%s > %s)", a, b)
			stack.Push(expr)
			pc++

		case op == 0x57: // JUMPI (Nhảy có điều kiện)
			opName = "JUMPI"
			dest := stack.Pop()
			cond := stack.Pop()
			pseudo = fmt.Sprintf("if %s goto label_%s;", cond, dest)
			pc++

		case op == 0x5b: // JUMPDEST (Điểm đến của lệnh nhảy)
			opName = "JUMPDEST"
			pseudo = fmt.Sprintf("label_%d:", offset)
			pc++

		default:
			opName = fmt.Sprintf("UNKNOWN(0x%02x)", op)
			pc++
		}

		fmt.Printf("[%02d] %s\n", offset, opName)

		if pseudo != "" {
			pseudoCodes = append(pseudoCodes, PseudoCode{Offset: offset, Code: pseudo})
		}
	}

	fmt.Println("\n=== 2. DECOMPILED PSEUDO-CODE (Mã giả Solidity) ===")
	for _, p := range pseudoCodes {
		// Thêm thụt lề cho đẹp nếu không phải label
		indent := "    "
		if p.Code[len(p.Code)-1] == ':' {
			indent = ""
		}
		fmt.Printf("[%02d] %s%s\n", p.Offset, indent, p.Code)
	}
}
