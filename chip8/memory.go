package chip8

type Memory interface{
	Write(pos int16, data []byte) error
	Read(pos int16, size int16) ([]byte, error)
}

