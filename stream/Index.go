package stream

import (
	"errors"
	"qianxi/lambda-go/stream/utils"
)

func (o *Object) Index(i int) *Object {
	arr, err := utils.ToInterfaceList(o.obj)
	if err != nil {
		return o.clone(nil, err)
	}
	if i < 0 || i >= len(arr) {
		return o.clone(nil, errors.New("index out of range"))
	}

	return o.clone(arr[i], err)
}
