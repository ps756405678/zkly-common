package slice

// 数组过滤
func Filter[T any](sli []T, test func(T, int) bool) []T {
	var result []T

	for idx, e := range sli {
		if test(e, idx) {
			result = append(result, e)
		}
	}

	return result
}

// 数组映射
func MapTo[O any, R any](sli []O, mapFunc func(O, int) R) []R {
	var result []R

	for idx, e := range sli {
		result = append(result, mapFunc(e, idx))
	}

	return result
}

// 数组任意匹配
func AnyMatch[T any](sli []T, test func(T, int) bool) bool {
	for idx, e := range sli {
		if test(e, idx) {
			return true
		}
	}
	return false
}

// 数组转Map
func ToMap[E any, K comparable, V any](sli []E, keyFun func(E) K, valFun func(E) V) map[K]V {
	var m = make(map[K]V)
	for _, e := range sli {
		m[keyFun(e)] = valFun(e)
	}
	return m
}

// 数组降维
func Reduce[E any](sli []E, accumulatorm func(E, E) E) E {
	var result E
	for i := 1; i < len(sli); i++ {
		result = accumulatorm(result, sli[i])
	}

	return result
}

func ForEach[E any](sli []E, consumer func(E, int)) {
	for idx, e := range sli {
		consumer(e, idx)
	}
}

func Sort[E any](sli *[]E, comparer func(*E, *E) int) {
	qs[E](sli, comparer, 0, len(*sli)-1)
}

func qs[E any](sli *[]E, comparer func(*E, *E) int, start int, end int) {
	if start >= end {
		return
	}
	var l = start
	var r = end
	cen := (*sli)[l]
	// 23154
	for l < r {
		for comparer(&(*sli)[r], &cen) >= 0 && l < r {
			r--
		}
		if l < r {
			swap[E](sli, l, r)
		}
		for comparer(&(*sli)[l], &cen) <= 0 && l < r {
			l++
		}
		if l < r {
			swap[E](sli, l, r)
		}
	}

	qs[E](sli, comparer, l+1, end)
	qs[E](sli, comparer, start, l-1)
}

func swap[E any](sli *[]E, i int, j int) {
	tmp := (*sli)[i]
	(*sli)[i] = (*sli)[j]
	(*sli)[j] = tmp
}
