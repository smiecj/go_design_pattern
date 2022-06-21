package bridge_

// 多次打印功能扩展类
type multipleDisplay struct {
	display
}

func (dp *multipleDisplay) CountDisplay(times int) {
	dp.open()
	for index := 0; index < times; index++ {
		dp.print()
	}
	dp.close()
}

func NewMultipleDisplay(impl displayImpl) *multipleDisplay {
	return &multipleDisplay{
		display: *NewDisplay(impl),
	}
}
