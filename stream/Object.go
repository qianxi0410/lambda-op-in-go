package stream

import "fmt"

// wrapper
type Object struct {
	// sting array slice
	obj interface{}
	err error
}

type Elem interface{}

func Init(obj interface{}) *Object {
	return &Object{obj: obj, err: nil}
}

var ErrBreak = fmt.Errorf("break range")
