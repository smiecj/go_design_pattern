package builder_

import (
	"bytes"
	"fmt"
)

type TextBuilder struct {
	buf bytes.Buffer
}

func (tb *TextBuilder) makeTitle(s string) {
	tb.buf.WriteString("=========================\n")
	tb.buf.WriteString(fmt.Sprintf("《%s》\n", s))
}

func (tb *TextBuilder) makeString(s string) {
	tb.buf.WriteString(fmt.Sprintf(" • %s\n", s))
}

func (tb *TextBuilder) makeItems(strArr []string) {
	for _, str := range strArr {
		tb.buf.WriteString(fmt.Sprintf("  • %s\n", str))
	}
}

func (tb *TextBuilder) close() error {
	tb.buf.WriteString("=========================\n")
	return nil
}

func (tb *TextBuilder) GetResult() string {
	return tb.buf.String()
}

func NewTextBuilder() *TextBuilder {
	return new(TextBuilder)
}
