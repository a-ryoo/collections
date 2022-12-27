package col

import (
	"crypto/md5"
	"encoding/base64"
	"reflect"
	"unsafe"
)

type ArrayList[E any] struct {
	container []E
}

func NewArrayList[E any](e ...E) List[E] {
	var container []E
	if len(e) < 1 {
		container = make([]E, 0)
	} else {
		container = e
	}

	return &ArrayList[E]{
		container: container,
	}
}

func (a *ArrayList[E]) Clear() {
	a.container = []E{}
}

func (a *ArrayList[E]) Add(e ...E) List[E] {
	for _, el := range e {
		a.container = append(a.container, el)
	}
	return a
}

func (a *ArrayList[E]) AddAll(col List[E]) List[E] {
	host := a.container
	hostLen := len(host)

	guest := col.Collect()
	guestLen := len(guest)
	totalLen := hostLen + guestLen

	placeholder := make([]E, totalLen)

	for i, e := range host {
		placeholder[i] = e
	}

	guestIndex := 0

	for i := hostLen; i < totalLen; i++ {
		placeholder[i] = guest[guestIndex]
		guestIndex++
	}

	a.container = placeholder
	return a
}

func (a *ArrayList[E]) Remove(e ...E) {
	var temp []E
	for _, h := range a.container {
		for _, g := range e {
			if !reflect.DeepEqual(h, g) {
				temp = append(temp, h)
			}
		}
	}
	a.container = temp
}

func (a *ArrayList[E]) RemoveIf(pred Predicate[E]) {
	var result []E
	for _, e := range a.container {
		if !pred(e) {
			result = append(result, e)
		}
	}
	a.container = result
}
func (a *ArrayList[E]) Filter(pred Predicate[E]) List[E] {
	var result []E
	for _, e := range a.container {
		if pred(e) {
			result = append(result, e)
		}
	}
	a.container = result

	return a
}

func (a *ArrayList[E]) Contains(e ...E) bool {
	for _, h := range a.container {
		for _, g := range e {
			if reflect.DeepEqual(h, g) {
				return true
			}
		}
	}
	return false
}

func (a *ArrayList[E]) Equals(e List[E]) (result bool) {
	if len(a.container) != len(e.Collect()) {
		return result
	}

	for _, h := range a.container {
		for _, g := range e.Collect() {
			if !reflect.DeepEqual(h, g) {
				return
			}
		}
	}
	return false
}

func (a *ArrayList[E]) DeepEquals(col List[E]) bool {
	return a.GetHash() == col.GetHash()
}

func (a *ArrayList[E]) ForEach(action Action[E]) {
	for _, i := range a.container {
		action(i)
	}
}

func (a *ArrayList[E]) IsEmpty() bool {
	return len(a.container) == 0
}

func (a *ArrayList[E]) Collect() []E {
	return a.container
}

func (a *ArrayList[E]) Iterator() Iterator[E] {
	return &ListIterator[E]{container: a.container}
}

func (a *ArrayList[E]) GetHash() string {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Ptr {
		v = v.Addr()
	}

	size := unsafe.Sizeof(v.Interface())
	b := (*[1 << 10]uint8)(unsafe.Pointer(v.Pointer()))[:size:size]

	h := md5.New()
	return base64.StdEncoding.EncodeToString(h.Sum(b))
}

type MultiArrayList[S, T any] struct {
	container []S
}

func NewMultiArrayList[S, T any](src ...S) MultiList[S, T] {
	var container []S
	if len(src) < 1 {
		container = make([]S, 0)
	} else {
		container = src
	}

	return &MultiArrayList[S, T]{
		container: container,
	}
}

func (a *MultiArrayList[S, T]) Add(e ...S) MultiList[S, T] {
	for _, el := range e {
		a.container = append(a.container, el)
	}
	return a
}

func (a *MultiArrayList[S, T]) AddAll(col List[S]) MultiList[S, T] {
	host := a.container
	hostLen := len(host)

	guest := col.Collect()
	guestLen := len(guest)
	totalLen := hostLen + guestLen

	placeholder := make([]S, totalLen)

	for i, e := range host {
		placeholder[i] = e
	}

	guestIndex := 0

	for i := hostLen; i < totalLen; i++ {
		placeholder[i] = guest[guestIndex]
		guestIndex++
	}

	a.container = placeholder
	return a
}

func (a *MultiArrayList[S, T]) Clear() {
	a.container = []S{}
}

func (a *MultiArrayList[S, T]) Collect() []S {
	return a.container
}

func (a *MultiArrayList[S, T]) Map(mapperFunc MapperFunc[S, T]) List[T] {
	result := NewArrayList[T]()
	for _, e := range a.container {
		result.Add(mapperFunc(e))
	}

	return result
}
