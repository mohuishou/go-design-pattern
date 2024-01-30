package singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	// once 内的方法只会执行一次，所以不需要再次判断
	once.Do(func() {
		lazySingleton = &Singleton{a: "test"}
	})

	return lazySingleton
}
