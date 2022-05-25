// package iterator 迭代器实现
package iterator_

// 可迭代对象（一般用来实现集合）
type iterable interface {
	iterator() iterator
}
