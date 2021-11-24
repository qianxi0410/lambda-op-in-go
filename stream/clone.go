package lambda

func (o *Object) clone(obj interface{}, err error) *Object {
	if err != nil {
		return &Object{
			err: err,
			obj: o.obj,
		}
	}
	if o.err != nil {
		return &Object{
			err: o.err,
			obj: o.obj,
		}
	}

	return &Object{
		err: nil,
		obj: obj,
	}
}
