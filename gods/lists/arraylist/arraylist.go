package arraylist

import (
	"learngo/gods/lists"
)

func assertListImplementation() {
	var _ lists.List = (*List)(nil)
}

type List struct {
	elements []interface{}
	size     int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

func (list *List) Contains(values ...interface{}) bool {
	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if element == searchValue {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Insert(index int, values ...interface{}) {
	if list.withinRange(index) {
		if index == list.size {
			list.Add(values)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

func (list *List) Values() []interface{} {
	return list.elements
}

func (list *List) Clear() {
	list.elements = []interface{}{}
	list.size = 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) withinRange(index int) bool {
	return 0 <= index && index < list.size
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n > currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}

	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}
