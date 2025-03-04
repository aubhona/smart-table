package utils

import "errors"

type SharedRef[T any] struct {
	value *T
}

func NewSharedRef[T any](v *T) (SharedRef[T], error) {
	if v == nil {
		return SharedRef[T]{}, errors.New("SharedRef cannot be nil")
	}

	return SharedRef[T]{value: v}, nil
}

func (r SharedRef[T]) Get() *T {
	return r.value
}

func (r SharedRef[T]) Value() T {
	return *r.value
}
