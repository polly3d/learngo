package containers

import (
	"learngo/gods/utils"
)

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}

func GetSortValues(container Container, comparator utils.Comparator) []interface{} {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	utils.Sort(values, comparator)
	return values
}
