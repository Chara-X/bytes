package bytes

import (
	"bytes"
	"io"
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
	return &Buffer{buf: buf}
}
func (b *Buffer) Read(p []byte) (n int, err error) {
	if Reference {
		return b.b.Read(p)
	}
	if b.off >= len(b.buf) {
		b.buf, b.off = b.buf[:0], 0
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	return n, nil
}
func (b *Buffer) Write(p []byte) (n int, err error) {
	if Reference {
		return b.b.Write(p)
	}
	if len(b.buf)+len(p) > cap(b.buf) {
		copy(b.buf, b.buf[b.off:])
		b.buf = b.buf[:len(b.buf)-b.off]
	}
	b.buf = append(b.buf, p...)
	return len(p), nil
}
func (b *Buffer) Bytes() []byte {
	if Reference {
		return b.b.Bytes()
	}
	return b.buf[b.off:]
}
