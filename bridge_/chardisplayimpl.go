package bridge_

import "fmt"

// charDisplayImpl: 字符打印功能实现类
type charDisplayImpl struct {
	toPrintC rune
	start    rune
	end      rune
}

func (impl *charDisplayImpl) rawOpen() {
	fmt.Printf("%c", impl.start)
}

func (impl *charDisplayImpl) rawPrint() {
	fmt.Printf("%c", impl.toPrintC)
}

func (impl *charDisplayImpl) rawClose() {
	fmt.Printf("%c\n", impl.end)
}

func NewCharDisplayImpl(start, toPrintC, end rune) *charDisplayImpl {
	return &charDisplayImpl{
		start:    start,
		toPrintC: toPrintC,
		end:      end,
	}
}
