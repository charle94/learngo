package design

//为子系统中的一组接口提供一个一致的界面，
//此模式定义了一个高层接口，这个接口使得这
//个子系统更加容易使用
//知道子系统类负责处理哪些请求将客户的请求代理给适当的子系统对象
//子系统类集合，实现子系统功能，处理facade对象指派的任务子类中没有facade
//的任何信息
//业务逻辑层和表示层之间建立外观facade
//增加外观模式减少依赖
//维护遗留的大型系统建立一个facade类完成新代码与旧系统的沟通
type SubSystenOne struct {
}

func (s SubSystenOne) MethodOne() {

}

type SubSystenTwo struct {
}

func (s SubSystenTwo) MethodTwo() {

}

type SubSystenThree struct {
}

func (s SubSystenThree) MethodThree() {

}

type Facade struct {
	One   SubSystenOne
	Two   SubSystenTwo
	Three SubSystenThree
}

func (f Facade) MethodA() {
	f.One.MethodOne()
	f.Two.MethodTwo()
}
func (f Facade) MethodB() {
	f.Three.MethodThree()
}
