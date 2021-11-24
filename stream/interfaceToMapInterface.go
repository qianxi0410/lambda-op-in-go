package stream

import (
	"fmt"
	"reflect"
)

func interfaceToInterfaceMap(v interface{}) (res map[interface{}]interface{}, err error) {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return nil, fmt.Errorf("%T unsupport to map lambda operator", v)
	}

	m := vv.MapRange()
	res = map[interface{}]interface{}{}
	for m.Next() {
		res[m.Key().Interface()] = m.Value().Interface()
	}
	return
}
