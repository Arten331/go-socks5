package socks5

import (
	"sync"
)

const BUFFER_SIZE = 4096

var BufferPool sync.Pool

func AcquireBuffer() []byte {
	v := BufferPool.Get()
	if v == nil {
		return make([]byte, BUFFER_SIZE)
	}
	return v.([]byte)
}

func ReleaseBuffer(buf []byte) {
	BufferPool.Put(buf)
}
