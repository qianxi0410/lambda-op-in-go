package lambda

import "qianxi/lambda-go/lambda/utils"

func (o *Object) Int() int {
	if o.err != nil {
		panic("To int panic")
	}
	i, err := utils.ToInt(o.obj)
	if err != nil {
		panic("To int panic")
	}
	return i
}

func (o *Object) IntWithErr() (int, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToInt(o.obj)
}

func (o *Object) Int8() int8 {
	if o.err != nil {
		panic("To int8 panic")
	}
	i8, err := utils.ToInt8(o.obj)
	if err != nil {
		panic("To int8 panic")
	}
	return i8
}

func (o *Object) Int8WithErr() (int8, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToInt8(o.obj)
}

func (o *Object) Int16() int16 {
	if o.err != nil {
		panic("To int16 panic")
	}
	i16, err := utils.ToInt16(o.obj)
	if err != nil {
		panic("To int16 panic")
	}
	return i16
}

func (o *Object) Int16WithErr() (int16, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToInt16(o.obj)
}

func (o *Object) Int32() int32 {
	if o.err != nil {
		panic("To int32 panic")
	}
	i32, err := utils.ToInt32(o.obj)
	if err != nil {
		panic("To int32 panic")
	}
	return i32
}

func (o *Object) Int32WithErr() (int32, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToInt32(o.obj)
}

func (o *Object) Int64() int64 {
	if o.err != nil {
		panic("To int panic")
	}
	i64, err := utils.ToInt64(o.obj)
	if err != nil {
		panic("To int panic")
	}
	return i64
}

func (o *Object) Int64WithErr() (int64, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToInt64(o.obj)
}

func (o *Object) Uint() uint {
	if o.err != nil {
		panic("To uint panic")
	}
	u, err := utils.ToUint(o.obj)
	if err != nil {
		panic("To uint panic")
	}
	return u
}

func (o *Object) UintWithErr() (uint, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToUint(o.obj)
}

func (o *Object) Uint8() uint8 {
	if o.err != nil {
		panic("To uint8 panic")
	}
	u8, err := utils.ToUint8(o.obj)
	if err != nil {
		panic("To uint8 panic")
	}
	return u8
}

func (o *Object) Uint8WithErr() (uint8, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToUint8(o.obj)
}

func (o *Object) Uint16() uint16 {
	if o.err != nil {
		panic("To uint16 panic")
	}
	u16, err := utils.ToUint16(o.obj)
	if err != nil {
		panic("To uint16 panic")
	}
	return u16
}

func (o *Object) Uint16WithErr() (uint16, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToUint16(o.obj)
}

func (o *Object) Uint32() uint32 {
	if o.err != nil {
		panic("To uint32 panic")
	}
	u32, err := utils.ToUint32(o.obj)
	if err != nil {
		panic("To uint32 panic")
	}
	return u32
}

func (o *Object) Uint32WithErr() (uint32, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToUint32(o.obj)
}

func (o *Object) Uint64() uint64 {
	if o.err != nil {
		panic("To uint panic")
	}
	u64, err := utils.ToUint64(o.obj)
	if err != nil {
		panic("To uint panic")
	}
	return u64
}

func (o *Object) Uint64WithErr() (uint64, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToUint64(o.obj)
}

func (o *Object) Uintptr() uintptr {
	if o.err != nil {
		panic("To uintptr panic")
	}
	up, err := utils.ToUintptr(o.obj)
	if err != nil {
		panic("To uintptr panic")
	}
	return up
}

func (o *Object) UintptrWithErr() (uintptr, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToUintptr(o.obj)
}

func (o *Object) Float32() float32 {
	if o.err != nil {
		panic("To float32 panic")
	}
	f32, err := utils.ToFloat32(o.obj)
	if err != nil {
		panic("To float32 panic")
	}
	return f32
}

func (o *Object) Float32WithErr() (float32, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToFloat32(o.obj)
}

func (o *Object) Float64() float64 {
	if o.err != nil {
		panic("To float64 panic")
	}
	f64, err := utils.ToFloat64(o.obj)
	if err != nil {
		panic("To float64 panic")
	}
	return f64
}

func (o *Object) Float64WithErr() (float64, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToFloat64(o.obj)
}

func (o *Object) Bool() bool {
	if o.err != nil {
		panic("To bool panic")
	}
	b, err := utils.ToBool(o.obj)
	if err != nil {
		panic("To bool panic")
	}
	return b
}

func (o *Object) BoolWithErr() (bool, error) {
	if o.err != nil {
		return false, o.err
	}
	return utils.ToBool(o.obj)
}

func (o *Object) Complex64() complex64 {
	if o.err != nil {
		panic("To complex64 panic")
	}
	c64, err := utils.ToComplex64(o.obj)
	if err != nil {
		panic("To complex64 panic")
	}
	return c64
}

func (o *Object) Complex64WithErr() (complex64, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToComplex64(o.obj)
}

func (o *Object) Complex128() complex128 {
	if o.err != nil {
		panic("To complex128 panic")
	}
	c128, err := utils.ToComplex128(o.obj)
	if err != nil {
		panic("To complex128 panic")
	}
	return c128
}

func (o *Object) Complex128WithErr() (complex128, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToComplex128(o.obj)
}

func (o *Object) String() string {
	if o.err != nil {
		panic("To string panic")
	}
	s, err := utils.ToString(o.obj)
	if err != nil {
		panic("To string panic")
	}
	return s
}

func (o *Object) StringWithErr() (string, error) {
	if o.err != nil {
		return "", o.err
	}
	return utils.ToString(o.obj)
}

func (o *Object) Byte() byte {
	if o.err != nil {
		panic("To byte panic")
	}
	b, err := utils.ToByte(o.obj)
	if err != nil {
		panic("To byte panic")
	}
	return b
}

func (o *Object) ByteWithErr() (byte, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToByte(o.obj)
}

func (o *Object) Rune() rune {
	if o.err != nil {
		panic("To rune panic")
	}
	r, err := utils.ToRune(o.obj)
	if err != nil {
		panic("To rune panic")
	}
	return r
}

func (o *Object) RuneWithErr() (rune, error) {
	if o.err != nil {
		return 0, o.err
	}
	return utils.ToRune(o.obj)
}
