package stream

func (o *Object) Chunk(size int) *Object {
	objs := []interface{}{}
	item := []interface{}{}
	transfer := func(idx int, obj Elem) error {
		if len(item) < size {
			item = append(item, obj)
		} else {
			objs = append(objs, item)
			item = []interface{}{obj}
		}
		return nil
	}

	if err := o.each(transfer); err != nil {
		return o.clone(nil, err)
	}

	if len(item) > 0 {
		objs = append(objs, item)
		item = []interface{}{}
	}

	return o.clone(objs, nil)
}
