package lambda

import (
	"errors"
	"qianxi/lambda-go/lambda/utils"
)

func (o *Object) each(f func(idx int, itm Elem) error) error {
	if o.err != nil {
		return o.err
	}

	arr, err := utils.ToInterfaceList(o.obj)
	if err != nil {
		return err
	}

	for i, v := range arr {
		if err := f(i, v); err != nil {
			if errors.Is(err, ErrBreak) {
				return nil
			}
			return err
		}
	}
	return nil
}
