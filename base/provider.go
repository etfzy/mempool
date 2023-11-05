package base

import (
	"sort"
	"sync"
)

type LevelsMemPool[T any] interface {
	Get(expect uint64) *Buffer[T]
	PutBack(b *Buffer[T])
}

func NewMemPool[T any](inputs []uint64) LevelsMemPool[T] {
	newlevel := RemoveRepByMap[uint64](inputs)
	sort.Slice(newlevel, func(i, j int) bool {
		return newlevel[i] < newlevel[j]
	})

	p := &LevelsPool[T]{
		maxSize: newlevel[len(newlevel)-1],
		sp:      make([]*sync.Pool, len(newlevel)),
		levels:  newlevel,
	}

	for k, v := range newlevel {
		temp := v
		p.sp[k] = &sync.Pool{}
		p.sp[k].New = func() any {
			return newBuffer[T](int(temp))
		}
	}

	return p
}
