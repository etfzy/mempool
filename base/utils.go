package base

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap[T comparable](slc []T) []T {
	result := []T{}
	tempMap := map[T]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
