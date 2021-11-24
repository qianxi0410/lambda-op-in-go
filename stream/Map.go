package stream

func (o *Object) Map(f func(idx int, obj Elem) Elem) *Object {
	objs := make([]interface{}, 0)

	transfer := func(idx int, obj Elem) error {
		objs = append(objs, f(idx, obj))
		return nil
	}

	err := o.each(transfer)
	return o.clone(objs, err)
}
