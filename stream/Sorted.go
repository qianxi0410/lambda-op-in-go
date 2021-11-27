package stream

import (
	"qianxi/lambda-go/stream/utils"
	"sort"
)

func (o *Object) Sorted(f func(e1, e2 Elem) bool) *Object {
	arr, err := utils.ToInterfaceList(o.obj)
	if err != nil {
		return o.clone(nil, err)
	}

	sort.Slice(arr, func(i, j int) bool {
		return f(arr[i], arr[j])
	})

	return o.clone(arr, err)
}
