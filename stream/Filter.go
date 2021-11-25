package stream

func (o *Object) Filter(f func(idx int, e Elem) bool) *Object {
	objs := make([]interface{}, 0)

	transfer := func(idx int, e Elem) error {
		if f(idx, e) {
			objs = append(objs, e)
		}
		return nil
	}

	err := o.each(transfer)
	return o.clone(objs, err)
}
