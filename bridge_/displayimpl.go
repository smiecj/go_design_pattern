// package bridge_ 桥接模式
package bridge_

// displayImpl: 输出功能定义
type displayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}
