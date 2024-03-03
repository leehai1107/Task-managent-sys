package wapper

import "fmt"

type WrappedError struct {
	Context string
	Err     string
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err.Error(),
	}
}
