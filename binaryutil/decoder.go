package binaryutil

import (
	"encoding/binary"
	"io"
)

type decoder struct {
	r   io.Reader
	err error
}

func (e *decoder) Error() error {
	return e.err
}

func (e *decoder) Skip(size uint) {
	var buf = make([]byte, size)
	_, err := io.ReadFull(e.r, buf)
	e.err = err
}

func (e *decoder) ReadLE(r any) {
	err := binary.Read(e.r, binary.LittleEndian, r)
	e.err = err
}

func (e *decoder) ReadBE(r any) {
	err := binary.Read(e.r, binary.BigEndian, r)
	e.err = err
}

func (e *decoder) ReadU32LE() uint32 {
	var r uint32
	err := binary.Read(e.r, binary.LittleEndian, &r)
	e.err = err
	return r
}

func (e *decoder) Read32LE() int32 {
	var r int32
	err := binary.Read(e.r, binary.LittleEndian, &r)
	e.err = err
	return r
}

func (e *decoder) ReadU16LE() uint16 {
	var r uint16
	err := binary.Read(e.r, binary.LittleEndian, &r)
	e.err = err
	return r
}

func (e *decoder) Read16LE() int16 {
	var r int16
	err := binary.Read(e.r, binary.LittleEndian, &r)
	e.err = err
	return r
}

func (e *decoder) ReadU8LE() uint8 {
	var r uint8
	err := binary.Read(e.r, binary.LittleEndian, &r)
	e.err = err
	return r
}

func (e *decoder) Read8LE() int8 {
	var r int8
	err := binary.Read(e.r, binary.LittleEndian, &r)
	e.err = err
	return r
}

func (e *decoder) ReadU32BE() uint32 {
	var r uint32
	err := binary.Read(e.r, binary.BigEndian, &r)
	e.err = err
	return r
}

func (e *decoder) Read32BE() int32 {
	var r int32
	err := binary.Read(e.r, binary.BigEndian, &r)
	e.err = err
	return r
}

func (e *decoder) ReadU16BE() uint16 {
	var r uint16
	err := binary.Read(e.r, binary.BigEndian, &r)
	e.err = err
	return r
}

func (e *decoder) Read16BE() int16 {
	var r int16
	err := binary.Read(e.r, binary.BigEndian, &r)
	e.err = err
	return r
}

func (e *decoder) ReadU8BE() uint8 {
	var r uint8
	err := binary.Read(e.r, binary.BigEndian, &r)
	e.err = err
	return r
}

func (e *decoder) Read8BE() int8 {
	var r int8
	err := binary.Read(e.r, binary.BigEndian, &r)
	e.err = err
	return r
}

func NewDecoder(r io.Reader) Decoder {
	return &decoder{r, nil}
}
