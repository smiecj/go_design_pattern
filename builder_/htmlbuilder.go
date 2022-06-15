package builder_

import (
	"fmt"
	"os"
)

type HTMLBuilder struct {
	file *os.File
}

func (hb *HTMLBuilder) makeTitle(s string) {
	var err error
	hb.file, err = os.OpenFile(s+".html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if nil != err {
		fmt.Println(err.Error())
	}
	hb.file.WriteString(fmt.Sprintf("<html><head><title>%s</title></head><body>", s))
}

func (hb *HTMLBuilder) makeString(s string) {
	hb.file.WriteString(fmt.Sprintf("<p>%s</p>", s))
}

func (hb *HTMLBuilder) makeItems(strArr []string) {
	hb.file.WriteString("<ul>")
	for _, str := range strArr {
		hb.file.WriteString(fmt.Sprintf("<li>%s</li>", str))
	}
	hb.file.WriteString("</ul>")
}

func (hb *HTMLBuilder) close() error {
	hb.file.WriteString("</body></html>")
	return hb.file.Close()
}

func (hb *HTMLBuilder) GetResult() string {
	return hb.file.Name()
}

func NewHTMLBuilder() *HTMLBuilder {
	return new(HTMLBuilder)
}
