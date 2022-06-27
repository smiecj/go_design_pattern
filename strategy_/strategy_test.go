package strategy_

import (
	"fmt"
	"testing"
	"time"
)

const (
	gameCount = 10000
)

func TestStrategy(t *testing.T) {
	// 模拟两个策略不同的选手，最后输出成功率
	winningStrategy := NewWinningStrategy(time.Now().UnixNano())
	probStrategy := NewProbStrategy(time.Now().UnixNano())
	player1 := NewPlayer("xiaoming", winningStrategy)
	player2 := NewPlayer("xiaohong", probStrategy)

	for index := 0; index < gameCount; index++ {
		hand1 := player1.nextHand()
		hand2 := player2.nextHand()
		if hand1.Stronger(hand2) {
			player1.win()
			player2.lose()
		} else if hand2.Stronger(hand1) {
			player2.win()
			player1.lose()
		} else {
			player1.draw()
			player2.draw()
		}
	}
	fmt.Printf("%s\n", player1)
	fmt.Printf("%s\n", player2)
}
