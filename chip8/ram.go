package chip8

import "errors"

type RAM []byte

func (r RAM) Write(pos int16, data []byte) error {
	for i, d := range data {
		r[pos+int16(i)] = d
	}
	return nil
}

func (r RAM) Read(pos uint16, size uint16) ([]byte, error) {
	if int(pos) >= len(r) || int(pos+size) >= len(r) {
		return nil, errors.New("Memory position not found")
	}
	return r[pos : pos+size], nil
}

func NewRAM() RAM {
	return make(RAM, 64)
}
