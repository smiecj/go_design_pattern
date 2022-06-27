package strategy_

import "math/rand"

// randomStrategy: 完全随机策略
type randomStrategy struct {
	rand *rand.Rand
}

func (strategy *randomStrategy) recordResult(win bool) {
}

func (strategy *randomStrategy) nextHand() *hand {
	return getHandByIndex(strategy.rand.Intn(handCount))
}

func NewRandomStrategy(seed int64) strategy {
	return &randomStrategy{
		rand: rand.New(rand.NewSource(seed)),
	}
}
