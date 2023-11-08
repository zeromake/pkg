package binaryutil

type Decoder interface {
	Error() error
	Skip(size uint)
	ReadLE(r any)
	ReadBE(r any)
	ReadU32LE() uint32
	Read32LE() int32
	ReadU16LE() uint16
	Read16LE() int16
	ReadU8LE() uint8
	Read8LE() int8
	ReadU32BE() uint32
	Read32BE() int32
	ReadU16BE() uint16
	Read16BE() int16
	ReadU8BE() uint8
	Read8BE() int8
}

type Encoder interface {
	Error() error
	Skip(size uint)
	WriteLE(r any)
	WriteBE(r any)
	WriteU32LE(r uint32)
	Write32LE(r int32)
	WriteU16LE(r uint16)
	Write16LE(r int16)
	WriteU8LE(r uint8)
	Write8LE(r int8)
	WriteU32BE(r uint32)
	Write32BE(r int32)
	WriteU16BE(r uint16)
	Write16BE(r int16)
	WriteU8BE(r uint8)
	Write8BE(r int8)
}
