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

		if f.Cap() != 2048 {
			t.Errorf("expect capacity is error %d", f.Cap())
		}

		if f.Len() != 1035 {
			t.Errorf("expect length is error %d", f.Len())
		}

		sp.PutBack(f)

		s := sp.Get(1068)
		s.Reset()
		fmt.Println(s.Len())
	})
}
