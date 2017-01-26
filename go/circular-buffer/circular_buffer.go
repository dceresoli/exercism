package circular

import "errors"

const testVersion = 4

type Buffer struct {
	data       []byte
	size       int
	start, end int
	ndata      int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: make([]byte, size), size: size, start: 0, end: 0, ndata: 0}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.ndata <= 0 {
		return 0, errors.New("can't read from empty buffer")
	}
	res := b.data[b.start]
	b.start = (b.start + 1) % b.size
	b.ndata--
	return res, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.ndata == b.size {
		return errors.New("can't write: buffer full")
	}
	b.data[b.end] = c
	b.end = (b.end + 1) % b.size
	b.ndata++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	b.data[b.start] = c
	b.start = (b.start + 1) % b.size
}

func (b *Buffer) Reset() {
	b.ndata = 0
	b.start = 0
	b.end = 0
}
