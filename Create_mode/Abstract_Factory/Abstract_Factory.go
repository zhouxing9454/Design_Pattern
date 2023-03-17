/*
问题
抽象工厂模式基于工厂方法模式。两者的区别在于：工厂方法模式是创建出一种产品，而抽象工厂模式是创建出一类产品。这二种都属于工厂模式，在设计上是相似的。
假设，有一个存储工厂，提供redis和mysql两种存储数据的方式。如果使用工厂方法模式，我们就需要一个存储工厂，并提供SaveRedis方法和SaveMysql方法。
如果此时业务还需要分成存储散文和古诗两种载体，这两种载体都可以进行redis和mysql存储。就可以使用抽象工厂模式，我们需要一个存储工厂作为父工厂，散文工厂和古诗工厂作为子工厂，并提供SaveRedis方法和SaveMysql方法。

解决
以上文的存储工厂业务为例，用抽象工厂模式的思路来设计代码，就像下面这样：
我们定义了存储工厂，也就是SaveArticle接口，并实现了CreateProse方法和CreateAncientPoetry方法，这2个方法分别用于创建散文工厂和古诗工厂。
然后我们又分别为散文工厂和古诗工厂实现了SaveProse方法和SaveAncientPoetry方法，并用Redis结构体和Mysql结构体分别重写了2种存储方法。
*/

package Abstract_Factory

import "fmt"

// SaveArticle 抽象模式工厂接口
type SaveArticle interface {
	CreateProse() Prose
	CreateAncientPoetry() AncientPoetry
}

type SaveRedis struct{}

func (*SaveRedis) CreateProse() Prose {
	return &RedisProse{}
}

func (*SaveRedis) CreateAncientPoetry() AncientPoetry {
	return &RedisProse{}
}

type SaveMysql struct{}

func (*SaveMysql) CreateProse() Prose {
	return &MysqlProse{}
}

func (*SaveMysql) CreateAncientPoetry() AncientPoetry {
	return &MysqlProse{}
}

// Prose 散文
type Prose interface {
	SaveProse()
}

// AncientPoetry 古诗
type AncientPoetry interface {
	SaveAncientPoetry()
}

type RedisProse struct{}

func (*RedisProse) SaveProse() {
	fmt.Println("Redis Save Prose")
}

func (*RedisProse) SaveAncientPoetry() {
	fmt.Println("Redis Save Ancient Poetry")
}

type MysqlProse struct{}

func (*MysqlProse) SaveProse() {
	fmt.Println("Mysql Save Prose")
}

func (*MysqlProse) SaveAncientPoetry() {
	fmt.Println("Mysql Save Ancient Poetry")
}
