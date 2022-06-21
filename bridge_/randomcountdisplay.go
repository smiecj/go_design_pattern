package bridge_

import (
	"math"
	"math/rand"
	"time"
)

// 多次打印功能扩展类
type randomDisplay struct {
	multipleDisplay
}

func (dp *randomDisplay) RandomDisplay(times int) {
	rand.Seed(time.Now().UnixNano())
	randomTimes := int(math.Floor(float64(rand.Float32() * float32((times + 1)))))
	dp.CountDisplay(randomTimes)
}

func NewRandomDisplay(impl displayImpl) *randomDisplay {
	return &randomDisplay{
		multipleDisplay: *NewMultipleDisplay(impl),
	}
}
