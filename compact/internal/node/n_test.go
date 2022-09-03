package node_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt/compact/internal/node"
)

const testBitsLen32Error = "BitsLen[uint32] Test: result is not 32"

func TestBitsLen32(t *testing.T) {
	if node.BitsLen[uint32]() != 32 {
		t.Error(testBitsLen32Error)
	}
}

const testBitsLen64Error = "BitsLen[uint64] Test: result is not 64"

func TestBitsLen64(t *testing.T) {
	if node.BitsLen[uint64]() != 64 {
		t.Error(testBitsLen64Error)
	}
}

var head32Tests = []struct {
	n      uint32
	s      byte
	result uint
}{
	{n: 0, s: 0, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, s: 255, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, s: 32, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, s: 27, result: 0b10110},
	{
		n: 0b10101011_11001101_11101111_01110110,
		s: 16,
		result: 0b11101111_01110110,
	},
	{
		n: 0b10101011_11001101_11101111_01110110,
		s: 0,
		result: 0b10101011_11001101_11101111_01110110,
	},
}

const testHead32Error = "Head[uint32] Test %d: got %b for result, should be %b"

func TestHead32(t *testing.T) {
	for i, tt := range head32Tests {
		result := node.Head(tt.n, tt.s)
		if result != tt.result {
			t.Errorf(testHead32Error, i, result, tt.result)
		}
	}
}

var head64Tests = []struct {
	n      uint64
	s      byte
	result uint
}{
	{n: 0, s: 0, result: 0},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 255,
		result: 0,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 64,
		result: 0,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 27,
		result: 0b10110_01010100_00110010_00010000_10011000,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 16,
		result: 0b11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 0,
		result: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
	},
}

const testHead64Error = "Head[uint64] Test %d: got %b for result, should be %b"

func TestHead64(t *testing.T) {
	for i, tt := range head64Tests {
		result := node.Head(tt.n, tt.s)
		if result != tt.result {
			t.Errorf(testHead64Error, i, result, tt.result)
		}
	}
}

var tail32Tests = []struct {
	n      uint32
	s      byte
	result uint
}{
	{n: 0, s: 0, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, s: 255, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, s: 32, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, s: 27, result: 0b10101},
	{
		n: 0b10101011_11001101_11101111_01110110,
		s: 16,
		result: 0b10101011_11001101,
	},
	{
		n: 0b10101011_11001101_11101111_01110110,
		s: 0,
		result: 0b10101011_11001101_11101111_01110110,
	},
}

const testTail32Error = "Tail[uint32] Test %d: got %b for result, should be %b"

func TestTail32(t *testing.T) {
	for i, tt := range tail32Tests {
		result := node.Tail(tt.n, tt.s)
		if result != tt.result {
			t.Errorf(testTail32Error, i, result, tt.result)
		}
	}
}

var tail64Tests = []struct {
	n      uint64
	s      byte
	result uint
}{
	{n: 0, s: 0, result: 0},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 255,
		result: 0,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 64,
		result: 0,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 27,
		result: 0b10101 << 32 |
			0b01111001_10111101_11101110_11001010,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 16,
		result: 0b10101011_11001101 << 32 |
			0b11101111_01110110_01010100_00110010,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		s: 0,
		result: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
	},
}

const testTail64Error = "Tail[uint64] Test %d: got %b for result, should be %b"

func TestTail64(t *testing.T) {
	for i, tt := range tail64Tests {
		result := node.Tail(tt.n, tt.s)
		if result != tt.result {
			t.Errorf(testTail64Error, i, result, tt.result)
		}
	}
}

var body32Tests = []struct {
	n      uint32
	ls     byte
	rs     byte
	result uint
}{
	{n: 0, ls: 0, rs: 0, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, ls: 255, rs: 0, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, ls: 32, rs: 0, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, ls: 5, rs: 32, result: 0},
	{n: 0b10101011_11001101_11101111_01110110, ls: 5, rs: 255, result: 0},
	{
		n: 0b10101011_11001101_11101111_01110110,
		ls: 22,
		rs: 27,
		result: 0b11011,
	},
	{
		n: 0b10101011_11001101_11101111_01110110,
		ls: 16,
		rs: 24,
		result: 0b11101111,
	},
	{
		n: 0b10101011_11001101_11101111_01110110,
		ls: 0,
		rs: 0,
		result: 0b10101011_11001101_11101111_01110110,
	},
}

const testBody32Error = "Body[uint32] Test %d: got %b for result, should be %b"

func TestBody32(t *testing.T) {
	for i, tt := range body32Tests {
		result := node.Body(tt.n, tt.ls, tt.rs)
		if result != tt.result {
			t.Errorf(testBody32Error, i, result, tt.result)
		}
	}
}

var body64Tests = []struct {
	n      uint64
	ls     byte
	rs     byte
	result uint
}{
	{n: 0, ls: 0, rs: 0, result: 0},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		ls: 255,
		rs: 0,
		result: 0,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		ls: 64,
		rs: 0,
		result: 0,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		ls: 5,
		rs: 64,
		result: 0,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		ls: 5,
		rs: 255,
		result: 0,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		ls: 42,
		rs: 47,
		result: 0b11_0010000_10000100,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		ls: 22,
		rs: 27,
		result: 0b11011 << 32 |
			0b10110010_10100001_10010000_10000100,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		ls: 16,
		rs: 24,
		result: 0b11101111 << 32 |
			0b01110110_01010100_00110010_00010000,
	},
	{
		n: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
		ls: 0,
		rs: 0,
		result: 0b10101011_11001101_11101111_01110110 << 32 |
			0b01010100_00110010_00010000_10011000,
	},
}

const testBody64Error = "Body[uint64] Test %d: got %b for result, should be %b"

func TestBody64(t *testing.T) {
	for i, tt := range body64Tests {
		result := node.Body(tt.n, tt.ls, tt.rs)
		if result != tt.result {
			t.Errorf(testBody64Error, i, result, tt.result)
		}
	}
}
