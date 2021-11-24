package stream

import (
	"fmt"
	"qianxi/lambda-go/stream/utils"
	"reflect"
)

func (o *Object) Err() error {
	return o.err
}

func (o *Object) Object(resp interface{}) {
	if o.err != nil {
		panic(o.err)
	}

	if reflect.TypeOf(resp).Kind() != reflect.Ptr {
		panic("args must be ptr")
	}

	t := reflect.TypeOf(resp).Elem()
	v := reflect.ValueOf(resp).Elem()

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	to := reflect.TypeOf(o.obj)
	tv := reflect.ValueOf(o.obj)

	for to.Kind() == reflect.Ptr {
		to = to.Elem()
		tv = tv.Elem()
	}

	// for i := 0; i < t.NumField(); i++ {
	// 	for j := 0; j < to.NumField(); j++ {
	// 		if t.Field(i).Name == to.Field(i).Name {
	// 			v.Field(i).Set(tv.Field(i))
	// 		}
	// 	}
	// }
	for i := 0; i < t.NumField(); i++ {
		v.Field(i).Set(tv.Field(i))
	}
}

func (o *Object) ObjectWithErr(resp interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("err: %v", r)
		}
	}()

	if o.err != nil {
		return fmt.Errorf("%v", o.err)
	}

	if reflect.TypeOf(resp).Kind() != reflect.Ptr {
		return fmt.Errorf("args mut be ptr")
	}

	t := reflect.TypeOf(resp).Elem()
	v := reflect.ValueOf(resp).Elem()

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	to := reflect.TypeOf(o.obj)
	tv := reflect.ValueOf(o.obj)

	for to.Kind() == reflect.Ptr {
		to = to.Elem()
		tv = tv.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		v.Field(i).Set(tv.Field(i))
	}

	return nil
}

func (o *Object) ObjectSlice(resp interface{}) {
	if o.err != nil {
		panic(o.err)
	}

	arr, err := utils.ToInterfaceList(o.obj)
	if err != nil {
		panic(err)
	}

	v := reflect.ValueOf(resp)
	t := reflect.TypeOf(resp)
	for t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Slice {
		panic("args must be ptr of slice")
	}

	for i := 0; i < len(arr); i++ {
		elem := reflect.New(t.Elem()).Elem()

		at := reflect.TypeOf(arr[i])
		av := reflect.ValueOf(arr[i])
		for at.Kind() == reflect.Ptr {
			at = at.Elem()
			av = av.Elem()
		}

		for i := 0; i < elem.NumField(); i++ {
			elem.Field(i).Set(av.Field(i))
		}

		v.Set(reflect.Append(v, elem))
	}

}

func (o *Object) ObjectSliceWithErr(resp interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("err: %v", r)
		}
	}()

	if o.err != nil {
		return fmt.Errorf("%v", o.err)
	}

	arr, err := utils.ToInterfaceList(o.obj)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	v := reflect.ValueOf(resp)
	t := reflect.TypeOf(resp)
	for t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Slice {
		return fmt.Errorf("args must be ptr of slice")
	}

	for i := 0; i < len(arr); i++ {
		elem := reflect.New(t.Elem()).Elem()

		at := reflect.TypeOf(arr[i])
		av := reflect.ValueOf(arr[i])
		for at.Kind() == reflect.Ptr {
			at = at.Elem()
			av = av.Elem()
		}

		for i := 0; i < elem.NumField(); i++ {
			elem.Field(i).Set(av.Field(i))
		}

		v.Set(reflect.Append(v, elem))
	}

	return nil
}
