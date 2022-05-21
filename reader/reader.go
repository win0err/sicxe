package reader

import (
	"bufio"
	"encoding/hex"
	"io"

	"github.com/win0err/sicxe/vm"
)

type ObjReader struct {
	rd bufio.Reader
}

func NewReader(rd io.Reader) *ObjReader {
	return &ObjReader{rd: *bufio.NewReader(rd)}
}

func (or *ObjReader) Read(p []byte) (n int, err error) {
	return or.rd.Read(p)
}

func (or *ObjReader) readBytes(symbols int) ([]byte, error) {
	buf := make([]byte, symbols)
	if _, err := or.rd.Read(buf); err != nil {
		return nil, err
	}

	k, err := hex.Decode(buf, buf)
	if err != nil {
		return nil, err
	}

	return buf[:k], nil
}

func (or *ObjReader) ReadBytes(count int) ([]vm.Byte, error) {
	data := make([]vm.Byte, count)

	for i := 0; i < count; i++ {
		value, err := or.ReadByte()
		if err != nil {
			return nil, err
		}

		data[i] = value
	}

	return data, nil
}

func (or *ObjReader) ReadByte() (vm.Byte, error) {
	buf, err := or.readBytes(2)
	if err != nil {
		return 0, err
	}

	return vm.Byte(buf[0]), nil
}

func (or *ObjReader) ReadWord() (vm.Word, error) {
	buf, err := or.readBytes(6)
	if err != nil {
		return 0, err
	}

	word := int(buf[2]) | int(buf[1])<<8 | int(buf[0])<<16

	return vm.Word(word), nil
}

func (or *ObjReader) SkipLine() error {
	_, err := or.rd.ReadBytes('\n')

	return err
}
