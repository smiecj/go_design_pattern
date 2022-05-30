package factory_

import "sync"

var (
	carFactoryOnce     = sync.Once{}
	carFactoryInstance factory
)

type carFactory struct {
	carList []product
}

func (f *carFactory) create(name string) product {
	return newCar(name)
}

func (f *carFactory) register(p product) {
	f.carList = append(f.carList, p)
}

type carFactoryImpl struct {
	withRegisterFactoryI
}

func (cf *carFactoryImpl) Create(name string) product {
	car := cf.withRegisterFactoryI.create(name)
	cf.withRegisterFactoryI.register(car)
	return car
}

// 单例模式: factory
func NewCarFactory() factory {
	carFactoryOnce.Do(func() {
		carFactoryInstance = &carFactoryImpl{
			withRegisterFactoryI: new(carFactory),
		}
	})
	return carFactoryInstance
}
