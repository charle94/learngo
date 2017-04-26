package design

//迭代器模式
//提供一种方法顺序访问一个聚合对象
//中各个元素，而又不暴露该对象的内部表示

//需要访问一个聚集对象，而不管这些对象是什么
//都需要遍历的时候，应考虑使用迭代器模式
type Iterator interface {
	First() interface{}
	Last() interface{}
	Pre()
	Next()
	IsDone() bool
	CurrentItem() interface{}
}
type Aggreate interface {
	CreateIterator()
}
type ConcreteIterator struct {
	agg     ConcreteAggreate
	current int
}

//
func (c ConcreteIterator) First() interface{} {
	return c.agg.items[0]
}

//获取最后一个同时将下标后置到最后
func (c *ConcreteIterator) Last() interface{} {
	c.current = c.agg.Count() - 1
	return c.agg.items[c.agg.Count()-1]
}

//前一个下标
func (c *ConcreteIterator) Pre() {
	if c.current >= 0 {
		c.current--
	}
}

//下一个下标
func (c *ConcreteIterator) Next() {
	if c.current < c.agg.Count() {
		c.current++
	}
}

//实现了前后遍历结束条件
func (c ConcreteIterator) IsDone() bool {
	if c.current >= c.agg.Count() || c.current < 0 {
		return true
	} else {
		return false
	}
}

//只负责返回当前位置的值
func (c ConcreteIterator) CurrentItem() interface{} {
	return c.agg.items[c.current]
}

type ConcreteAggreate struct {
	items []string
}

func (c *ConcreteAggreate) CreateIterator(a []string) {
	c.items = a
}
func (c ConcreteAggreate) Count() int {
	return len(c.items)
}
