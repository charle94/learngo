package design

/**
 * vs simplefactory
 * simplefactory ：工厂类中包含逻辑判断根据客户端的条件决定具体实现哪个类
 * factory：定义一个用于创建对象的接口，让子类来决定实例化哪一个类。工厂的方法使类的实例化延迟到子类当中
 */

type LeiFeng interface {
	Sweep()
	Wash()
}
type UnderGraduate struct {
}

func (u UnderGraduate) Sweep() {

}
func (u UnderGraduate) Wash() {

}

type Volunteer struct {
}

func (v Volunteer) Sweep() {

}
func (v Volunteer) Wash() {

}

/*type SimpleFactory struct {
}

func (s SimpleFactory) GetLei(who string) LeiFeng {
	switch who {
	case "graduate":
		return UnderGraduate{}
	case "volunteer":
		return Volunteer{}
	default:
		return UnderGraduate{}
	}
}*/
//工厂模式
//将实例化移动到子类当中
//更能好的方法应该是使用反射而不是使用case条件
//工厂模式就是根据条件先获取到工厂再用工厂实例化相应的类
type LeiFactory interface {
	CreateLei()
}
type UnderGraduateFactory struct {
}

func (u UnderGraduateFactory) CreateLei() LeiFeng {
	return UnderGraduate{}
}

type VolunteerFactory struct {
}

func (v VolunteerFactory) CreateLei() LeiFeng {
	return Volunteer{}
}
