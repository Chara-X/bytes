package bytes

import (
	"bytes"
)

type Buffer struct {
	b   *bytes.Buffer
	buf []byte
	off int
}

func NewBuffer(buf []byte) *Buffer {
	if Reference {
		return &Buffer{b: bytes.NewBuffer(buf)}
	}
	panic("not implemented")
}
func (b *Buffer) Read(p []byte) (n int, err error) {
	if Reference {
		return b.b.Read(p)
	}
	panic("not implemented")
}
func (b *Buffer) Write(p []byte) (n int, err error) {
	if Reference {
		return b.b.Write(p)
	}
	panic("not implemented")
}
func (b *Buffer) Bytes() []byte {
	if Reference {
		return b.b.Bytes()
	}
	return b.buf[b.off:]
}
