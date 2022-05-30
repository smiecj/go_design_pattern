package factory_

// 和 template method 类似，将接口分成对外开放和不对外开放两部分，分别实现
type factory interface {
	Create(name string) product
}

type withRegisterFactoryI interface {
	create(name string) product
	register(product)
}

type withRegisterFactory struct {
	withRegisterFactoryI
}

func (f *withRegisterFactory) Create(name string) product {
	p := f.withRegisterFactoryI.create(name)
	f.withRegisterFactoryI.register(p)
	return p
}
