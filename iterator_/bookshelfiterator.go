package iterator_

type bookShelfIterator struct {
	shelf *bookShelf
	index int
}

func (it *bookShelfIterator) hasNext() bool {
	if it.index >= it.shelf.size {
		return false
	}
	return true
}

func (it *bookShelfIterator) next() interface{} {
	it.index = it.index + 1
	book, _ := it.shelf.getBookAt(it.index - 1)
	return book
}
