package gopwn

import (
	"encoding/binary"
)

// P16L packs a uint16 into a byte slice in little endian format
func P16L(i uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return b
}

// P16L packs a uint16 into a byte slice in big endian format
func P16B(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

// P32L packs a uint32 into a byte slice in little endian format
func P32L(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

// P32L packs a uint32 into a byte slice in big endian format
func P32B(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

// P64L packs a uint64 into a byte slice in little endian format
func P64L(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

// P32L packs a uint32 into a byte slice in big endian format
func P64B(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

// U16L unpacks a byte slice in little endian format into a uint16
func U16L(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

// U16L unpacks a byte slice in big endian format into a uint16
func U16B(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

// U32L unpacks a byte slice in little endian format into a uint32
func U32L(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

// U32L unpacks a byte slice in big endian format into a uint32
func U32B(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

// U64L unpacks a byte slice in little endian format into a uint64
func U64L(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

// U64L unpacks a byte slice in big endian format into a uint64
func U64B(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
