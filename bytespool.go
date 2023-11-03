package mempool

import "github.com/etfzy/mempool/base"

func NewBytesPool(levels []uint64) base.LevelsMemPool[byte] {
	return base.NewMemPool[byte](levels)
}
