// package strategy_ 策略模式
package strategy_

// strategy: 策略接口
type strategy interface {
	recordResult(bool)
	nextHand() *hand
}
