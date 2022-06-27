package strategy_

import (
	"math/rand"
)

// probStrategy: 结合连续两次结果统计策略
type probStrategy struct {
	rand                  *rand.Rand
	winCountRecordArr     [][]int
	lastHand, currentHand *hand
}

func (strategy *probStrategy) recordResult(win bool) {
	if win {
		strategy.winCountRecordArr[strategy.lastHand.index][strategy.currentHand.index]++
	}
}

func (strategy *probStrategy) nextHand() *hand {
	sum := 0
	for index := 0; index < handCount; index++ {
		sum += strategy.winCountRecordArr[strategy.currentHand.index][index]
	}
	randInt := strategy.rand.Intn(sum) + 1
	total := 0
	for index := 0; index < handCount; index++ {
		total += strategy.winCountRecordArr[strategy.currentHand.index][index]
		if total >= randInt {
			strategy.lastHand = strategy.currentHand
			strategy.currentHand = getHandByIndex(index)
			return strategy.currentHand
		}
	}
	return nil
}

func NewProbStrategy(seed int64) strategy {
	strategy := new(probStrategy)
	strategy.rand = rand.New(rand.NewSource(seed))
	strategy.winCountRecordArr = make([][]int, handCount)
	for index := 0; index < handCount; index++ {
		strategy.winCountRecordArr[index] = make([]int, handCount)
		for y := 0; y < handCount; y++ {
			strategy.winCountRecordArr[index][y] = 1
		}
	}
	// 初始化: 设置出拳方式都为第一种出拳
	strategy.lastHand, strategy.currentHand = getHandByIndex(0), getHandByIndex(0)

	return strategy
}
