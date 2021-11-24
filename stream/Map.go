package stream

func (o *Object) Map(f func(idx int, e Elem) Elem) *Object {
	objs := make([]interface{}, 0)

	transfer := func(idx int, e Elem) error {
		objs = append(objs, f(idx, e))
		return nil
	}

	err := o.each(transfer)
	return o.clone(objs, err)
}
