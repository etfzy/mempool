package base

import (
	"reflect"
	"unsafe"
)

type Buffer[T any] struct {
	buf []T
}

func (b *Buffer[T]) SetLen(length uint64) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b.buf))
	sliceHeader.Len = int(length)
}

func (b *Buffer[T]) Reset() {
	b.buf = b.buf[:0]
}

func (b *Buffer[T]) Cap() int {
	return cap(b.buf)
}

func (b *Buffer[T]) Len() int {
	return len(b.buf)
}

func (b *Buffer[T]) Buf() *[]T {
	return &b.buf
}

func NewBuffer[T any](capacity int) *Buffer[T] {
	return &Buffer[T]{
		buf: make([]T, 0, capacity),
	}
}
