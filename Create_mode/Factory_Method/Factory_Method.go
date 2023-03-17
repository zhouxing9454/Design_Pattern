/*
问题:
假设我们的业务需要一个支付渠道，我们开发了一个Pay方法，其可以用于支付
如果业务需求变更，需要我们提供多种支付方式，一种叫APay，一种叫BPay，这二种支付方式所需的参数不同，APay只需要订单号OrderId，BPay则需要订单号OrderId和Uid。此时如何修改？
我们为Pay接口实现了APay() 和BPay() 方法。虽然暂时实现了业务需求，但却使得结构体PayReq变得冗余了，APay() 并不需要Uid参数。如果之后再增加CPay、DPay、EPay，可想而知，代码会变得越来越难以维护。
随着后续业务迭代，将不得不编写出复杂的代码。

解决:
让我们想象一个工厂类，这个工厂类需要生产电线和开关等器具，我们可以为工厂类提供一个生产方法，当电线机器调用生产方法时，就产出电线，当开关机器调用生产方法时，就产出开关。
套用到我们的支付业务来，就是我们不再为接口提供APay方法、BPay方法，而只提供一个Pay方法，并将A支付方式和B支付方式的区别下放到子类。
我们用APay和BPay两个结构体重写了Pay() 方法，如果需要添加一种新的支付方式， 只需要重写新的Pay() 方法即可。
工厂方法的优点就在于避免了创建者和具体产品之间的紧密耦合，从而使得代码更容易维护。
*/

package Factory_Method

import "fmt"

type Pay interface {
	Pay(string) int
}

type PayReq struct {
	OrderId string
}

type APayReq struct {
	PayReq
}

type BPayReq struct {
	PayReq
	Uid int64
}

func (p *APayReq) Pay() string {
	fmt.Println(p.OrderId)
	return "APay支付成功"
}

func (p *BPayReq) Pay() string {
	fmt.Println(p.OrderId)
	fmt.Println(p.Uid)
	return "BPay支付成功"
}
