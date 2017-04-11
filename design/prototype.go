package design

//原型实例指定创建对象的种类，
//并通过拷贝创建这些对象
//从一个对象创建另一个可定制的对象，而且不需要
//知道创建的细节

//copy 是一个复制函数通过序列化与反序列化进行深度的复制，
//但是不知道为什么引用类型时两个结构体指向的地址不相同
//所以在此处采用了直接赋值并修改Id的方法
/*func Copy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}*/

type Prototype interface {
	Clone(Id int) Prototype
}
type hello struct {
	Str string
}
type PrototypeInstance struct {
	Id         int
	TestStruct *hello
}

func (p PrototypeInstance) Clone(Id int) PrototypeInstance {
	var result PrototypeInstance
	result = p
	result.Id = Id
	return result

}
