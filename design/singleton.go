package design

import (
	"sync"
)

//保证一个类仅有一个实例
//提供一个全局访问点
type singleton struct {
	Name string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{Name: "john"}
	})
	return instance
}
