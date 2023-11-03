package example

import (
	"fmt"
	"testing"

	"github.com/xpizy2020/mempool"
)

func TestBytes(t *testing.T) {
	t.Run("test levels bytes", func(t *testing.T) {
		sp := mempool.NewLevelsBytesPool([]uint64{2048, 1024, 1024, 4096})
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

	t.Run("test constant bytes", func(t *testing.T) {
		sp := mempool.NewConstantBytesPool(1024)
		f := sp.Get()

		if cap(*f) != 1024 {
			t.Errorf("expect length is error %d", cap(*f))
		}

		sp.PutBack(f)

		s := sp.Get()

		p1 := &s
		p2 := &f
		fmt.Println(p1)
		fmt.Println(p2)
	})
}
