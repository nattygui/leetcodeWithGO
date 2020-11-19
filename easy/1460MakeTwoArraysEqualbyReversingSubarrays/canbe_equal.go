package canbeequal

import (
	"reflect"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(arr)
	sort.Ints(target)
	return reflect.DeepEqual(target, arr)
}
