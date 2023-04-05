package command

import "fmt"

// 请求者

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// 具体命令接口

type command interface {
	execute()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// 接收者

type device interface {
	on()
	off()
}

type tv struct{}

func (t *tv) on() {
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	fmt.Println("Turning tv off")
}

type airConditioner struct{}

func (t *airConditioner) on() {
	fmt.Println("Turning air conditioner on")
}

func (t *airConditioner) off() {
	fmt.Println("Turning air conditioner off")
}
