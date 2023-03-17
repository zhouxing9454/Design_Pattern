/*
问题
存储着重要对象的全局变量，往往意味着“不安全”，因为你无法保证这个全局变量的值不会在项目的某个引用处被覆盖掉。
对数据的修改经常导致出乎意料的的结果和难以发现的bug。我在一处更新数据，却没有意识到软件中的另一处期望着完全不同的数据，于是一个功能就失效了，而且找出故障的原因也会非常困难。
一个较好的解决方案是：将这样的“可变数据”封装起来，写一个查询方法专门用来获取这些值。
单例模式则更进一步：除了要为“可变数据”提供一个全局访问方法，它还要保证获取到的只有同一个实例。也就是说，如果你打算用一个构造函数创建一个对象，单例模式将保证你得到的不是一个新的对象，而是之前创建过的对象，并且每次它所返回的都只有这同一个对象，也就是单例。这可以保护该对象实例不被篡改。


解决
单例模式需要一个全局构造函数，这个构造函数会返回一个私有的对象，无论何时调用，它总是返回相同的对象。

单例实例singleton被保存为一个私有的变量，以保证不被其他包的函数引用。
用构造方法GetInstance可以获得单例实例，函数中使用了sync包的once方法，以保证实例只会在首次调用时被初始化一次，之后再调用构造方法都只会返回同一个实例。
测试代码：
*/

package Singleton

import (
	"sync"
)

// 单例实例
type singleton struct {
	Value int
}

type Singleton interface {
	getValue() int
}

func (s singleton) getValue() int {
	return s.Value
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 构造方法，用于获取单例模式对象
func GetInstance(v int) Singleton {
	once.Do(func() {
		instance = &singleton{Value: v}
	})

	return instance
}
