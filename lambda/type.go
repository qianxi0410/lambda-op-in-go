package lambda

import "reflect"

type objType int

const (
	othert objType = 0
	arrt   objType = 1
	mapt   objType = 2
)

func (o *Object) objType() objType {
	vv := reflect.ValueOf(o.obj)

	switch vv.Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		return arrt
	case reflect.Map:
		return mapt
	default:
		return othert
	}
}
