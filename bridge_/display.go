package bridge_

// display: 打印功能基础扩展类
type display struct {
	impl displayImpl
}

func (dp *display) open() {
	dp.impl.rawOpen()
}

func (dp *display) print() {
	dp.impl.rawPrint()
}

func (dp *display) close() {
	dp.impl.rawClose()
}

func (dp *display) Display() {
	dp.open()
	dp.print()
	dp.close()
}

func NewDisplay(impl displayImpl) *display {
	return &display{
		impl: impl,
	}
}
