package utils

import (
	"fmt"
	"reflect"
)

func ToIntSlice(v interface{}) (res []int, err error) {
	switch v := v.(type) {
	case []int:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				ii, err := ToInt(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, int(ii))
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []int", v, v)
	}
}

func ToInt8Slice(v interface{}) (res []int8, err error) {
	switch v := v.(type) {
	case []int8:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt8(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				i8, err := ToInt8(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, i8)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []int8", v, v)
	}
}

func ToInt16Slice(v interface{}) (res []int16, err error) {
	switch v := v.(type) {
	case []int16:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt16(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				i16, err := ToInt16(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, i16)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []int16", v, v)
	}
}

func ToInt32Slice(v interface{}) (res []int32, err error) {
	switch v := v.(type) {
	case []int32:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt32(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				i32, err := ToInt32(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, i32)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []int32", v, v)
	}
}

func ToInt64Slice(v interface{}) (res []int64, err error) {
	switch v := v.(type) {
	case []int64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt64(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				i64, err := ToInt64(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, i64)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []int64", v, v)
	}
}

func ToByteSlice(v interface{}) (res []byte, err error) {
	switch v := v.(type) {
	case []byte:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToByte(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				ib, err := ToByte(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, ib)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []byte", v, v)
	}
}

func ToRuneSlice(v interface{}) (res []rune, err error) {
	switch v := v.(type) {
	case []rune:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToRune(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				ir, err := ToRune(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, ir)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []rune", v, v)
	}
}

func ToUintSlice(v interface{}) (res []uint, err error) {
	switch v := v.(type) {
	case []uint:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				u, err := ToUint(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, u)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []uint", v, v)
	}
}

func ToUint8Slice(v interface{}) (res []uint8, err error) {
	switch v := v.(type) {
	case []uint8:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint8(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				u8, err := ToUint8(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, u8)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []uint8", v, v)
	}
}

func ToUint16Slice(v interface{}) (res []uint16, err error) {
	switch v := v.(type) {
	case []uint16:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint16(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				u16, err := ToUint16(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, u16)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []uint16", v, v)
	}
}

func ToUint32Slice(v interface{}) (res []uint32, err error) {
	switch v := v.(type) {
	case []uint32:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint32(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				u32, err := ToUint32(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, u32)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []uint32", v, v)
	}
}

func ToUint64Slice(v interface{}) (res []uint64, err error) {
	switch v := v.(type) {
	case []uint64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint64(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				u64, err := ToUint64(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, u64)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []uint64", v, v)
	}
}

func ToUintptrSlice(v interface{}) (res []uintptr, err error) {
	switch v := v.(type) {
	case []uintptr:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUintptr(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		flag := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				up, err := ToUintptr(vv.Index(i).Interface())
				if err != nil {
					flag = true
					break
				}
				res = append(res, up)
			}
			if !flag {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []uintptr", v, v)
	}
}

func ToFloat32Slice(v interface{}) (res []float32, err error) {
	switch v := v.(type) {
	case []float32:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToFloat32(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		catchErr := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				f32, err := ToFloat32(vv.Index(i).Interface())
				if err != nil {
					catchErr = true
					break
				}
				res = append(res, f32)
			}
			if !catchErr {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []float32", v, v)
	}
}

func ToFloat64Slice(v interface{}) (res []float64, err error) {
	switch v := v.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToFloat64(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		catchErr := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				f64, err := ToFloat64(vv.Index(i).Interface())
				if err != nil {
					catchErr = true
					break
				}
				res = append(res, f64)
			}
			if !catchErr {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []float64", v, v)
	}
}

func ToBoolSlice(v interface{}) (res []bool, err error) {
	switch v := v.(type) {
	case []bool:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToBool(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		catchErr := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				b, err := ToBool(vv.Index(i).Interface())
				if err != nil {
					catchErr = true
					break
				}
				res = append(res, b)
			}
			if !catchErr {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []bool", v, v)
	}
}

func ToComplex64Slice(v interface{}) (res []complex64, err error) {
	switch v := v.(type) {
	case []complex64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToComplex64(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		catchErr := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				c64, err := ToComplex64(vv.Index(i).Interface())
				if err != nil {
					catchErr = true
					break
				}
				res = append(res, c64)
			}
			if !catchErr {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []complex64", v, v)
	}
}

func ToComplex128Slice(v interface{}) (res []complex128, err error) {
	switch v := v.(type) {
	case []complex128:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToComplex128(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		catchErr := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				c128, err := ToComplex128(vv.Index(i).Interface())
				if err != nil {
					catchErr = true
					break
				}
				res = append(res, c128)
			}
			if !catchErr {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []complex128", v, v)
	}
}

func ToStringSlice(v interface{}) (res []string, err error) {
	switch v := v.(type) {
	case []string:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToString(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		vv := reflect.ValueOf(v)
		catchErr := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				s, err := ToString(vv.Index(i).Interface())
				if err != nil {
					catchErr = true
					break
				}
				res = append(res, s)
			}
			if !catchErr {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v[%T] can't convert to []string", v, v)
	}
}
