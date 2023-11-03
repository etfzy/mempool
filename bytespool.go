package mempool

import "github.com/xpizy2020/mempool/base"

func NewLevelsBytesPool(levels []uint64) base.LevelsMemPool[byte] {
	return base.NewLevelsMemPool[byte](levels)
}

func NewConstantBytesPool(expect uint64) base.ConstantMemPool[byte] {
	return base.NewConstantPool[byte](expect)
}
