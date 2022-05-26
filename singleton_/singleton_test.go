package singleton_

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestSingleton(t *testing.T) {
	/* singleton1 := &singleton{
		name: "1",
	}
	singleton2 := &singleton{
		name: "2",
	} */
	singleton1 := GetSingleton()
	singleton2 := GetSingleton()
	pointer1 := (*float32)(unsafe.Pointer(singleton1)) // 输出效果和 %p 一致
	pointer2 := (*float32)(unsafe.Pointer(singleton2))

	// fmt.Println(reflect.DeepEqual(singleton1, singleton2))
	// fmt.Println(singleton1 == singleton2)
	// require.Same(t, singleton1, singleton2)
	fmt.Printf("pointer1: %v, pointer2: %v\n", pointer1, pointer2)

	// todo: 整理空 struct 地址都一样的原因，以及不同 type 的空 struct 是否地址也一样？
	// https://github.com/golang/go/issues/23440
	// https://github.com/golang/go/issues/12884
	// https://dave.cheney.net/2014/03/25/the-empty-struct
	// 思路: 尝试用 go build 工具分析源码

	require.Equal(t, singleton1, singleton2)
}
