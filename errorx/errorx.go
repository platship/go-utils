package errorx

import "fmt"

func ErrorFuncBreak(fns ...func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func ErrorFuncPanic(fns ...func() error) {
	if err := ErrorFuncBreak(fns...); err != nil {
		panic(err)
	}
}

func ErrorReturn(errs ...error) error {
	fmt.Println(errs, errs)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
