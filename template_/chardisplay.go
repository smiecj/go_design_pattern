package template_

import "fmt"

type charDisplay struct {
	display
}

type charEchoer struct {
	r rune
}

func (ce *charEchoer) open() {
	fmt.Print(">>")
}

func (ce *charEchoer) print() {
	fmt.Printf("%c", ce.r)
}

func (ce *charEchoer) close() {
	fmt.Println("<<")
}

func NewCharDisplay(r rune) display {
	return &charDisplay{
		display: &print5TimesDisplay{
			echo: &charEchoer{r: r},
		},
	}
}
