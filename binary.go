// Package binutils implements easy and simple reading and writing of byte arrays.
// It writes and reads directly to and from byte arrays, rather than readers.
package binutils

import (
	"github.com/pkg/errors"
	"math"
)

const (
	BigEndian EndianType = iota
	LittleEndian
)

type EndianType byte

// Read reads from buffer at the given offset with the given length.
func Read(buffer *[]byte, offset *int, length int) ([]byte, error) {
	if len(*buffer) < *offset+length {
		return nil, errors.New("Not enough bytes to read")
	}
	var b = (*buffer)[*offset : *offset+length]
	*offset += length
	return b, nil
}

// Write writes a byte to the buffer.
func Write(buffer *[]byte, v byte) {
	*buffer = append(*buffer, v)
}

func WriteBool(buffer *[]byte, bool bool) {
	if bool {
		WriteByte(buffer, 0x01)
		return
	}
	WriteByte(buffer, 0x00)
}

func ReadBool(buffer *[]byte, offset *int) (bool, error) {
	out, err := Read(buffer, offset, 1)
	if err != nil {
		return false, errors.Wrap(err, "Error reading byte for bool")
	}
	return out[0] != 0x00, nil
}

func WriteByte(buffer *[]byte, byte byte) {
	Write(buffer, byte)
}

func ReadByte(buffer *[]byte, offset *int) (byte, error) {
	out, err := Read(buffer, offset, 1)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading byte")
	}
	return out[0], nil
}

func WriteUnsignedByte(buffer *[]byte, unsigned uint8) {
	WriteByte(buffer, byte(unsigned))
}

func ReadUnsignedByte(buffer *[]byte, offset *int) (byte, error) {
	out, err := Read(buffer, offset, 1)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading unsigned byte")
	}
	return out[0], nil
}

func WriteShort(buffer *[]byte, signed int16) {
	var b = make([]byte, 2)
	var v = uint16(signed)
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadShort(buffer *[]byte, offset *int) (int16, error) {
	b, err := Read(buffer, offset, 2)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for short")
	}
	return int16(uint16(b[1]) | uint16(b[0])<<8), nil
}

func WriteUnsignedShort(buffer *[]byte, v uint16) {
	var b = make([]byte, 2)
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadUnsignedShort(buffer *[]byte, offset *int) (uint16, error) {
	b, err := Read(buffer, offset, 2)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for unsigned short")
	}
	return uint16(b[1]) | uint16(b[0])<<8, nil
}

func WriteInt(buffer *[]byte, int int32) {
	var b = make([]byte, 4)
	var v = uint32(int)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadInt(buffer *[]byte, offset *int) (int32, error) {
	b, err := Read(buffer, offset, 4)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for int")
	}
	return int32(uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24), nil
}

func WriteUnsignedInt(buffer *[]byte, v uint32) {
	var b = make([]byte, 4)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadUnsignedInt(buffer *[]byte, offset *int) (uint32, error) {
	b, err := Read(buffer, offset, 4)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for unsigned int")
	}
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24, nil
}

func WriteLong(buffer *[]byte, long int64) {
	var b = make([]byte, 8)
	var v = uint64(long)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLong(buffer *[]byte, offset *int) (int64, error) {
	b, err := Read(buffer, offset, 8)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for long")
	}
	return int64(uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56), nil
}

func WriteUnsignedLong(buffer *[]byte, v uint64) {
	var b = make([]byte, 8)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadUnsignedLong(buffer *[]byte, offset *int) (uint64, error) {
	b, err := Read(buffer, offset, 8)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for unsigned long")
	}
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56, nil
}

func WriteFloat(buffer *[]byte, float float32) {
	var b = make([]byte, 4)
	var v = math.Float32bits(float)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadFloat(buffer *[]byte, offset *int) (float32, error) {
	b, err := Read(buffer, offset, 4)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for float")
	}

	var out = uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	return math.Float32frombits(out), nil
}

func WriteDouble(buffer *[]byte, double float64) {
	var b = make([]byte, 8)
	var v = math.Float64bits(double)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadDouble(buffer *[]byte, offset *int) (float64, error) {
	b, err := Read(buffer, offset, 8)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for double")
	}
	var out = uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	return math.Float64frombits(out), nil
}

func WriteLittleShort(buffer *[]byte, signed int16) {
	var b = make([]byte, 2)
	var v = uint16(signed)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleShort(buffer *[]byte, offset *int) (int16, error) {
	b, err := Read(buffer, offset, 2)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for little endian short")
	}
	return int16(uint16(b[0]) | uint16(b[1])<<8), nil
}

func WriteLittleUnsignedShort(buffer *[]byte, v uint16) {
	var b = make([]byte, 2)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleUnsignedShort(buffer *[]byte, offset *int) (uint16, error) {
	b, err := Read(buffer, offset, 2)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for unsigned little endian short")
	}
	return uint16(b[0]) | uint16(b[1])<<8, nil
}

func WriteLittleInt(buffer *[]byte, int int32) {
	var b = make([]byte, 4)
	var v = uint32(int)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleInt(buffer *[]byte, offset *int) (int32, error) {
	b, err := Read(buffer, offset, 4)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for little endian int")
	}
	return int32(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24), nil
}

func WriteLittleUnsignedInt(buffer *[]byte, v uint32) {
	var b = make([]byte, 4)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleUnsignedInt(buffer *[]byte, offset *int) (uint32, error) {
	b, err := Read(buffer, offset, 4)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for unsigned little endian int")
	}
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24, nil
}

func WriteLittleLong(buffer *[]byte, long int64) {
	var b = make([]byte, 8)
	var v = uint64(long)
	b[7] = byte(v >> 56)
	b[6] = byte(v >> 48)
	b[5] = byte(v >> 40)
	b[4] = byte(v >> 32)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleLong(buffer *[]byte, offset *int) (int64, error) {
	b, err := Read(buffer, offset, 8)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for little endian long")
	}
	return int64(uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56), nil
}

func WriteLittleUnsignedLong(buffer *[]byte, v uint64) {
	var b = make([]byte, 8)
	b[7] = byte(v >> 56)
	b[6] = byte(v >> 48)
	b[5] = byte(v >> 40)
	b[4] = byte(v >> 32)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleUnsignedLong(buffer *[]byte, offset *int) (uint64, error) {
	b, err := Read(buffer, offset, 8)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for unsigned little endian long")
	}
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56, nil
}

func WriteLittleFloat(buffer *[]byte, float float32) {
	var b = make([]byte, 4)
	var v = math.Float32bits(float)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleFloat(buffer *[]byte, offset *int) (float32, error) {
	b, err := Read(buffer, offset, 4)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for little endian float")
	}
	var out = uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
	return math.Float32frombits(out), nil
}

func WriteLittleDouble(buffer *[]byte, double float64) {
	var b = make([]byte, 8)
	var v = math.Float64bits(double)
	b[7] = byte(v >> 56)
	b[6] = byte(v >> 48)
	b[5] = byte(v >> 40)
	b[4] = byte(v >> 32)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleDouble(buffer *[]byte, offset *int) (float64, error) {
	b, err := Read(buffer, offset, 8)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for little endian double")
	}
	var out = uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	return math.Float64frombits(out), nil
}

func ReadBigTriad(buffer *[]byte, offset *int) (uint32, error) {
	var out uint32
	var b, err = Read(buffer, offset, 3)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for big endian triad")
	}
	out = (uint32(b[2]) & 0xFF) | ((uint32(b[1]) & 0xFF) << 8) | ((uint32(b[0]) & 0x0F) << 16)

	return out, nil
}

func WriteLittleTriad(buffer *[]byte, uint uint32) {
	Write(buffer, byte(uint&0xFF))
	Write(buffer, byte(uint>>8)&0xFF)
	Write(buffer, byte(uint>>16)&0xFF)
}

func ReadLittleTriad(buffer *[]byte, offset *int) (uint32, error) {
	var b, err = Read(buffer, offset, 3)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading bytes for little endian triad")
	}
	return (uint32(b[0]) & 0xFF) | ((uint32(b[1]) & 0xFF) << 8) | ((uint32(b[2]) & 0x0F) << 16), nil
}

func WriteBigTriad(buffer *[]byte, uint uint32) {
	Write(buffer, byte(uint>>16)&0xFF)
	Write(buffer, byte(uint>>8)&0xFF)
	Write(buffer, byte(uint&0xFF))
}

func toZigZag32(n int32) uint32 {
	return (uint32)((n << 1) ^ (n >> 31))
}

func fromZigZag32(n uint32) int32 {
	return (int32)(n>>1) ^ -(int32)(n&1)
}

func toZigZag64(n int64) uint64 {
	return (uint64)((n << 1) ^ (n >> 63))
}

func fromZigZag64(n uint64) int64 {
	return (int64)(n>>1) ^ -(int64)(n&1)
}

func WriteVarInt(buffer *[]byte, value int32) {
	WriteUnsignedVarInt(buffer, toZigZag32(value))
}

func ReadVarInt(buffer *[]byte, offset *int) (int32, error) {
	u, err := ReadUnsignedVarInt(buffer, offset)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading varint")
	}
	return fromZigZag32(u), nil
}

func WriteVarLong(buffer *[]byte, value int64) {
	WriteUnsignedVarLong(buffer, toZigZag64(value))
}

func ReadVarLong(buffer *[]byte, offset *int) (int64, error) {
	u, err := ReadUnsignedVarLong(buffer, offset)
	if err != nil {
		return 0, errors.Wrap(err, "Error reading var long")
	}
	return fromZigZag64(u), nil
}

func WriteUnsignedVarInt(buffer *[]byte, value uint32) {
	var x int32 = -128
	for (value & uint32(x)) != 0 {
		Write(buffer, byte((value&0x7F)|0x80))
		value >>= 7
	}

	Write(buffer, byte(value))
}

func ReadUnsignedVarInt(buffer *[]byte, offset *int) (uint32, error) {
	result := uint32(0)
	j := uint32(0)
	var b0 byte

	// do-while https://stackoverflow.com/a/32844744
	for ok := true; ok; ok = (b0 & 0x80) != 0 {
		b, err := Read(buffer, offset, 1)
		if err != nil {
			return 0, errors.Wrap(err, "Couldn't get next byte of unsigned varint")
		}
		b0 = b[0]
		if b0 < 0 {
			return 0, errors.New("not enough bytes for unsigned varint")
		}
		result |= uint32(b0&0x7f) << (j * 7)
		j++
		if j > 5 { // up to 5 bytes in varint
			return 0, errors.New("Unsigned varint too big")
		}
	}

	return result, nil
}

func WriteUnsignedVarLong(buffer *[]byte, value uint64) {
	var x int64 = -128
	for (value & uint64(x)) != 0 {
		Write(buffer, byte((value&0x7F)|0x80))
		value >>= 7
	}

	Write(buffer, byte(value))
}

func ReadUnsignedVarLong(buffer *[]byte, offset *int) (uint64, error) {
	result := uint64(0)
	j := uint64(0)
	var b0 byte

	// do-while https://stackoverflow.com/a/32844744
	for ok := true; ok; ok = (b0 & 0x80) != 0 {
		b, err := Read(buffer, offset, 1)
		if err != nil {
			return 0, errors.Wrap(err, "Error reading unsigned var long")
		}
		b0 = b[0]
		if b0 < 0 {
			return 0, errors.New("Not enough bytes for unsigned var long")
		}
		result |= uint64(b0&0x7f) << (j * 7)
		j++
		if j > 10 { // up to 10 bytes in varlong
			return 0, errors.New("Unsigned var long too big")
		}
	}
	return result, nil
}

func WriteString(buffer *[]byte, str string) {
	WriteUnsignedVarInt(buffer, uint32(len(str)))
	*buffer = append(*buffer, []byte(str)...)
}

func ReadString(buffer *[]byte, offset *int) (string, error) {
	l, err := ReadUnsignedVarInt(buffer, offset)
	if err != nil {
		return "", errors.Wrap(err, "Error reading length of string")
	}
	strbytes, err := Read(buffer, offset, int(l))
	if err != nil {
		return "", errors.Wrap(err, "Error reading the bytes of the string")
	}
	return string(strbytes), nil
}
