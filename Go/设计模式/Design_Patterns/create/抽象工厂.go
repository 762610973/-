package main

import "fmt"

//抽象工厂
type factory interface {
	workShop() workShop
	eatery() eatery
}

//车间标准
type workShop interface {
	work()
}

//食堂标准
type eatery interface {
	eating()
}
type factoryAInfo struct{}

func (o *factoryAInfo) workShop() workShop { return &factoryA{} }
func (o *factoryAInfo) eatery() eatery     { return &factoryA{} }

// factoryA 工厂A的车间和食堂
type factoryA struct{}

func (o *factoryA) work()   { fmt.Println("factoryA :: work") }
func (o *factoryA) eating() { fmt.Println("factoryA :: eating") }

//factoryB 工厂B
type factoryBInfo struct{}

func (o *factoryBInfo) workShop() workShop { return &factoryB{} }
func (o *factoryBInfo) eatery() eatery     { return &factoryB{} }

//factoryB 工厂B的车间和食堂
type factoryB struct{}

func (o *factoryB) work()   { fmt.Println("factoryB :: work") }
func (o *factoryB) eating() { fmt.Println("factoryB :: eating") }

//提供不同的参数，创建不同的工厂
func findWork(money int) factory {
	if money > 1000 {
		return &factoryAInfo{}
	} else {
		return &factoryBInfo{}
	}
}

func main() {
	//findWork 返回factoryA/BInfo 这个能调用workShop 和 eatery ,然后workShop()和eatery()函数返回的是factoryA/B
	//factoryA/B能调用work和eating函数
	wk1 := findWork(20)
	wk1.workShop().work()
	wk1.eatery().eating()

	wk2 := findWork(2000)
	wk2.workShop().work()
	wk2.eatery().eating()
}
