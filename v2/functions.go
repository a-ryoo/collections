package v2

import "reflect"

type Predicate[E any] func(e E) bool

type Action[E any] func(e E)

type MultiAction[E any] func(x E, z E)

type Producer[E any] func(e E) E

type Collector[E any] func(x E, z E) []E

type MapperFunc[S, T any] func(h S) T

func Equals[E any](g E) Predicate[E] {
	return func(h E) bool {
		return reflect.DeepEqual(h, g)
	}
}

func NotEquals[E any](g E) Predicate[E] {
	return func(h E) bool {
		return !reflect.DeepEqual(h, g)
	}
}
