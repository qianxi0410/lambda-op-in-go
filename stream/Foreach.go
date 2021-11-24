package lambda

import "qianxi/lambda-go/lambda/utils"

func (o *Object) Foreach(fn func(e Elem)) {
	list, err := utils.ToInterfaceList(o.obj)
	if err != nil {
		panic("can's iterator non-list obj")
	}

	for i := range list {
		fn(list[i])
	}
}
