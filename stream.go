package binutils

// Stream is a container of a byte array and an offset.
// Reading from the stream increments the offset.
type Stream struct {
	Offset int
	Buffer []byte
	err    *streamErr
}

type StreamErr interface {
	// Offset gives the offset at which the error happened
	Offset() int
	// Buffer gives the buffer which caused the error
	Buffer() []byte
}

type streamErr struct {
	err    error
	offset int
	buf    *[]byte
}

func (se streamErr) Error() string {
	return se.err.Error()
}

func (se streamErr) Offset() int {
	return se.offset
}

func (se streamErr) Buffer() []byte {
	return *se.buf
}

// NewStream returns a new stream.
func NewStream() *Stream {
	return &Stream{0, []byte{}, &streamErr{}}
}

// NewGetStream gets a stream for reading
func NewGetStream(buf []byte, offset int) *Stream {
	return &Stream{offset, buf, &streamErr{buf: &buf}}
}

// Error returns any error that has been encountered on this stream
func (stream *Stream) Error() error {
	return stream.err
}

// SetError allows to set the error message on the stream
func (stream *Stream) SetError(err error) {
	stream.err.err = err
	stream.err.offset = stream.Offset
}

// GetOffset returns the current stream offset.
func (stream *Stream) GetOffset() int {
	return stream.Offset
}

// SetOffset sets the offset of the stream.
func (stream *Stream) SetOffset(offset int) {
	stream.Offset = offset
}

// SetBuffer sets the buffer of the stream.
func (stream *Stream) SetBuffer(buffer []byte) {
	stream.Buffer = buffer
}

// GetBuffer returns the buffer of the stream.
func (stream *Stream) GetBuffer() []byte {
	return stream.Buffer
}

// Feof checks if the stream offset reached the end of its buffer.
func (stream *Stream) Feof() bool {
	return stream.Offset >= len(stream.Buffer)-1
}

// Get reads the given amount of bytes from the buffer.
// If length is negative, reads the leftover bytes.
func (stream *Stream) Get(length int) []byte {
	if stream.err != nil {
		return nil
	}
	if length < 0 {
		length = len(stream.Buffer) - stream.Offset - 1
	}
	b, err := Read(&stream.Buffer, &stream.Offset, length)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutBool(v bool) {
	WriteBool(&stream.Buffer, v)
}

func (stream *Stream) GetBool() bool {
	b, err := ReadBool(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutByte(v byte) {
	WriteByte(&stream.Buffer, v)
}

func (stream *Stream) GetByte() byte {
	b, err := ReadByte(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutUnsignedByte(v byte) {
	WriteUnsignedByte(&stream.Buffer, v)
}

func (stream *Stream) GetUnsignedByte() byte {
	b, err := ReadUnsignedByte(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutShort(v int16) {
	WriteShort(&stream.Buffer, v)
}

func (stream *Stream) GetShort() int16 {
	b, err := ReadShort(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutUnsignedShort(v uint16) {
	WriteUnsignedShort(&stream.Buffer, v)
}

func (stream *Stream) GetUnsignedShort() uint16 {
	b, err := ReadUnsignedShort(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutInt(v int32) {
	WriteInt(&stream.Buffer, v)
}

func (stream *Stream) GetInt() int32 {
	b, err := ReadInt(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutUnsignedInt(v uint32) {
	WriteUnsignedInt(&stream.Buffer, v)
}

func (stream *Stream) GetUnsignedInt() uint32 {
	b, err := ReadUnsignedInt(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLong(v int64) {
	WriteLong(&stream.Buffer, v)
}

func (stream *Stream) GetLong() int64 {
	b, err := ReadLong(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutUnsignedLong(v uint64) {
	WriteUnsignedLong(&stream.Buffer, v)
}

func (stream *Stream) GetUnsignedLong() uint64 {
	b, err := ReadUnsignedLong(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutFloat(v float32) {
	WriteFloat(&stream.Buffer, v)
}

func (stream *Stream) GetFloat() float32 {
	b, err := ReadFloat(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutDouble(v float64) {
	WriteDouble(&stream.Buffer, v)
}

func (stream *Stream) GetDouble() float64 {
	b, err := ReadDouble(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutVarInt(v int32) {
	WriteVarInt(&stream.Buffer, v)
}

func (stream *Stream) GetVarInt() int32 {
	if stream.err != nil {
		return 0
	}
	i, err := ReadVarInt(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return i
}

func (stream *Stream) PutVarLong(v int64) {
	WriteVarLong(&stream.Buffer, v)
}

func (stream *Stream) GetVarLong() int64 {
	if stream.err != nil {
		return 0
	}
	i, err := ReadVarLong(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return i
}

func (stream *Stream) PutUnsignedVarInt(v uint32) {
	WriteUnsignedVarInt(&stream.Buffer, v)
}

func (stream *Stream) GetUnsignedVarInt() uint32 {
	if stream.err != nil {
		return 0
	}
	i, err := ReadUnsignedVarInt(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return i
}

func (stream *Stream) PutUnsignedVarLong(v uint64) {
	WriteUnsignedVarLong(&stream.Buffer, v)
}

func (stream *Stream) GetUnsignedVarLong() uint64 {
	if stream.err != nil {
		return 0
	}
	i, err := ReadUnsignedVarLong(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return i
}

func (stream *Stream) PutString(v string) {
	WriteUnsignedVarInt(&stream.Buffer, uint32(len(v)))
	stream.Buffer = append(stream.Buffer, []byte(v)...)
}

func (stream *Stream) GetString() string {
	if stream.err != nil {
		return ""
	}
	i, err := ReadString(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return i
}

func (stream *Stream) PutLittleShort(v int16) {
	WriteLittleShort(&stream.Buffer, v)
}

func (stream *Stream) GetLittleShort() int16 {
	b, err := ReadLittleShort(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLittleUnsignedShort(v uint16) {
	WriteLittleUnsignedShort(&stream.Buffer, v)
}

func (stream *Stream) GetLittleUnsignedShort() uint16 {
	b, err := ReadLittleUnsignedShort(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLittleInt(v int32) {
	WriteLittleInt(&stream.Buffer, v)
}

func (stream *Stream) GetLittleInt() int32 {
	b, err := ReadLittleInt(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLittleUnsignedInt(v uint32) {
	WriteLittleUnsignedInt(&stream.Buffer, v)
}

func (stream *Stream) GetLittleUnsignedInt() uint32 {
	b, err := ReadLittleUnsignedInt(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLittleLong(v int64) {
	WriteLittleLong(&stream.Buffer, v)
}

func (stream *Stream) GetLittleLong() int64 {
	b, err := ReadLittleLong(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLittleUnsignedLong(v uint64) {
	WriteLittleUnsignedLong(&stream.Buffer, v)
}

func (stream *Stream) GetLittleUnsignedLong() uint64 {
	b, err := ReadLittleUnsignedLong(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLittleFloat(v float32) {
	WriteLittleFloat(&stream.Buffer, v)
}

func (stream *Stream) GetLittleFloat() float32 {
	b, err := ReadLittleFloat(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLittleDouble(v float64) {
	WriteLittleDouble(&stream.Buffer, v)
}

func (stream *Stream) GetLittleDouble() float64 {
	b, err := ReadLittleDouble(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutTriad(v uint32) {
	WriteBigTriad(&stream.Buffer, v)
}

func (stream *Stream) GetTriad() uint32 {
	b, err := ReadBigTriad(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutLittleTriad(v uint32) {
	WriteLittleTriad(&stream.Buffer, v)
}

func (stream *Stream) GetLittleTriad() uint32 {
	b, err := ReadLittleTriad(&stream.Buffer, &stream.Offset)
	if err != nil {
		stream.SetError(err)
	}
	return b
}

func (stream *Stream) PutBytes(bytes []byte) {
	stream.Buffer = append(stream.Buffer, bytes...)
}

func (stream *Stream) PutLengthPrefixedBytes(bytes []byte) {
	stream.PutUnsignedVarInt(uint32(len(bytes)))
	stream.PutBytes(bytes)
}

func (stream *Stream) GetLengthPrefixedBytes() []byte {
	return []byte(stream.GetString())
}

func (stream *Stream) ResetStream() {
	stream.Offset = 0
	stream.Buffer = []byte{}
}
