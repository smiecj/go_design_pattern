// package builder_ 工厂模式
package builder_

// builder 行为定义
type builder interface {
	makeTitle(string)
	makeString(string)
	makeItems([]string)
	close() error
}
