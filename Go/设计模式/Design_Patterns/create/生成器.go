package main

import "fmt"

//生成接口定义
type builder interface {
	produce()   //生产
	check()     //检测
	deepCheck() //深度检测
}

//工厂A
//type factoryA struct{}

func (o *factoryA) produce()   { fmt.Println("factoryA::produce") }
func (o *factoryA) check()     { fmt.Println("factoryA::check") }
func (o *factoryA) deepCheck() { fmt.Println("nothing") }

//工厂B
//type factoryB struct{}

func (o *factoryB) produce()   { fmt.Println("factoryB::produce") }
func (o *factoryB) check()     { fmt.Println("factoryB::check") }
func (o *factoryB) deepCheck() { fmt.Println("factoryB::deepCheck") }

//导向器（主管 导演）
type director struct {
	builders builder
}

func newDirector(b builder) *director {
	return &director{b}
}

//build 普通生产:生产和检测，没有深度检测
func (o *director) build() {
	o.builders.produce()
	o.builders.check()
}

func (o *director) buildFroVip() {
	o.builders.produce()
	o.builders.check()
	o.builders.deepCheck()
}

//创建生成器
func findBuilder(money int) builder {
	if money > 10000 {
		return &factoryA{}
	} else {
		return &factoryB{}
	}
}

func main() {
	normal := findBuilder(2000)
	normalDirector := newDirector(normal)
	normalDirector.build()
	normalDirector.buildFroVip()

	vip := findBuilder(20000)
	vipDirector := newDirector(vip)
	vipDirector.build()
	vipDirector.buildFroVip()
}
