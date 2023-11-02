package example

import (
	"fmt"
	"testing"

	"github.com/xpizy2020/mempool"
)

func TestBytes(t *testing.T) {
	t.Run("test bytes", func(t *testing.T) {
		sp := mempool.NewBytesPool(32*1024, []uint64{1024, 2048, 4096})
		f := sp.Get(1035)

		if cap(*f) != 2048 {
			t.Errorf("expect length is error %d", cap(*f))
		}

		sp.PutBack(f)

		s := sp.Get(2048)

		p1 := &s
		p2 := &f
		fmt.Println(p1)
		fmt.Println(p2)
	})
}
