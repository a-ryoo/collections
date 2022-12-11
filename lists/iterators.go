package collections

type Iterator[E any] interface {
	First()
	Next()
	IsDone() bool
	CurrentItem() E
	ForEach(action Action[E])
}

type ListIterator[E any] struct {
	index     int
	container []E
}

func (i *ListIterator[E]) First() {
	i.index = 0
}

func (i *ListIterator[E]) Next() {
	i.index++
}

func (i *ListIterator[E]) IsDone() bool {
	return i.index >= len(i.container)
}

func (i *ListIterator[E]) CurrentItem() E {
	return i.container[i.index]
}

func (i *ListIterator[E]) ForEach(action Action[E]) {
	for i.First(); i.IsDone(); i.Next() {
		action(i.CurrentItem())
	}
}
