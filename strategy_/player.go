package strategy_

import "fmt"

// Player: 猜拳玩家
type player struct {
	name     string
	strategy strategy
	total_   int
	win_     int
	lose_    int
	draw_    int
}

func (player *player) nextHand() *hand {
	return player.strategy.nextHand()
}

func (player *player) win() {
	player.total_++
	player.win_++
}

func (player *player) lose() {
	player.total_++
	player.lose_++
}

func (player *player) draw() {
	player.total_++
	player.draw_++
}

func (player *player) String() string {
	return fmt.Sprintf("Player: %s, total: %d, win: %d, draw: %d, lose: %d",
		player.name, player.total_, player.win_, player.draw_, player.lose_)
}

func NewPlayer(name string, strategy strategy) *player {
	return &player{
		name:     name,
		strategy: strategy,
	}
}
