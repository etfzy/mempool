package mempool

import (
	"sync"

	"github.com/xpizy2020/mempool/base"
)

type ConstantMemPool[T any] interface {
	Get() *[]T
	PutBack(b *[]T)
}

type LevelsMemPool[T any] interface {
	Get(expect uint64) *[]T
	PutBack(b *[]T)
}

func NewLevelsMemPool[T any](levels []uint64) LevelsMemPool[T] {
	p := &base.LevelsPool[T]{
		maxSize: levels[len(levels)-1],
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

func NewConstantPool[T any](expect uint64) ConstantMemPool[T] {
	p := &base.ConstantPool[T]{
		expect: expect,
		sp:     &sync.Pool{},
	}

	p.sp.New = func() any {
		var b = make([]T, 0, expect)
		return &b
	}

	return p
}
