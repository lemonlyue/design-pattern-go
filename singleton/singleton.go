package singleton

import "sync"

type singleton struct {

}

var (
	instance *singleton
	once sync.Once
)

// 获取单例模式对象
func GetInstance() singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return *instance
}
