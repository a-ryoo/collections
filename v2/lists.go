package v2

type List[E any] interface {
	Add(e ...E) List[E]
	AddAll(col List[E]) List[E]
	Clear()
	Contains(e ...E) bool
	DeepEquals(col List[E]) bool
	Equals(col List[E]) bool
	Filter(pred Predicate[E])
	ForEach(action Action[E])
	GetHash() string
	IsEmpty() bool
	Iterator() Iterator[E]
	Collect() []E
	Remove(e ...E)
	RemoveIf(pred Predicate[E])
}

type MultiList[S, T any] interface {
	Add(e ...S) MultiList[S, T]
	AddAll(lst List[S]) MultiList[S, T]
	Clear()
	Collect() []S
	Map(mapperFunc MapperFunc[S, T]) List[T]
}
