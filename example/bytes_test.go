package example

import (
	"fmt"
	"testing"

	"github.com/etfzy/mempool"
)

func TestBytes(t *testing.T) {
	t.Run("test levels bytes", func(t *testing.T) {
		sp := mempool.NewBytesPool([]uint64{2048, 1024, 1024, 4096})
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
