package stream

func (o *Object) Limit(limit int) *Object {
	objs := make([]interface{}, 0, limit)

	transfer := func(idx int, e Elem) error {
		if idx < limit {
			objs = append(objs, e)
		}
		return nil
	}

	err := o.each(transfer)
	return o.clone(objs, err)
}
