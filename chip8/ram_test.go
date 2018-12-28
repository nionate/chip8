package chip8_test

import (
	"bytes"
	"testing"

	"github.com/nionate/chip8/chip8"
)

func TestShouldWriteOnRAM(t *testing.T) {
	ram := chip8.NewRAM()
	ram.Write(5, []byte{1, 2, 3, 4})
	//fmt.Println(ram)
}

func TestShouldReadFromRAM(t *testing.T) {
	ram := chip8.NewRAM()
	ram[12] = 9
	ram[13] = 8

	value, err := ram.Read(11, 4)
	if err != nil {
		t.Error(err)
	}

	expected := []byte{0, 9, 8, 0}
	if !bytes.Equal(expected, value) {
		t.Errorf("%x should be %x", value, expected)
	}
}
