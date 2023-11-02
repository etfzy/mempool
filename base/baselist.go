package base

import (
	"sync"
)

// buffer pool中包含了多个不同规模的内存池:
// 512字节、1KB、2KB、4KB、8KB、16KB、32KB，
// 超过32KB的内存不放入pool中
const (
	level1  = 512
	level2  = 1 * 1024
	level3  = 2 * 2048
	level4  = 4 * 1024
	level5  = 8 * 1024
	level6  = 16 * 1024
	level7  = 32 * 1024
	MaxSize = 32 * 1024
)

var DefaultLevelHierarchies = []int64{level1, level2, level3, level4, level5, level6, level7}

type ListMemPool[T any] interface {
	Get(expectLen int64) *[]T
	PutBack(b *[]T)
}

type MemPoolInfos[T any] struct {
	maxSize int64
	sp      []*sync.Pool
	levels  []uint64
}

func NewMemPool[T any](maxSize int, levels []uint64) ListMemPool[T] {
	p := &MemPoolInfos[T]{
		maxSize: MaxSize,
		sp:      make([]*sync.Pool, len(levels)),
		levels:  levels,
	}

	for k, v := range levels {
		temp := v
		p.sp[k] = &sync.Pool{}
		p.sp[k].New = func() any {
			var b = make([]T, 0, temp)
			return &b
		}

	}

	return p
}

func (m *MemPoolInfos[T]) Get(expectLen int64) *[]T {

	if expectLen > m.maxSize {
		b := make([]T, 0, expectLen)
		return &b
	}
	index := m.findHierachicalIndex(uint64(expectLen))
	if index == len(m.levels) {
		b := make([]T, 0, expectLen)
		return &b
	}

	sp := m.sp[index]

	b := sp.Get().(*[]T)
	return b
}

func (m *MemPoolInfos[T]) PutBack(b *[]T) {
	if b == nil {
		return
	}
	ex := cap(*b)
	if ex > MaxSize {
		return
	}

	//清零放回
	*b = (*b)[:0]

	index := m.findHierachicalIndex(uint64(ex))
	if index == len(m.levels) {
		return
	}

	sp := m.sp[index]
	sp.Put(b)
	return
}

func (m *MemPoolInfos[T]) findHierachicalIndex(el uint64) int {
	//level 索引初始化为7（有效值为0-6）
	index := len(m.levels)
	for k, v := range m.levels {
		if el > v {
			continue
		} else {
			//找到等级break，否则跳入下一级
			index = k
			break
		}
	}

	return index
}
