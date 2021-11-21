package utils

import (
	"fmt"
	"reflect"
)

func ToIntArray(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(int(0)))).Elem()
	for i := 0; i < length; i++ {
		ii, err := ToInt(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]int", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(ii))
	}
	return vtt.Interface(), nil
}

func ToInt8Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(int8(0)))).Elem()
	for i := 0; i < length; i++ {
		i8, err := ToInt8(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]int8", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(i8))
	}
	return vtt.Interface(), nil
}

func ToInt16Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(int16(0)))).Elem()
	for i := 0; i < length; i++ {
		i16, err := ToInt16(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]int16", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(i16))
	}
	return vtt.Interface(), nil
}

func ToInt32Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(int32(0)))).Elem()
	for i := 0; i < length; i++ {
		i32, err := ToInt32(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]int32", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(i32))
	}
	return vtt.Interface(), nil
}

func ToInt64Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(int64(0)))).Elem()
	for i := 0; i < length; i++ {
		i64, err := ToInt64(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]int64", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(i64))
	}
	return vtt.Interface(), nil
}

func ToByteArray(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(byte(0)))).Elem()
	for i := 0; i < length; i++ {
		b, err := ToByte(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]byte", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(b))
	}
	return vtt.Interface(), nil
}

func ToRuneArray(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(rune(0)))).Elem()
	for i := 0; i < length; i++ {
		r, err := ToRune(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]rune", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(r))
	}
	return vtt.Interface(), nil
}

func ToUintArray(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(uint(0)))).Elem()
	for i := 0; i < length; i++ {
		u, err := ToUint(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]uint", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(u))
	}
	return vtt.Interface(), nil
}

func ToUint8Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(uint8(0)))).Elem()
	for i := 0; i < length; i++ {
		u8, err := ToUint8(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]uint8", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(u8))
	}
	return vtt.Interface(), nil
}

func ToUint16Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(uint16(0)))).Elem()
	for i := 0; i < length; i++ {
		u16, err := ToUint16(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]uint16", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(u16))
	}
	return vtt.Interface(), nil
}

func ToUint32Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(uint32(0)))).Elem()
	for i := 0; i < length; i++ {
		u32, err := ToUint32(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]uint32", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(u32))
	}
	return vtt.Interface(), nil
}

func ToUint64Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(uint64(0)))).Elem()
	for i := 0; i < length; i++ {
		u64, err := ToUint64(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]uint64", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(u64))
	}
	return vtt.Interface(), nil
}

func ToUintptrArray(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(uintptr(0)))).Elem()
	for i := 0; i < length; i++ {
		up, err := ToUintptr(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]uintptr", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(up))
	}
	return vtt.Interface(), nil
}

func ToFloat32Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(float32(0)))).Elem()
	for i := 0; i < length; i++ {
		f32, err := ToFloat32(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]float32", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(f32))
	}
	return vtt.Interface(), nil
}

func ToFloat64Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(float64(0)))).Elem()
	for i := 0; i < length; i++ {
		f64, err := ToFloat64(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]float64", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(f64))
	}
	return vtt.Interface(), nil
}

func ToBoolArray(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(bool(false)))).Elem()
	for i := 0; i < length; i++ {
		b, err := ToBool(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]bool", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(b))
	}
	return vtt.Interface(), nil
}

func ToComplex64Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(complex64(0)))).Elem()
	for i := 0; i < length; i++ {
		c64, err := ToComplex64(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]complex64", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(c64))
	}
	return vtt.Interface(), nil
}

func ToComplex128Array(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(complex128(0)))).Elem()
	for i := 0; i < length; i++ {
		c128, err := ToComplex128(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]complex128", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(c128))
	}
	return vtt.Interface(), nil
}

func ToStringArray(v interface{}) (interface{}, error) {
	arr, err := ToInterfaceList(v)
	if err != nil {
		return nil, err
	}

	length := len(arr)
	vtt := reflect.New(reflect.ArrayOf(length, reflect.TypeOf(string("")))).Elem()
	for i := 0; i < length; i++ {
		s, err := ToString(arr[i])
		if err != nil {
			return nil, fmt.Errorf("%v[%T] can't convert to [%d]string", v, v, length)
		}
		vtt.Index(i).Set(reflect.ValueOf(s))
	}
	return vtt.Interface(), nil
}
