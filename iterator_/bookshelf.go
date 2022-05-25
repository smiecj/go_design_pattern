package iterator_

type bookShelf struct {
	books []book
	size  int
}

func newBookShelf(size int) *bookShelf {
	shelf := bookShelf{
		size:  0,
		books: make([]book, size),
	}
	return &shelf
}

func (shelf *bookShelf) putBook(book book) bool {
	if shelf.size >= len(shelf.books) {
		return false
	}
	shelf.books[shelf.size] = book
	shelf.size++
	return true
}

func (shelf *bookShelf) getBookAt(index int) (book, bool) {
	if index < 0 || index >= shelf.size {
		return book{}, false
	}
	return shelf.books[index], true
}

func (shelf *bookShelf) iterator() iterator {
	return &bookShelfIterator{shelf: shelf}
}
