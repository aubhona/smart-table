package utils

type Optional[T any] struct {
	value T
	valid bool
}

func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{value: v, valid: true}
}

func EmptyOptional[T any]() Optional[T] {
	return Optional[T]{valid: false}
}

func (o Optional[T]) Value() T {
	if !o.valid {
		panic("bad optional access")
	}

	return o.value
}

func (o Optional[T]) HasValue() bool {
	return o.valid
}

func (o Optional[T]) ToPointer() *T {
	if o.valid {
		return &o.value
	}

	return nil
}

func OptionalFromPointer[T any](ptr *T) Optional[T] {
	if ptr == nil {
		return EmptyOptional[T]()
	}

	return NewOptional(*ptr)
}
