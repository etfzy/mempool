package base

import (
	"sync"
)

// buffer pool中包含了多个不同规模的内存池:
// 512字节、1KB、2KB、4KB、8KB、16KB、32KB，
// 超过32KB的内存不放入pool中
type ConstantPool[T any] struct {
	expect uint64
	sp     *sync.Pool
}

func (m *ConstantPool[T]) Get() *[]T {
	b := m.sp.Get().(*[]T)
	return b
}

func (m *ConstantPool[T]) PutBack(b *[]T) {
	if uint64(cap(*b)) != m.expect {
		return
	}
	m.sp.Put(b)
	return
}
