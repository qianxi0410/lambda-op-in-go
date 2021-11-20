package lambda

// wrapper
type Object struct {
	// sting array slice
	obj interface{}
	err error
}

func New(obj interface{}) *Object {
	return &Object{obj: obj, err: nil}
}
