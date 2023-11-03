package base

import (
	"sort"
	"sync"
)

type LevelsMemPool[T any] interface {
	Get(expect uint64) *[]T
	PutBack(b *[]T)
}

func NewMemPool[T any](levels []uint64) LevelsMemPool[T] {
	newlevel := RemoveRepByMap[uint64](levels)
	sort.Slice(newlevel, func(i, j int) bool {
		return newlevel[i] < newlevel[j]
	})

	p := &LevelsPool[T]{
		maxSize: levels[len(newlevel)-1],
		sp:      make([]*sync.Pool, len(newlevel)),
		levels:  levels,
	}

	for k, v := range newlevel {
		temp := v
		p.sp[k] = &sync.Pool{}
		p.sp[k].New = func() any {
			var b = make([]T, 0, temp)
			return &b
		}

	}

	return p
}
