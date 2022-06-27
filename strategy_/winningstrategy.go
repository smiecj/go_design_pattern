package strategy_

import "math/rand"

// winningStrategy: 上次赢方式优先策略
type winningStrategy struct {
	lastHand *hand
	lastWin  bool
	rand     *rand.Rand
}

func (strategy *winningStrategy) recordResult(win bool) {
	strategy.lastWin = win
}

func (strategy *winningStrategy) nextHand() *hand {
	if strategy.lastHand != nil && strategy.lastWin {
		return strategy.lastHand
	} else {
		nextHandIndex := rand.Intn(handCount)
		strategy.lastHand = getHandByIndex(nextHandIndex)
	}
	return strategy.lastHand
}

func NewWinningStrategy(seed int64) strategy {
	strategy := new(winningStrategy)
	strategy.rand = rand.New(rand.NewSource(seed))
	return strategy
}
