/*
问题
假设业务需要按步骤创建一系列复杂的对象，实现这些步骤的代码加在一起非常繁复，我们可以将这些代码放进一个包含了众多参数的构造函数中，但这个构造函数看起来将会非常杂乱无章，且难以维护。
假设业务需要建造一个房子对象，需要先打地基、建墙、建屋顶、建花园、放置家具……。我们需要非常多的步骤，并且这些步骤之间是有联系的，即使将各个步骤从一个大的构造函数抽出到其他小函数中，整个程序的层次结构看起来依然很复杂。
如何解决呢？像这种复杂的有许多步骤的构造函数，就可以用建造者模式来设计。
建造者模式的用处就在于能够分步骤创建复杂对象。

解决
在建造者模式中，我们需要清晰的定义每个步骤的代码，然后在一个构造函数中操作这些步骤，我们需要一个主管类，用这个主管类来管理各步骤。这样我们就只需要将所需参数传给一个构造函数，构造函数再将参数传递给对应的主管类，最后由主管类完成后续所有建造任务。

*/

package Builder

import "fmt"

// Builder 建造者接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Director 管理类
type Director struct {
	builder Builder
}

// NewDirector 构造函数
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct 建造
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Builder2 struct{}

func (b *Builder2) Part1() {
	fmt.Println("part1")
}

func (b *Builder2) Part2() {
	fmt.Println("part2")
}

func (b *Builder2) Part3() {
	fmt.Println("part3")
}
