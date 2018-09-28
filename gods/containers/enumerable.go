package containers

type EnumerableWithIndex interface {
	Each(func(index int, value interface{}))
	Any(func(index int, value interface{}) bool) bool
	All(func(index int, value interface{}) bool) bool
	Find(func(index int, value interface{}) bool) (int, interface{})
}
