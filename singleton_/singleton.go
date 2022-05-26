// package singleton_ 单例模式实现
package singleton_

import "sync"

var (
	// 单例实例，私有
	singletonInstance *singleton
	// 通过 once 实现单例效果
	once sync.Once
)

// 单例类，私有
// 这里注意: 空 struct 即使 new 之后也是同一个对象
type singleton struct {
	name string
}

func GetSingleton() *singleton {
	once.Do(func() {
		singletonInstance = &singleton{
			name: "test",
		}
	})
	return singletonInstance
}
