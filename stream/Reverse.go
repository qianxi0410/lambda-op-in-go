package stream

import "qianxi/lambda-go/stream/utils"

func (o *Object) Reverse() *Object {
	arr, err := utils.ToInterfaceList(o.obj)
	if err != nil {
		return o.clone(nil, err)
	}

	objs := make([]interface{}, len(arr))

	for i := 0; i < len(arr); i++ {
		objs[i] = arr[len(arr)-i-1]
	}

	return o.clone(objs, nil)
}
