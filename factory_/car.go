package factory_

import "fmt"

type car struct {
	name string
}

func (c *car) use() {
	fmt.Println("用车: " + c.name)
}

func newCar(name string) *car {
	return &car{
		name: name,
	}
}
