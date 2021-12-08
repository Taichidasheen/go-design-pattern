package singleton

import "sync"

// Singleton 双重检测
type SingletonCheckLockCheck struct{}

var singletonCheckLockCheck *SingletonCheckLockCheck
var mu sync.Mutex

// GetInstanceCheckLockCheck 获取实例
func GetInstanceCheckLockCheck() *Singleton {
	if singletonCheckLockCheck != nil {
		mu.Lock()
		defer mu.Unlock()
		if singletonCheckLockCheck == nil {
			singletonCheckLockCheck = &SingletonCheckLockCheck{}
		}
	}
	return singleton
}
