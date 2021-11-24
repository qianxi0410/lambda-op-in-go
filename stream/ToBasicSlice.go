package stream

import "qianxi/lambda-go/stream/utils"

func (o *Object) IntSlice() []int {
	if o.err != nil {
		panic("To intslice panic")
	}

	res, err := utils.ToIntSlice(o.obj)
	if err != nil {
		panic("To intslice panic")
	}

	return res
}

func (o *Object) IntSliceWithErr() ([]int, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToIntSlice(o.obj)
}

func (o *Object) Int8Slice() []int8 {
	if o.err != nil {
		panic("To int8slice panic")
	}

	res, err := utils.ToInt8Slice(o.obj)
	if err != nil {
		panic("To int8slice panic")
	}

	return res
}

func (o *Object) Int8SliceWithErr() ([]int8, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToInt8Slice(o.obj)
}

func (o *Object) Int16Slice() []int16 {
	if o.err != nil {
		panic("To int16slice panic")
	}

	res, err := utils.ToInt16Slice(o.obj)
	if err != nil {
		panic("To int16slice panic")
	}

	return res
}

func (o *Object) Int16SliceWithErr() ([]int16, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToInt16Slice(o.obj)
}

func (o *Object) Int32Slice() []int32 {
	if o.err != nil {
		panic("To int32slice panic")
	}

	res, err := utils.ToInt32Slice(o.obj)
	if err != nil {
		panic("To int32slice panic")
	}

	return res
}

func (o *Object) Int32SliceWithErr() ([]int32, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToInt32Slice(o.obj)
}

func (o *Object) Int64Slice() []int64 {
	if o.err != nil {
		panic("To int64slice panic")
	}

	res, err := utils.ToInt64Slice(o.obj)
	if err != nil {
		panic("To int64slice panic")
	}

	return res
}

func (o *Object) Int64SliceWithErr() ([]int64, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToInt64Slice(o.obj)
}

func (o *Object) UintSlice() []uint {
	if o.err != nil {
		panic("To uintslice panic")
	}

	res, err := utils.ToUintSlice(o.obj)
	if err != nil {
		panic("To uintslice panic")
	}

	return res
}

func (o *Object) UintSliceWithErr() ([]uint, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToUintSlice(o.obj)
}

func (o *Object) Uint8Slice() []uint8 {
	if o.err != nil {
		panic("To uint8slice panic")
	}

	res, err := utils.ToUint8Slice(o.obj)
	if err != nil {
		panic("To uint8slice panic")
	}

	return res
}

func (o *Object) Uint8SliceWithErr() ([]uint8, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToUint8Slice(o.obj)
}

func (o *Object) Uint16Slice() []uint16 {
	if o.err != nil {
		panic("To uint16slice panic")
	}

	res, err := utils.ToUint16Slice(o.obj)
	if err != nil {
		panic("To uint16slice panic")
	}

	return res
}

func (o *Object) Uint16SliceWithErr() ([]uint16, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToUint16Slice(o.obj)
}

func (o *Object) Uint32Slice() []uint32 {
	if o.err != nil {
		panic("To uint32slice panic")
	}

	res, err := utils.ToUint32Slice(o.obj)
	if err != nil {
		panic("To uint32slice panic")
	}

	return res
}

func (o *Object) Uint32SliceWithErr() ([]uint32, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToUint32Slice(o.obj)
}

func (o *Object) Uint64Slice() []uint64 {
	if o.err != nil {
		panic("To uint64slice panic")
	}

	res, err := utils.ToUint64Slice(o.obj)
	if err != nil {
		panic("To uint64slice panic")
	}

	return res
}

func (o *Object) Uint64SliceWithErr() ([]uint64, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToUint64Slice(o.obj)
}

func (o *Object) UintptrSlice() []uintptr {
	if o.err != nil {
		panic("To uintptrslice panic")
	}

	res, err := utils.ToUintptrSlice(o.obj)
	if err != nil {
		panic("To uintptrslice panic")
	}

	return res
}

func (o *Object) UintptrSliceWithErr() ([]uintptr, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToUintptrSlice(o.obj)
}

func (o *Object) Float32Slice() []float32 {
	if o.err != nil {
		panic("To float32slice panic")
	}

	res, err := utils.ToFloat32Slice(o.obj)
	if err != nil {
		panic("To float32slice panic")
	}

	return res
}

func (o *Object) Float32SliceWithErr() ([]float32, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToFloat32Slice(o.obj)
}

func (o *Object) Float64Slice() []float64 {
	if o.err != nil {
		panic("To float64slice panic")
	}

	res, err := utils.ToFloat64Slice(o.obj)
	if err != nil {
		panic("To float64slice panic")
	}

	return res
}

func (o *Object) Float64SliceWithErr() ([]float64, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToFloat64Slice(o.obj)
}

func (o *Object) Complex64Slice() []complex64 {
	if o.err != nil {
		panic("To complex64slice panic")
	}

	res, err := utils.ToComplex64Slice(o.obj)
	if err != nil {
		panic("To complex64slice panic")
	}

	return res
}

func (o *Object) Complex64SliceWithErr() ([]complex64, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToComplex64Slice(o.obj)
}

func (o *Object) Complex128Slice() []complex128 {
	if o.err != nil {
		panic("To complex128slice panic")
	}

	res, err := utils.ToComplex128Slice(o.obj)
	if err != nil {
		panic("To complex128slice panic")
	}

	return res
}

func (o *Object) Complex128SliceWithErr() ([]complex128, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToComplex128Slice(o.obj)
}

func (o *Object) BoolSlice() []bool {
	if o.err != nil {
		panic("To boolslice panic")
	}

	res, err := utils.ToBoolSlice(o.obj)
	if err != nil {
		panic("To complex64slice panic")
	}

	return res
}

func (o *Object) BoolSliceWithErr() ([]bool, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToBoolSlice(o.obj)
}

func (o *Object) StringSlice() []string {
	if o.err != nil {
		panic("To stringslice panic")
	}

	res, err := utils.ToStringSlice(o.obj)
	if err != nil {
		panic("To stringslice panic")
	}

	return res
}

func (o *Object) StringSliceWithErr() ([]string, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToStringSlice(o.obj)
}

func (o *Object) ByteSlice() []byte {
	if o.err != nil {
		panic("To byteslice panic")
	}

	res, err := utils.ToByteSlice(o.obj)
	if err != nil {
		panic("To complex64slice panic")
	}

	return res
}

func (o *Object) ByteSliceWithErr() ([]byte, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToByteSlice(o.obj)
}

func (o *Object) RuneSlice() []rune {
	if o.err != nil {
		panic("To runeslice panic")
	}

	res, err := utils.ToRuneSlice(o.obj)
	if err != nil {
		panic("To runeslice panic")
	}

	return res
}

func (o *Object) RuneSliceWithErr() ([]rune, error) {
	if o.err != nil {
		return nil, o.err
	}

	return utils.ToRuneSlice(o.obj)
}
