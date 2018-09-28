package singlylinkedlist

import (
	"learngo/gods/lists"
	"learngo/gods/utils"
)

func assertImplementation() {
	var _ lists.List = (*List)(nil)
}

type element struct {
	value interface{}
	next  *element
}

type List struct {
	first *element
	last  *element
	size  int
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	element := list.first
	for i := 0; i != index; i, element = i+1, element.next {
	}
	return element, true
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var preEle *element
	element := list.first
	for i := 0; i != index; i, element = i+1, element.next {
		preEle = element
	}
	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = preEle
	}
	if preEle != nil {
		preEle.next = element.next
	}
	element = nil
	list.size--
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values)
		}
	}

	list.size += len(values)

	var preEle *element
	foundElement := list.first
	for i := 0; i != index; i, foundElement = i+1, foundElement.next {
		preEle = foundElement
	}

	if foundElement == list.first {
		for _, value := range values {
			newElement := &element{value: value, next: list.first}
			list.first = newElement
		}
	} else {
		for _, value := range values {
			newElement := &element{value: value, next: foundElement}
			foundElement = newElement
		}
		preEle.next = foundElement
	}
}

func (list *List) Set(index int, value interface{}) {

}

func (list *List) Sort(comparator utils.Comparator) {

}

func (list *List) Contains(values ...interface{}) bool {
	return false
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (lists *List) Size() int {
	return lists.size
}

func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *List) withinRange(index int) bool {
	return 0 <= index && index < list.size
}
