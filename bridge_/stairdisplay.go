package bridge_

// 阶梯形状打印功能扩展类
type stairDisplay struct {
	multipleDisplay
}

func (dp *stairDisplay) StairDisplay(maxLen, step int) {
	for index := 0; index < maxLen; index += step {
		dp.CountDisplay(index)
	}
}

func NewStairDisplay(impl displayImpl) *stairDisplay {
	return &stairDisplay{
		multipleDisplay: *NewMultipleDisplay(impl),
	}
}
