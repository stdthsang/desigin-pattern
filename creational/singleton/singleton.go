package singleton

import "sync"

var Instance *singleton

type singleton struct {
}

func init() {
	Instance = &singleton{}
}

func GetInstance1() *singleton {
	if Instance == nil {
		Instance = &singleton{}
	}
	return Instance
}

var lock sync.RWMutex

func GetInstance2() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if Instance == nil {
		Instance = &singleton{}
	}
	return Instance
}

func GetInstance3() *singleton {
	if Instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if Instance == nil {
			Instance = &singleton{}
		}
	}
	return Instance
}

var once sync.Once

func GetInstance4() *singleton {
	once.Do(func() {
		Instance = &singleton{}
	})
	return Instance
}
