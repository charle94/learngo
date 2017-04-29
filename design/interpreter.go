package design

//给定一个语言定义它的文法的一种表示
//并定义一个解释器。这个解释器使用该表示来解释语言中的句子
//如果一种特定类型的问题发生的频率足够高，那么可能就值得将该问题的各个实例
//表述为一个简单语言中的句子，这样就可以构建一个解释器，该解释器通过解释这些句子来解决该问题
type PlayContext struct {
	text string
}
type Expression interface {
	Interpret(context PlayContext)
}
