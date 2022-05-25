package iterator_

import (
	"log"
	"testing"
)

func TestBookShelfIterator(t *testing.T) {
	shelf := newBookShelf(100)
	shelf.putBook(book{name: "语文"})
	shelf.putBook(book{name: "数学"})
	shelf.putBook(book{name: "英语"})
	it := shelf.iterator()
	for it.hasNext() {
		book := it.next().(book)
		log.Println(book.name)
	}
}
