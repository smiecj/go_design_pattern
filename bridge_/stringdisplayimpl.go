package bridge_

import "fmt"

// stringDisplayImpl: 字符串打印功能实现类
type stringDisplayImpl struct {
	str string
}

func (impl *stringDisplayImpl) rawOpen() {
	impl.println()
}

func (impl *stringDisplayImpl) rawPrint() {
	fmt.Print("|")
	fmt.Print(impl.str)
	fmt.Println(("|"))
}

func (impl *stringDisplayImpl) rawClose() {
	impl.println()
}

func (impl *stringDisplayImpl) println() {
	fmt.Print("+")
	for index := 0; index < len(impl.str); index++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func NewStringDisplayImpl(str string) *stringDisplayImpl {
	return &stringDisplayImpl{
		str: str,
	}
}
