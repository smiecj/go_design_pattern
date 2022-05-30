// package template_ 模板方法模式
package template_

// go 中没有 abstract class，必须通过 “多继承” 实现
// 接口: echo, display (将接口拆成 需要对外开放、不需要对外开放 两部分)
// display 实现: print5TimesDisplay
// echo 实现: charEchoer
// 实际实现类要引用: charEchoer、print5TimesDisplay
type echo interface {
	open()
	print()
	close()
}

type display interface {
	echo
	Display()
}

type print5TimesDisplay struct {
	echo
}

func (d *print5TimesDisplay) Display() {
	d.open()
	for index := 0; index < 5; index++ {
		d.print()
	}
	d.close()
}
