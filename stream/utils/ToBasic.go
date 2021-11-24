package utils

import (
	"fmt"
	"math"
)

func ToInt8(v interface{}) (int8, error) {
	switch v := v.(type) {
	case int:
		if v >= math.MinInt8 && v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case int8:
		return v, nil
	case int16:
		if v >= math.MinInt8 && v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case int32:
		if v >= math.MinInt8 && v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case int64:
		if v >= math.MinInt8 && v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case uint:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case uint8:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case uint16:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case uint32:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case uint64:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	case uintptr:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int8", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to int8", v)
	}
}

func ToInt16(v interface{}) (int16, error) {
	switch v := v.(type) {
	case int:
		if v >= math.MinInt16 && v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int16", v, v)
	case int8:
		return int16(v), nil
	case int16:
		return v, nil
	case int32:
		if v >= math.MinInt16 && v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int16", v, v)
	case int64:
		if v >= math.MinInt16 && v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int16", v, v)
	case uint:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int16", v, v)
	case uint8:
		return int16(v), nil
	case uint16:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int16", v, v)
	case uint32:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int16", v, v)
	case uint64:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int16", v, v)
	case uintptr:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int16", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to int16", v)
	}
}

func ToInt32(v interface{}) (int32, error) {
	switch v := v.(type) {
	case int:
		if v >= math.MinInt32 && v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int32", v, v)
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		if v >= math.MinInt32 && v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int32", v, v)
	case uint:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int32", v, v)
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int32", v, v)
	case uint64:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int32", v, v)
	case uintptr:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int32", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to int32", v)
	}
}

func ToInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		if v <= math.MaxInt64 {
			return int64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int64", v, v)
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		if v <= math.MaxInt64 {
			return int64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int64", v, v)
	case uintptr:
		if v <= math.MaxInt64 {
			return int64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int64", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to int64", v)
	}
}

func ToInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		if v >= math.MinInt && v <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int", v, v)
	case uint:
		if v <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int", v, v)
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		if v <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int", v, v)
	case uintptr:
		if v <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of int", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to int", v)
	}
}

func ToUint8(v interface{}) (uint8, error) {
	switch v := v.(type) {
	case int:
		if v >= 0 && v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case int8:
		if v >= 0 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case int16:
		if v >= 0 && v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case int32:
		if v >= 0 && v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case int64:
		if v >= 0 && v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case uint:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case uint8:
		return v, nil
	case uint16:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case uint32:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case uint64:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	case uintptr:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint8", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to int8", v)
	}
}

func ToUint16(v interface{}) (uint16, error) {
	switch v := v.(type) {
	case int:
		if v >= 0 && v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	case int8:
		if v >= 0 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	case int16:
		if v >= 0 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	case int32:
		if v >= 0 && v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	case int64:
		if v >= 0 && v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	case uint:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	case uint8:
		return uint16(v), nil

	case uint16:
		return v, nil
	case uint32:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	case uint64:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	case uintptr:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint16", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to uint16", v)
	}
}

func ToUint32(v interface{}) (uint32, error) {
	switch v := v.(type) {
	case int:
		if v >= 0 && v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint32", v, v)
	case int8:
		if v >= 0 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint32", v, v)
	case int16:
		if v >= 0 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint32", v, v)
	case int32:
		if v >= 0 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint32", v, v)
	case int64:
		if v >= 0 && v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint32", v, v)
	case uint:
		if v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint32", v, v)
	case uint8:
		return uint32(v), nil

	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint64:
		if v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint32", v, v)
	case uintptr:
		if v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint32", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to uint32", v)
	}
}

func ToUint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case int:
		if v >= 0 {
			return uint64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint64", v, v)
	case int8:
		if v >= 0 {
			return uint64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint64", v, v)
	case int16:
		if v >= 0 {
			return uint64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint64", v, v)
	case int32:
		if v >= 0 {
			return uint64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint64", v, v)
	case int64:
		if v >= 0 {
			return uint64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint64", v, v)
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case uintptr:
		if v <= math.MaxUint64 {
			return uint64(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint64", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to uint64", v)
	}
}

func ToUint(v interface{}) (uint, error) {
	switch v := v.(type) {
	case int:
		if v >= 0 {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint", v, v)
	case int8:
		if v >= 0 {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint", v, v)
	case int16:
		if v >= 0 {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint", v, v)
	case int32:
		if v >= 0 {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint", v, v)
	case int64:
		if v >= 0 {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint", v, v)
	case uint:
		return v, nil
	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint64:
		if v <= math.MaxUint {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint", v, v)
	case uintptr:
		if v <= math.MaxUint {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uint", v, v)
	default:
		return 0, fmt.Errorf("%T can't convert to uint", v)
	}
}

func ToUintptr(v interface{}) (uintptr, error) {
	switch v := v.(type) {
	case int:
		if v >= 0 {
			return uintptr(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uintptr", v, v)
	case int8:
		if v >= 0 {
			return uintptr(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uintptr", v, v)
	case int16:
		if v >= 0 {
			return uintptr(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uintptr", v, v)
	case int32:
		if v >= 0 {
			return uintptr(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uintptr", v, v)
	case int64:
		if v >= 0 {
			return uintptr(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of uintptr", v, v)
	case uint:
		return uintptr(v), nil
	case uint8:
		return uintptr(v), nil
	case uint16:
		return uintptr(v), nil
	case uint32:
		return uintptr(v), nil
	case uint64:
		return uintptr(v), nil
	case uintptr:
		return v, nil
	default:
		return 0, fmt.Errorf("%T can't convert to uintptr", v)
	}
}

func ToFloat32(v interface{}) (float32, error) {
	switch v := v.(type) {
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case uintptr:
		return float32(v), nil
	case float32:
		return v, nil
	case float64:
		if v <= math.MaxFloat32 {
			return float32(v), nil
		}
		return 0, fmt.Errorf("%v[%T] beyound the scope of float32", v, v)
	default:
		return 0, fmt.Errorf("%v[%T] can't convert to float32", v, v)
	}
}

func ToFloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uintptr:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("%v[%T] can't convert to float64", v, v)
	}
}

func ToBool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("%v[%T] can't convert to bool", v, v)
	}
}

func ToComplex64(v interface{}) (complex64, error) {
	switch v := v.(type) {
	case complex64:
		return v, nil
	default:
		return 0, fmt.Errorf("%v[%T] can't convert to complex64", v, v)
	}
}

func ToComplex128(v interface{}) (complex128, error) {
	switch v := v.(type) {
	case complex64:
		return complex128(v), nil
	case complex128:
		return v, nil
	default:
		return 0, fmt.Errorf("%v[%T] can't convert to complex128", v, v)
	}
}

func ToString(v interface{}) (string, error) {
	switch v := v.(type) {
	case []rune:
		return string(v), nil
	case []byte:
		return string(v), nil
	case string:
		return v, nil
	case interface{}:
		resp, err := ToInterfaceList(v)
		if err != nil {
			return "", err
		}
		rs := make([]rune, 0, len(resp))
		for i := 0; i < len(resp); i++ {
			rr, ok := resp[i].(rune)
			if !ok {
				return "", fmt.Errorf("cant convert to rune")
			}
			rs = append(rs, rr)
		}
		return string(rs), nil
	default:
		return "", fmt.Errorf("%v[%T] can't convert to string", v, v)
	}
}

func ToByte(v interface{}) (byte, error) {
	val, err := ToInt8(v)
	if err != nil {
		return 0, err
	}
	return byte(val), nil
}

func ToRune(v interface{}) (rune, error) {
	val, err := ToInt32(v)
	if err != nil {
		return 0, err
	}
	return rune(val), nil
}
