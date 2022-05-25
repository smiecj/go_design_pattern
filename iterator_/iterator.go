package iterator_

// 迭代器
type iterator interface {
	hasNext() bool
	next() interface{}
}
