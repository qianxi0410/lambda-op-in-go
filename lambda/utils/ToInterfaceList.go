package utils

import (
	"fmt"
	"reflect"
)

func ToInterfaceList(v interface{}) (res []interface{}, err error) {
	vv := reflect.ValueOf(v)
	switch vv.Kind() {
	case reflect.String:
		for _, v := range vv.String() {
			res = append(res, v)
		}
		return res, nil
	case reflect.Slice:
		for i := 0; i < vv.Len(); i++ {
			res = append(res, vv.Index(i).Interface())
		}
		return res, nil
	case reflect.Array:
		for i := 0; i < vv.Len(); i++ {
			res = append(res, vv.Index(i).Interface())
		}
		return res, nil
	default:
		return nil, fmt.Errorf("%v[%T] can't convert to []interface", v, v)
	}
}
