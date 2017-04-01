package design

type Operation interface {
	GetResult() (result float64)
}
type Add struct {
	Data []float64
}
type Muti struct {
	Data []float64
}
type Min struct {
	Data []float64
}
type Divi struct {
	Data []float64
}

func GetOperation(what string, data ...float64) Operation {
	switch what {
	case "add":
		return Add{Data: data}
	case "min":
		return Min{Data: data}
	case "muti":
		return Muti{Data: data}
	case "divi":
		return Divi{Data: data}
	default:
		return Add{Data: []float64{1, 2, 3, 5}}

	}
}
func (m Muti) GetResult() (result float64) {
	result = 1
	for _, v := range m.Data {
		result = result * v
	}
	return
}
func (a Add) GetResult() (result float64) {
	for _, v := range a.Data {
		result = result + v
	}
	return
}
func (m Min) GetResult() (result float64) {
	result = m.Data[0]
	for i := 1; i < len(m.Data); i++ {
		result = result - m.Data[i]
	}
	return
}
func (d Divi) GetResult() (result float64) {
	result = d.Data[0]
	for i := 1; i < len(d.Data); i++ {
		result = result / d.Data[i]
	}
	return
}
