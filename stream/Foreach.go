package stream

func (o *Object) Foreach(f func(idx, e Elem)) {
	transfer := func(idx int, e Elem) error {
		f(idx, e)
		return nil
	}

	o.each(transfer)
}
