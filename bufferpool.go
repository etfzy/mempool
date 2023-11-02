package mempool

/*
type buffersPool struct {
	sp      []*sync.Pool
	maxSize int64
}

var buffersPoolInst *buffersPool

func init() {
	buffersPoolInst = NewBufferPool()
}


func NewBufferPool() *buffersPool {
	p := &buffersPool{
		maxSize: MaxSize,
		sp:      make([]*sync.Pool, len(LevelHierarchies)),
	}

	for k, v := range LevelHierarchies {
		temp := v
		p.sp[k] = &sync.Pool{}
		p.sp[k].New = func() any {
			var b = make([]byte, 0, temp)
			buff := bytes.NewBuffer(b)
			return buff
		}

	}

	return p
}
*/

/*
func GetBuffersPool(expectLen int64) *bytes.Buffer {
	cap := expectLen * 2
	if expectLen > MaxSize {
		var b = make([]byte, 0, cap)
		buff := bytes.NewBuffer(b)
		return buff
	}
	index := findHierachicalIndex(expectLen)
	if index == NoneLevelIndex {
		var b = make([]byte, 0, cap)
		buff := bytes.NewBuffer(b)
		return buff
	}

	sp := buffersPoolInst.sp[index]

	b := sp.Get().(*bytes.Buffer)
	return b
}
*/

/*
func PutBufferPool(b *bytes.Buffer) {
	if b == nil {
		return
	}
	ex := len(b.Bytes())
	if ex > MaxSize {
		return
	}

	//清零放回
	b.Reset()
	index := findHierachicalIndex(int64(ex))
	if index == NoneLevelIndex {
		return
	}
	sp := buffersPoolInst.sp[index]
	sp.Put(b)
	return
}
*/
