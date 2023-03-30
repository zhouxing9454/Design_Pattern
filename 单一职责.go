package main

import "fmt"

type ClothesShop struct{}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

type ClothesWork struct{}

func (cs *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}

func main() {
	cs := ClothesShop{}
	cs.Style()

	cw := ClothesWork{}
	cw.Style()
}
