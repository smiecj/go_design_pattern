package builder_

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	director := new(Director)

	textBuilder := new(TextBuilder)
	director.Construct(textBuilder)
	fmt.Println(textBuilder.GetResult())

	htmlBuilder := NewHTMLBuilder()
	director.Construct(htmlBuilder)
	fmt.Println(htmlBuilder.GetResult())
}
