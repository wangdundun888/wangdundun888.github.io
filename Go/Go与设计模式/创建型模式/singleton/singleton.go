package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton

var once sync.Once

func GetSingleton() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
