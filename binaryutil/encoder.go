package binaryutil

import (
	"encoding/binary"
	"io"
)

type encoder struct {
	w   io.Writer
	err error
}

func (e *encoder) Error() error {
	return e.err
}

func (e *encoder) Skip(size uint) {
	var buf = make([]byte, size)
	_, err := e.w.Write(buf)
	e.err = err
}

func (e *encoder) WriteLE(r any) {
	err := binary.Write(e.w, binary.LittleEndian, r)
	e.err = err
}

func (e *encoder) WriteBE(r any) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) WriteU32LE(r uint32) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) Write32LE(r int32) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) WriteU16LE(r uint16) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) Write16LE(r int16) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) WriteU8LE(r uint8) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) Write8LE(r int8) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) WriteU32BE(r uint32) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) Write32BE(r int32) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) WriteU16BE(r uint16) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) Write16BE(r int16) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) WriteU8BE(r uint8) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func (e *encoder) Write8BE(r int8) {
	err := binary.Write(e.w, binary.BigEndian, r)
	e.err = err
}

func NewEncoder(w io.Writer) Encoder {
	return &encoder{w, nil}
}
