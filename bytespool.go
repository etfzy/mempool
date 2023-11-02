package mempool

import "github.com/xpizy2020/mempool/base"

func NewBytesPool(maxSize int, levels []uint64) base.ListMemPool[byte] {
	return base.NewMemPool[byte](maxSize, levels)
}
