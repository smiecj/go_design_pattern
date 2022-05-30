package factory_

import "testing"

func TestCarFactory(t *testing.T) {
	factory := NewCarFactory()
	c1 := factory.Create("奔驰")
	c2 := factory.Create("宝马")
	c3 := factory.Create("奥迪")
	c1.use()
	c2.use()
	c3.use()
}
