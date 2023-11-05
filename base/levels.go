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

type LevelsPool[T any] struct {
	maxSize uint64
	sp      []*sync.Pool
	levels  []uint64
}

func (m *LevelsPool[T]) Get(expectLen uint64) *Buffer[T] {

	if expectLen > m.maxSize {
		return NewBuffer[T](int(expectLen))
	}
	index := m.findHierachicalIndex(uint64(expectLen))
	if index == len(m.levels) {
		return NewBuffer[T](int(expectLen))
	}

	sp := m.sp[index]

	b := sp.Get().(*Buffer[T])
	b.SetLen(expectLen)
	return b
}

func (m *LevelsPool[T]) PutBack(b *Buffer[T]) {
	if b == nil {
		return
	}
	ex := b.Cap()
	if ex > MaxSize {
		return
	}

	//清零放回
	b.Reset()

	index := m.findHierachicalIndex(uint64(ex))
	if index == len(m.levels) {
		return
	}

	sp := m.sp[index]
	sp.Put(b)
	return
}

func (m *LevelsPool[T]) findHierachicalIndex(el uint64) int {
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
