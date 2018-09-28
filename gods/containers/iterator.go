package containers

type IteratorWithIndex interface {
	Next() bool
	Value() interface{}
	Index() int
	Begin()
	First() bool
}
