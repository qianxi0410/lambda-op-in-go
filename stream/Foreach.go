package stream

import "qianxi/lambda-go/stream/utils"

func (o *Object) Foreach(fn func(e Elem)) {
	list, err := utils.ToInterfaceList(o.obj)
	if err != nil {
		panic(err)
	}

	for i := range list {
		fn(list[i])
	}
}
